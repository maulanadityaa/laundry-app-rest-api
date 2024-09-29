package entity

import "time"

type Account struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RoleID    string    `json:"role_id"`

	Customer Customer `gorm:"foreignKey:AccountID" json:"customer"`
	Employee Employee `gorm:"foreignKey:AccountID" json:"employee"`
}

func (c *Account) TableName() string {
	return "mst_account"
}
