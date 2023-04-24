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

type RoleUserGroup struct {
	Label string `json:"label" gorm:"label"`
	Value int    `json:"id" gorm:"primary_key"`
}
