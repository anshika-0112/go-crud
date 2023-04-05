package product

import (
	"database/sql"
	"fmt"

	"github.com/anshika-0112/go-crud/database"
)

func getProduct(productID int) (*Product, error) {
	row := database.DbConn.QueryRow(`SELECT productID,manufacturer,sku,upc,pricePerUnit,quantityOnHand,name FROM products WHERE productID = ?`, productID)
	var product Product
	err := row.Scan(&product.ProductID, &product.Manufacturer, &product.Sku, &product.Upc, &product.PricePerUnit, &product.QuantityOnHand, &product.ProductName)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return &product, err
	}
}

func getProductList() ([]Product, error) {
	result, err := database.DbConn.Query(`SELECT productID,manufacturer,sku,upc,pricePerUnit,quantityOnHand,name FROM products`)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	products := make([]Product, 0)
	for result.Next() {
		var product Product
		result.Scan(&product.ProductID, &product.Manufacturer, &product.Sku, &product.Upc, &product.PricePerUnit, &product.QuantityOnHand, &product.ProductName)
		products = append(products, product)
	}
	return products, nil
}

func removeProduct(productID int) error {
	_, err := database.DbConn.Query(`DELETE from products where productId=?`, productID)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(product Product) error {
	_, err := database.DbConn.Exec(`UPDATE products SET 
	manufacturer=?,
	sku=?,
	upc=?,
	pricePerUnit=?,
	quantityOnHand=?,
	name=? 
	WHERE 
	productId=?`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName,
		product.ProductID,
	)
	if err != nil {
		return err
	}
	fmt.Println(err)
	return err
}

func addProduct(product Product) (int, error) {
	result, err := database.DbConn.Exec(`INSERT INTO products
	(manufacturer,
		sku,
		upc,
		pricePerUnit,
		quantityOnHand,
		name) VALUES (?,?,?,?,?,?)`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(insertID), nil
}
