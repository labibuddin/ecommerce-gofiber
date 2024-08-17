package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title" example:"Product A"`
	Description string `gorm:"not null" json:"description" example:"A high-quality product"`
	Amount      int    `gorm:"not null" json:"amount" example:"10"`
}
// ProductSwagger is a struct used only for Swagger documentation, containing only the fields you want to expose
type ProductSwagger struct {
	Title       string `json:"title" example:"Product A"`
	Description string `json:"description" example:"A high-quality product"`
	Amount      int    `json:"amount" example:"10"`
}


// type ProductSwagger struct {
// 	ID          uint   `json:"id"`
// 	Title       string `json:"title" example:"Product A"`
// 	Description string `json:"description" example:"A high-quality product"`
// 	Amount      int    `json:"amount" example:"10"`
// 	CreatedAt   time.Time `json:"created_at" `
// 	UpdatedAt   time.Time `json:"updated_at" `
// 	DeletedAt   *time.Time `json:"deleted_at,omitempty" `
// }

// type Product struct {
// 	gorm.Model
// 	Title       string `json:"title" example:"Product A"`
// 	Description string `json:"description" example:"A high-quality product"`
// 	Amount      int    `json:"amount" example:"10"`
// }