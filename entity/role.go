package entity

type Role struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Accounts []Account `gorm:"foreignKey:RoleID" json:"accounts"`
}

func (r *Role) TableName() string {
	return "mst_role"
}
