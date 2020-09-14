package types

import (
	"reflect"
	"time"
)

var (
	// ProductNameMongoTag ...
	ProductNameMongoTag string
	// PriceMogoTag ...
	PriceMogoTag string
	// LastUpdateMongoTag ...
	LastUpdateMongoTag string
	// CountUpdateMongoTag ...
	CountUpdateMongoTag string
)

// ProductPrice ...
type ProductPrice struct {
	ProductName string    `bson:"product_name"`
	Price       float64   `bson:"price"`
	CountUpdate int64     `bson:"count"`
	LastUpdate  time.Time `bson:"lastModified"`
}

const (
	// ASC ...
	ASC = 1
	// DESC ...
	DESC = -1
)

// ProductPriceListParams ...
type ProductPriceListParams struct {
	SortField string
	SortOrder int64
	Page      int64
	LenPage   int64
}

const tagName = "bson"

func init() {
	p := ProductPrice{}
	f, _ := reflect.TypeOf(p).FieldByName("ProductName")
	ProductNameMongoTag = f.Tag.Get(tagName)
	f, _ = reflect.TypeOf(p).FieldByName("Price")
	PriceMogoTag = f.Tag.Get(tagName)
	f, _ = reflect.TypeOf(p).FieldByName("LastUpdate")
	LastUpdateMongoTag = f.Tag.Get(tagName)
	f, _ = reflect.TypeOf(p).FieldByName("CountUpdate")
	CountUpdateMongoTag = f.Tag.Get(tagName)
}
