package model

type Role int

const (
	CUSTOMER Role = iota
	STAFF
	ADMIN
	MANAGER
	HOTELIER
	ACCOUNTANT
	SUPERADMIN
)

func (r Role) String() string {
	return []string{"1", "4", "51", "5", "2", "3", "66"}[r]
}

type RoleModel struct {
	ID    int    `json:"-" gorm:"primary_key;autoIncrement"`
	Role  string `json:"role" gorm:"role"`
	Label string `json:"label" gorm:"label"`
	Value int    `json:"value" gorm:"values"`
}
