package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserJson struct {
	ID       uuid.UUID       `gorm:"type:uuid;primaryKey"`
	UserData json.RawMessage `gorm:"type:jsonb"`
	Notes    json.RawMessage `gorm:"type:jsonb"`
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// Auto Migrate
	if err := db.AutoMigrate(&UserJson{}); err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	// Insert sample user
	// user := UserJson{
	// 	ID:       uuid.New(),
	// 	UserData: json.RawMessage(`{"name": "Faiz", "email": "f@zenithive.com"}`),
	// 	Notes:    json.RawMessage(`{"note1": "hello", "todoooooo": "launch CRM"}`),
	// }

	// if err := db.Create(&user).Error; err != nil {
	// 	log.Fatal("Failed to insert:", err)
	// }

	// fmt.Println("Inserted user:", user.ID)

	// Query users where notes contain key "todo"
	var users []UserJson
	db.Raw(`SELECT * FROM user_jsons WHERE notes ? 'todoooooo'`).Scan(&users)

	for _, u := range users {
		fmt.Printf("Found user with TODO: %s\n", u.ID)
	}
}
