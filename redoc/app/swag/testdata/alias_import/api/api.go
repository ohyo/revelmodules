package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ohyo/revelmodules/redoc/app/swag/testdata/alias_import/data"
	"github.com/ohyo/revelmodules/redoc/app/swag/testdata/alias_type/types"
)

// @Summary Get application
// @Description test get application
// @ID get-application
// @Accept  json
// @Produce  json
// @Success 200 {object} data.ApplicationResponse	"ok"
// @Router /testapi/application [get]
func GetApplication(c *gin.Context) {
	var foo = data.ApplicationResponse{
		Application: types.Application{
			Name: "name",
		},
		ApplicationArray: []types.Application{
			{Name: "name"},
		},
	}
	log.Println(foo)
	//write your code
}
