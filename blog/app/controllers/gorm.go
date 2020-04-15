package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/ohyo/revelmodules/blog/app/models"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

var (
	db *gorm.DB
)

const (
	// DefaultName is
	DefaultName = "Admin"
	// DefaultRole is
	DefaultRole = "admin"
	// DefaultUsername is
	DefaultUsername = "admin"
	// DefaultPassword is
	DefaultPassword = "admin"
)

// GormController is
type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

// InitDB is
func InitDB() {
	var (
		driver, spec string
		found        bool
	)

	// Read configuration.
	if driver, found = revel.Config.String("db.driver"); !found {
		revel.AppLog.Fatal("No db.driver found.")
	}
	if spec, found = revel.Config.String("db.spec"); !found {
		revel.AppLog.Fatal("No db.spec found.")
	}

	// Open a connection.
	var err error
	db, err = gorm.Open(driver, spec)
	if err != nil {
		revel.AppLog.Fatal(err.Error())
	}

	// Enable Logger
	db.LogMode(true)
	migrate()
}

func migrate() {
	db.AutoMigrate(&models.Post{}, &models.Comment{}, &models.User{})
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(DefaultPassword), bcrypt.DefaultCost)
	db.Where(models.User{Name: DefaultName, Role: DefaultRole, Username: DefaultUsername}).
		Attrs(models.User{Password: bcryptPassword}).
		FirstOrCreate(&models.User{})
}

// Begin a transaction
func (ctrl *GormController) Begin() revel.Result {
	ctrl.Txn = db.Begin()
	return nil
}

// Rollback if it's still going (must have panicked).
func (ctrl *GormController) Rollback() revel.Result {
	if ctrl.Txn != nil {
		ctrl.Txn.Rollback()
		ctrl.Txn = nil
	}
	return nil
}

// Commit the transaction.
func (ctrl *GormController) Commit() revel.Result {
	if ctrl.Txn != nil {
		ctrl.Txn.Commit()
		ctrl.Txn = nil
	}
	return nil
}

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*GormController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GormController).Commit, revel.AFTER)
	revel.InterceptMethod((*GormController).Rollback, revel.FINALLY)
}
