package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gin/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CREATE TABLE `products` (
//
//		`id` bigint unsigned NOT NULL AUTO_INCREMENT,
//		`created_at` datetime(3) DEFAULT NULL,
//		`updated_at` datetime(3) DEFAULT NULL,
//		`deleted_at` datetime(3) DEFAULT NULL,
//		`code` longtext COLLATE utf8mb4_general_ci,
//		`price` double DEFAULT NULL,
//		PRIMARY KEY (`id`),
//		KEY `idx_products_deleted_at` (`deleted_at`)
//	  ) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
type Product struct {
	gorm.Model
	Code  string  `json:"code"`
	Price float64 `json:"price"`
}

// CREATE TABLE `users` (
//
//		`id` bigint unsigned NOT NULL AUTO_INCREMENT,
//		`name` longtext COLLATE utf8mb4_general_ci,
//		`email` longtext COLLATE utf8mb4_general_ci,
//		`age` tinyint unsigned DEFAULT NULL,
//		`birthday` datetime(3) DEFAULT NULL,
//		`member_number` longtext COLLATE utf8mb4_general_ci,
//		`activated_at` datetime(3) DEFAULT NULL,
//		`created_at` datetime(3) DEFAULT NULL,
//		`updated_at` datetime(3) DEFAULT NULL,
//		PRIMARY KEY (`id`)
//	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// CREATE TABLE `user2` (
//
//		`created_at` datetime(3) DEFAULT NULL,
//		`updated_at` bigint DEFAULT NULL,
//		`updated` bigint DEFAULT NULL,
//		`created` bigint DEFAULT NULL
//	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
type User2 struct {
	CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
	UpdatedAt int       // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
	// Updated int64     `gorm:"autoUpdateTime:nano"`  // 使用时间戳填纳秒数充更新时间
	Updated int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	Created int64 `gorm:"autoCreateTime"`       // 使用时间戳秒数填充创建时间
}

func CreateController(ctx *gin.Context) {
	db := model.Db
	err := db.AutoMigrate(&Product{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}

	err = db.AutoMigrate(&User2{})
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
	db := model.Db
	db.Where("Code = ?", "777").Delete(&Product{})
}
