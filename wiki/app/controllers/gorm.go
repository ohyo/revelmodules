package controllers

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"

	"github.com/ohyo/revelmodules/wiki/app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
)

type GormController struct {
	*revel.Controller
	db *gorm.DB
}

var (
	DB *gorm.DB
)

// Perform automatic migration
func InitDB() {
	var err error

	dbDriver := revel.Config.StringDefault("database.driver", "")
	dbSource := revel.Config.StringDefault("database.source", "")
	databaseUrl := ""

	if dbDriver == "" || dbSource == "" {
		if os.Getenv("DATABASE_URL") != "" {
			// Heroku Postgres
			databaseUrl = os.Getenv("DATABASE_URL")
		} else if os.Getenv("CLEARDB_DATABASE_URL") != "" {
			// Heroku ClearDB MySQL Database
			databaseUrl = os.Getenv("CLEARDB_DATABASE_URL")
		}

		if databaseUrl != "" {
			re, _ := regexp.Compile("([^:]+)://([^:]+):([^@]+)@([^/]+)/([^?]+)")
			match := re.FindStringSubmatch(databaseUrl)

			dbDriver = match[1]
			if dbDriver == "mysql" {
				dbSource = fmt.Sprintf(
					"%s:%s@tcp(%s:3306)/%s?parseTime=true",
					match[2],
					match[3],
					match[4],
					match[5],
				)
			} else {
				dbSource = databaseUrl
			}
		}
	}

	if dbDriver == "" {
		dbDriver = "sqlite3"
	}

	if dbSource == "" {
		dbSource = "./wiki.db"
	}

	DB, err = gorm.Open(dbDriver, dbSource)

	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	DB.LogMode(revel.Config.BoolDefault("mode.dev", false))

	DB.AutoMigrate(&models.Page{}, &models.Revision{})

	DB.Model(models.Page{}).AddUniqueIndex("unique_title", "title")
}

// Start a transaction on request
func (ctrl *GormController) Begin() revel.Result {
	if ctrl.db != nil {
		return nil
	}
	db := DB.Begin()
	if db.Error != nil {
		panic(db.Error)
	}
	ctrl.db = db
	return nil
}

// Commit the transaction at the end of the request
func (ctrl *GormController) Commit() revel.Result {
	if ctrl.db == nil {
		return nil
	}
	ctrl.db.Commit()
	if err := ctrl.db.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	ctrl.db = nil
	return nil
}

// Abandon transaction when abnormal
func (ctrl *GormController) Rollback() revel.Result {
	if ctrl.db == nil {
		return nil
	}
	ctrl.db.Rollback()
	if err := ctrl.db.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	ctrl.db = nil
	return nil
}
