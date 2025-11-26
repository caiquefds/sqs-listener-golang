package domain

import (
	"time"
)

type OrderStatus string

const (
	Pending   OrderStatus = "PENDING"
	Delivered OrderStatus = "DELIVERED"
)

type Order struct {
	ID         uint        `gorm:"primaryKey;autoIncrement"`
	CustomerID int         `gorm:"column:customer_id"`
	Amount     float64     `gorm:"column:amount"`
	Status     OrderStatus `gorm:"type:varchar(50)"`
	CreatedAt  time.Time   `gorm:"column:created_at"`
}

type OrderListenerDTO struct {
	CustomerID int
	Amount     float64
}
