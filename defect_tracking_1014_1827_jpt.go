// 代码生成时间: 2025-10-14 18:27:32
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// DefectModel 定义缺陷信息结构体
type DefectModel struct {
    gorm.Model
    Title       string
    Description string
    Status      string // 例如: Open, In Progress, Resolved
}

func main() {
    // 连接数据库
    db, err := gorm.Open(sqlite.Open("defects.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database"))
    }
    defer db.Close()

    // 自动迁移模式
    db.AutoMigrate(&DefectModel{})

    // 创建缺陷
    defect := DefectModel{
        Title:       "Example Defect",
        Description: "This is an example defect.",
        Status:      "Open",
    }
    if err := db.Create(&defect).Error; err != nil {
        fmt.Printf("Failed to create defect: %s
", err)
    } else {
        fmt.Printf("Defect created successfully with ID: %d
", defect.ID)
    }

    // 读取所有缺陷
    var defects []DefectModel
    if err := db.Find(&defects).Error; err != nil {
        fmt.Printf("Failed to retrieve defects: %s
", err)
    } else {
        fmt.Printf("Retrieved %d defects: 
", len(defects))
        for _, defect := range defects {
            fmt.Printf("ID: %d, Title: %s, Status: %s
", defect.ID, defect.Title, defect.Status)
        }
    }
}
