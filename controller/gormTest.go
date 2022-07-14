package controller

import (
	"encoding/json"
	"fmt"
	"gin_jwt/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string  `json:"code"`
	Price float64 `json:"price"`
}

func CreateController(ctx *gin.Context) {
	db := model.Db
	err := db.AutoMigrate(&Product{})
	if err != nil {
		fmt.Println(err)
	}
}

func InsertController(ctx *gin.Context) {
	parms, _ := ctx.GetRawData() // 定义map或结构体
	var req map[string]interface{}
	err := json.Unmarshal(parms, &req) // 反序列化
	if err != nil {
		fmt.Println("err: ", err)
	}

	price, ok := req["price"].(float64)
	if !ok {
		fmt.Println(ok)
	}

	fmt.Printf("%T", req["price"])

	code := req["code"].(string)
	product := Product{Code: code, Price: price}
	db := model.Db
	db.Create(&product)

}

func UpdateController(ctx *gin.Context) {
	db := model.Db
	db.Model(&Product{}).Where("id = ?", 1).Update("code", "hello man")
	db.Model(&Product{}).Where("id = ?", 2).Updates(Product{Code: "777", Price: 777})
}

func SelectController(ctx *gin.Context) {
	db := model.Db
	var products = []Product{}
	// db.First(&product)
	// db.First(&product, 2)
	// db.Find(&products)
	db.Raw("select * from products").Scan(&products)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": gin.H{"data": products},
	})
}

func DeleteController(ctx *gin.Context) {
	db:= model.Db
	db.Where("Code = ?", "777").Delete(&Product{})
}
