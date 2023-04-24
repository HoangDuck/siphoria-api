package model

type Role int

const (
	CUSTOMER Role = iota
	STAFF
	ADMIN
)

func (r Role) String() string {
	return []string{"CUSTOMER", "STAFF", "ADMIN"}[r]
}

type RoleModel struct {
	ID    int    `json:"-" gorm:"primary_key;autoIncrement"`
	Role  string `json:"role" gorm:"role"`
	Label string `json:"label" gorm:"label"`
	Value int    `json:"value" gorm:"values"`
}
