// 代码生成时间: 2025-10-05 17:07:44
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Product represents a product in the live commerce system.
type Product struct {
    gorm.Model
    Name        string  `gorm:"column:name;type:varchar(255)"`
    Description string  `gorm:"column:description;type:text"`
    Price       float64 `gorm:"column:price;type:decimal(10,2)"`
    Quantity    int     `gorm:"column:quantity"`
}

// LiveStream represents a live stream session in the live commerce system.
type LiveStream struct {
    gorm.Model
    Title       string    `gorm:"column:title;type:varchar(255)"`
    HostID      uint      `gorm:"column:host_id"`
    Host        User      `gorm:"foreignKey:HostID"`
    Products    []Product `gorm:"foreignKey:LiveStreamID"`
    LiveStreamID uint
}

// User represents a user in the live commerce system.
type User struct {
    gorm.Model
    Name     string `gorm:"column:name;type:varchar(255)"`
    Email    string `gorm:"column:email;type:varchar(255);uniqueIndex"`
    Password string `gorm:"column:password;type:varchar(255)"`
}

// DB represents the database connection.
var DB *gorm.DB

func main() {
    var err error
    // Connect to the SQLite database
    DB, err = gorm.Open(sqlite.Open("live_commerce.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    DB.AutoMigrate(&Product{}, &LiveStream{}, &User{})

    // Seed some data into the database
    seedData()

    // Start a new live stream session
    startLiveStream()
}

// seedData populates the database with some initial data for demonstration purposes.
func seedData() {
    users := []User{
        {Name: "Alice", Email: "alice@example.com", Password: "password123"},
        {Name: "Bob", Email: "bob@example.com", Password: "password123"},
    }
    for _, user := range users {
        DB.Create(&user)
    }

    products := []Product{
        {Name: "Laptop", Description: "A high-performance laptop", Price: 999.99, Quantity: 10},
        {Name: "Smartphone", Description: "A latest model smartphone", Price: 599.99, Quantity: 20},
    }
    for _, product := range products {
        DB.Create(&product)
    }
}

// startLiveStream creates a new live stream session and associates it with a host and products.
func startLiveStream() {
    host := User{Name: "Alice", Email: "alice@example.com"}
    DB.Where(&User{Name: host.Name, Email: host.Email}).First(&host)

    liveStream := LiveStream{
        Title:       "Tech Gadgets Sale",
        HostID:      host.ID,
        LiveStreamID: host.ID,
    }
    
    DB.Create(&liveStream)

    products := []Product{
        {Name: "Laptop", Description: "A high-performance laptop", Price: 999.99, Quantity: 10},
        {Name: "Smartphone", Description: "A latest model smartphone", Price: 599.99, Quantity: 20},
    }
    for _, product := range products {
        DB.Model(&liveStream).Association("Products").Append(&product)
    }

    fmt.Println("Live stream started successfully!")
}
