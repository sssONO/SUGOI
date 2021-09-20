package controller

import (
	//strconv "strconv"
	db "main/models/db"
	entity "main/models/entity"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	productName := c.PostForm("Name")
	var product = entity.Product{
		Name: productName,
	}
	db.InsertProduct(&product)
}

// func DeleteProduct(c *gin.Context) {
// 	productIDStr := c.PostForm("productID")
// 	productID, _ := strconv.Atoi(productIDStr)
// 	db.DeleteProduct(productID)
// }
