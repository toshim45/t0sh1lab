package main

import (
	"fmt"

	"github.com/guregu/null"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/satori/go.uuid"
)

func setupDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&timeout=5s",
		"root",
		"root",
		"localhost:3306",
		"ponirah",
	))
	if err != nil {
		fmt.Printf("Cannot connect to MySQL DB:\n%v\n", err)
		return
	}

	fmt.Println("Connected to MySQL %s", "localhost")
	db.DB().SetMaxOpenConns(5)
	db.SingularTable(true)
	db.LogMode(true)

	return
}

type PurchaseOrderItem struct {
	ID                 uuid.UUID   `json:"id" gorm:"column:purchase_order_item_id;primary_key"`
	PurchaseOrderID    uuid.UUID   `json:"purchaseOrderID"`
	ProductEntityID    uuid.UUID   `json:"productEntityID"`
	Title              string      `json:"title"`
	ItemType           int         `json:"itemType"` // enum ItemType?
	CategoryID         int64       `json:"categoryID"`
	Quantity           int         `json:"quantity"`
	Price              float64     `json:"price"`
	NetPrice           float64     `json:"netPrice"`
	SkuAggregator      string      `json:"skuAggregator"`
	SkuVariant         null.String `json:"skuVariant"`
	QuantityReceived   null.Int    `json:"quantityReceived"`
	QuantityRejected   null.Int    `json:"quantityRejected"`
	QuantityQuarantine null.Int    `json:"quantityQuarantine"`
}

func main() {
	db, err := setupDB()
	if err != nil {
		return
	}
	defer db.Close()

	var result PurchaseOrderItem

	poiID := uuid.FromStringOrNil("ad8b8c29-6b56-4c77-9215-138557d9a1c7")

	if err := db.First(&result, "purchase_order_item_id = ?", poiID.String()).Error; err != nil {
		fmt.Printf("Could not find 'purchase_order_item' with poiID %s \n%v\n", poiID.String(), err)
	}

	fmt.Printf("result: %#v\n", result)

	var results []PurchaseOrderItem
	poiIDs := []uuid.UUID{poiID}

	if err = db.Where("purchase_order_item_id IN (?)", poiIDs).Find(&results).Error; err != nil {
		fmt.Printf("Could not find 'purchase_order_item' with poiID %v \n%v\n", poiIDs, err)
	}

	fmt.Printf("results: %#v\n", results)
}
