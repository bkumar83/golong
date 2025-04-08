// Configuration update for Ory Kratos integration and database connection
package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "backend/models"
)

var DB *gorm.DB

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file, using system environment variables")
    }
}

// By default, use Ory Kratos
func IsKratosEnabled() bool {
    useKratos := os.Getenv("USE_ORY_KRATOS")
    return useKratos == "true" || useKratos == ""
}

func ConnectDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // AutoMigrate to keep the schema updated
    if err := DB.AutoMigrate(&models.User{}); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    log.Println("Database connected and migrated successfully")
}
