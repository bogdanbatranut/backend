package dbmodels

import "time"

type InvoiceDetail struct {
	ID        uint
	CreatedAt time.Time
	InvoiceID uint
	ArticleID uint    `gorm:"column:article_id"`
	Article   Article `gorm:"foreignKey:article_id"`
}

func (InvoiceDetail) TableName() string {

	return "invoices_details"
}
