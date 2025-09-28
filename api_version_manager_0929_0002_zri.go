// 代码生成时间: 2025-09-29 00:02:58
package main

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// APIVersion is the model for API versions
type APIVersion struct {
    gorm.Model
    Version string `gorm:"unique;not null"` // API version string
    Status  string `gorm:"not null"`         // Status of the API version: 'active', 'deprecated', 'archived'
}

// APIVersionManager is the struct that handles API version operations
type APIVersionManager struct {
    DB *gorm.DB
}

// NewAPIVersionManager creates a new APIVersionManager instance
func NewAPIVersionManager() *APIVersionManager {
    db, err := gorm.Open(sqlite.Open("api_versions.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Migrate the schema
    db.AutoMigrate(&APIVersion{})

    return &APIVersionManager{DB: db}
}

// AddVersion adds a new API version to the database
func (m *APIVersionManager) AddVersion(version, status string) error {
    var apiVersion APIVersion
    apiVersion.Version = version
    apiVersion.Status = status

    result := m.DB.Create(&apiVersion)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// GetVersions returns a list of all API versions
func (m *APIVersionManager) GetVersions() ([]APIVersion, error) {
    var versions []APIVersion
    result := m.DB.Find(&versions)
    if result.Error != nil {
        return nil, result.Error
    }
    return versions, nil
}

// UpdateVersion updates the status of an API version
func (m *APIVersionManager) UpdateVersion(id, status string) error {
    result := m.DB.Model(&APIVersion{}).Where("id = ?", id).Update("status", status)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// DeleteVersion deletes an API version from the database
func (m *APIVersionManager) DeleteVersion(id string) error {
    result := m.DB.Delete(&APIVersion{}, id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func main() {
    manager := NewAPIVersionManager()
    defer manager.DB.Close()

    // Add a new API version
    if err := manager.AddVersion("1.0.0", "active"); err != nil {
        log.Printf("Error adding version: %v", err)
    }

    // Get all API versions
    versions, err := manager.GetVersions()
    if err != nil {
        log.Printf("Error getting versions: %v", err)
    } else {
        for _, v := range versions {
            fmt.Printf("Version: %s, Status: %s
", v.Version, v.Status)
        }
    }
}
