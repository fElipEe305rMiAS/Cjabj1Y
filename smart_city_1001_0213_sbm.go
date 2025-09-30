// 代码生成时间: 2025-10-01 02:13:28
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// SmartCity represents the structure for our smart city data
type SmartCity struct {
    gorm.Model
    Name        string  `gorm:"type:varchar(100)"`
    Population int     `gorm:"type:int"`
    EnergyUsage float64 `gorm:"type:float"`
}

// Database connection variable
var db *gorm.DB
var err error

// ConnectToDatabase sets up the connection to the SQLite database
func ConnectToDatabase() {
    db, err = gorm.Open(sqlite.Open("smart_city.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&SmartCity{})
}

// CreateSmartCity adds a new smart city to the database
func CreateSmartCity(city SmartCity) error {
    if result := db.Create(&city); result.Error != nil {
        return result.Error
    }
    return nil
}

// GetSmartCities retrieves all smart cities from the database
func GetSmartCities() ([]SmartCity, error) {
    var cities []SmartCity
    if result := db.Find(&cities); result.Error != nil {
        return nil, result.Error
    }
    return cities, nil
}

// UpdateSmartCity updates an existing smart city in the database
func UpdateSmartCity(id uint, updates map[string]interface{}) error {
    if result := db.Model(&SmartCity{}).Where("id = ?", id).Updates(updates); result.Error != nil {
        return result.Error
    }
    return nil
}

// DeleteSmartCity removes a smart city from the database by ID
func DeleteSmartCity(id uint) error {
    if result := db.Delete(&SmartCity{}, id); result.Error != nil {
        return result.Error
    }
    return nil
}

func main() {
    // Connect to the database
    ConnectToDatabase()

    // Create a new smart city
    newCity := SmartCity{Name: "New Metropolis", Population: 1000000, EnergyUsage: 500000.0}
    if err := CreateSmartCity(newCity); err != nil {
        fmt.Println("Error creating smart city: ", err)
    } else {
        fmt.Println("Smart city created successfully")
    }

    // Retrieve all smart cities
    cities, err := GetSmartCities()
    if err != nil {
        fmt.Println("Error retrieving smart cities: ", err)
    } else {
        for _, city := range cities {
            fmt.Printf("Name: %s, Population: %d, Energy Usage: %.2f
", city.Name, city.Population, city.EnergyUsage)
        }
    }

    // Update an existing smart city
    updates := map[string]interface{}{
        "Population": 1200000,
        "EnergyUsage": 600000.0,
    }
    if err := UpdateSmartCity(newCity.ID, updates); err != nil {
        fmt.Println("Error updating smart city: ", err)
    } else {
        fmt.Println("Smart city updated successfully")
    }

    // Delete a smart city
    if err := DeleteSmartCity(newCity.ID); err != nil {
        fmt.Println("Error deleting smart city: ", err)
    } else {
        fmt.Println("Smart city deleted successfully")
    }
}
