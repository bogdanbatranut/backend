package dbmodels

import "time"

type InvoiceMaster struct {
	ID         uint
	CreatedAt  time.Time
	CustomerID uint     `gorm:"column:customer_id"`
	Customer   Customer `gorm:"foreignKey:customer_id"`
	Number     uint
	Series     string
	Details    []InvoiceDetail `gorm:"foreignKey:invoice_id"`
}

type Tabler interface {
	TableName() string
}

func (InvoiceMaster) TableName() string {
	return "invoices_master"
}
