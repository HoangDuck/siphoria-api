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
