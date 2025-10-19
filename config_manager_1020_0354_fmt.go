// 代码生成时间: 2025-10-20 03:54:15
// config_manager.go

package main

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Config 用于存储配置信息
type Config struct {
    ID        uint   "gorm:"primary_key" json:"id""
    Key       string "gorm:"type:varchar(100);" json:"key""
    Value     string "gorm:"type:text" json:"value""
    CreatedAt string "gorm:"type:datetime" json:"createdAt""
    UpdatedAt string "gorm:"type:datetime" json:"updatedAt""
}

// ConfigManager 配置文件管理器
type ConfigManager struct {
    db *gorm.DB
}

// NewConfigManager 创建一个新的ConfigManager实例
func NewConfigManager(dbPath string) (*ConfigManager, error) {
    var db *gorm.DB
    var err error
    
    // 尝试连接数据库
    db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    // 迁移数据库模式
    err = db.AutoMigrate(&Config{})
    if err != nil {
        return nil, err
    }
    
    return &ConfigManager{db: db}, nil
}

// SaveConfig 保存配置项
func (cm *ConfigManager) SaveConfig(key, value string) error {
    var config Config
    
    // 尝试查找配置项
    if err := cm.db.Where("key = ?", key).First(&config).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // 如果配置项不存在，则创建新配置项
            config = Config{
                Key:   key,
                Value: value,
            }
            if err := cm.db.Create(&config).Error; err != nil {
                return err
            }
        } else {
            return err
        }
    } else {
        // 如果配置项存在，则更新配置项
        config.Value = value
        if err := cm.db.Save(&config).Error; err != nil {
            return err
        }
    }
    return nil
}

// LoadConfig 加载配置项
func (cm *ConfigManager) LoadConfig(key string) (string, error) {
    var config Config
    
    // 尝试查找配置项
    if err := cm.db.Where("key = ?", key).First(&config).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return "", nil
        }
        return "", err
    }
    
    return config.Value, nil
}

func main() {
    // 创建ConfigManager实例
    cm, err := NewConfigManager("config.db")
    if err != nil {
        log.Fatalf("Failed to create ConfigManager: %v", err)
    }
    
    // 保存配置项
    if err := cm.SaveConfig("example_key", "example_value"); err != nil {
        log.Fatalf("Failed to save config: %v", err)
    }
    
    // 加载配置项
    configValue, err := cm.LoadConfig("example_key")
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    fmt.Println("Loaded config value:", configValue)
}
