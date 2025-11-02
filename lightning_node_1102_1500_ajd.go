// 代码生成时间: 2025-11-02 15:00:19
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// LightningNode represents a lightning node in the network
type LightningNode struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string
    IP       string
    Port     int
    IsActive bool
}

// DBClient represents a database client
type DBClient struct {
    *gorm.DB
}

// NewDBClient initializes a new database client
func NewDBClient() *DBClient {
    db, err := gorm.Open(sqlite.Open("lightning_nodes.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    return &DBClient{db}
}

func main() {
    dbClient := NewDBClient()
    defer dbClient.Close()

    // Migrate the schema
    dbClient.AutoMigrate(&LightningNode{})

    // Example usage: creating a new lightning node
    newNode := LightningNode{
        Name:     "Node1",
        IP:       "192.168.1.100",
        Port:     9735,
        IsActive: true,
    }
    if err := dbClient.Create(&newNode).Error; err != nil {
        fmt.Printf("Failed to create new node: %v
", err)
        return
    }

    fmt.Printf("New lightning node created with ID: %d
", newNode.ID)

    // Fetching the node by ID
    var node LightningNode
    if err := dbClient.First(&node, newNode.ID).Error; err != nil {
        fmt.Printf("Failed to fetch node: %v
", err)
        return
    }

    fmt.Printf("Node Name: %s, IP: %s, Port: %d, IsActive: %t
", node.Name, node.IP, node.Port, node.IsActive)
}
