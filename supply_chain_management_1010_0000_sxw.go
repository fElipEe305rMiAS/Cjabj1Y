// 代码生成时间: 2025-10-10 00:00:35
package main

import (
    "fmt"
    "gorm.io/driver/sqlite" // Use sqlite for this example
    "gorm.io/gorm"
)

// Define the model structures for the supply chain entities
type Supplier struct {
    gorm.Model
    Name    string
    Address string
}

type Product struct {
    gorm.Model
    Name        string
    Description string
    Price       float64
    SupplierID  uint
    Supplier    *Supplier `gorm:"foreignKey:SupplierID"`
}

type Order struct {
    gorm.Model
    OrderDate   string
    OrderStatus string
    TotalAmount float64
    ProductID   uint
    Product     *Product `gorm:"foreignKey:ProductID"`
}

func main() {
    // Initialize a database connection
    db, err := gorm.Open(sqlite.Open("supply_chain.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&Supplier{}, &Product{}, &Order{})

    // Create a new supplier
    supplier := Supplier{Name: "Acme Corporation", Address: "123 Main St"}
    db.Create(&supplier)

    // Create a new product
    product := Product{Name: "Widget", Description: "A useful widget", Price: 19.99, SupplierID: supplier.ID}
    db.Create(&product)

    // Create a new order
    order := Order{OrderDate: "2023-04-01", OrderStatus: "Pending", TotalAmount: 19.99, ProductID: product.ID}
    db.Create(&order)

    // Retrieve and display all suppliers
    var suppliers []Supplier
    db.Find(&suppliers)
    fmt.Println("Suppliers:")
    for _, supplier := range suppliers {
        fmt.Printf("ID: %d, Name: %s, Address: %s
", supplier.ID, supplier.Name, supplier.Address)
    }

    // Retrieve and display all products
    var products []Product
    db.Find(&products)
    fmt.Println("Products:")
    for _, product := range products {
        fmt.Printf("ID: %d, Name: %s, Description: %s, Price: %.2f, SupplierID: %d
",
            product.ID, product.Name, product.Description, product.Price, product.SupplierID)
    }

    // Retrieve and display all orders
    var orders []Order
    db.Find(&orders)
    fmt.Println("Orders:")
    for _, order := range orders {
        fmt.Printf("ID: %d, OrderDate: %s, OrderStatus: %s, TotalAmount: %.2f, ProductID: %d
",
            order.ID, order.OrderDate, order.OrderStatus, order.TotalAmount, order.ProductID)
    }
}
