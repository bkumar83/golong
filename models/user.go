package models

type User struct {
    ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
    CreatedAt string `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt string `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
    Username  string `json:"username" gorm:"not null;unique"`
    Password  string `json:"password" gorm:"not null"`
}
