package model

type User struct {
	Id_account string `gorm:"unique" json:"id_account"`
	Username   string `gorm:"primaryKey" json:"username"`
	Email      string `gorm:"type:varchar(255)" json:"email"`
	Password   string `gorm:"type:varchar(255)" json:"password"`
}
