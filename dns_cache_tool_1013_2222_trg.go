// 代码生成时间: 2025-10-13 22:22:49
package main

import (
    "crypto/tls"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "github.com/joho/godotenv"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Configuration for the application
type Config struct {
    CacheExpiration time.Duration
}

// DNSRecord represents a DNS record in the cache
type DNSRecord struct {
    ID        string    `gorm:"primaryKey" json:"id"`
    Host      string    `json:"host"`
    IP        string    `json:"ip"`
    CreatedAt time.Time `json:"createdAt"`
}

var db *gorm.DB
var redisClient *redis.Client
var config Config

func init() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func main() {
    // Initialize the database connection
    db, err := gorm.Open(sqlite.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&DNSRecord{})

    // Initialize the Redis client
    redisClient = redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
        Password: os.Getenv("REDIS_PASSWORD"),
        DB:       0,  // use default DB
    })

    // Set up the Gin router
    router := gin.Default()

    // Route to resolve DNS and cache the result
    router.GET("/resolve", resolveDNS)

    // Start the server
    log.Println("Starting DNS cache tool on :8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

// resolveDNS handles the DNS resolution and caching
func resolveDNS(c *gin.Context) {
    host := c.Query("host")
    if host == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Host parameter is required"})
        return
    }

    // Check if the host is cached in Redis
    result, err := redisClient.Get(context.Background(), host).Result()
    if err == nil {
        c.JSON(http.StatusOK, gin.H{"ip": result})
        return
    }

    // If not cached, perform DNS resolution
    records, err := net.LookupIP(host)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to resolve DNS"})
        return
    }

    // Cache the result in Redis
    redisClient.Set(context.Background(), host, records[0].String(), config.CacheExpiration)

    // Cache the result in the database
    if err := db.Create(&DNSRecord{ID: host, Host: host, IP: records[0].String(), CreatedAt: time.Now()}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cache DNS result"})
        return
    }

    // Return the IP address
    c.JSON(http.StatusOK, gin.H{"ip": records[0].String()})
}
