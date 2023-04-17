package dao

type AccountDAO struct {
	ID         string `json:"-" gorm:"primary_key"`
	Email      string `json:"email" gorm:"email"`
	Password   string `json:"-" gorm:"password"`
	CustomerID string `json:"profile_id"`
	FirstName  string `json:"first_name" gorm:"first_name"`
	LastName   string `json:"last_name" gorm:"last_name"`
}
