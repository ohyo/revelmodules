package controllers

import (
	"math"
	"strconv"

	"github.com/ohyo/revelmodules/wiki/app/models"

	"github.com/revel/revel"
)

type List struct {
	App
}

// Display page list
func (ctrl List) Index() revel.Result {
	order := ctrl.Params.Get("order")
	if order == "" {
		order = "title"
	}

	paginateCurrent, _ := strconv.Atoi(ctrl.Params.Get("page"))
	if paginateCurrent < 1 {
		paginateCurrent = 1
	}

	limit := 20

	pages := []models.Page{}

	db := ctrl.db.Model(models.Page{})

	query := ctrl.Params.Get("query")
	if query != "" {
		db = db.Where("title like ?", "%"+query+"%")
	}

	db.Order(order).Limit(limit).Offset(limit * (paginateCurrent - 1)).Find(&pages)

	paginateTotal := 0
	db.Count(&paginateTotal)

	paginateLast := int(math.Ceil(float64(paginateTotal) / float64(limit)))
	paginatePages := make([]int, paginateLast)

	return ctrl.Render(pages, order, query, paginateCurrent, paginateTotal, paginateLast, paginatePages)
}
