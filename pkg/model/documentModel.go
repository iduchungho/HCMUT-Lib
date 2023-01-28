package model

type Document struct {
	Id         string `gorm:"primaryKey" json:"id"`
	Author     string `gorm:"type:varchar(255)" json:"author"`
	Author_id  string `gorm:"type:varchar(255)" json:"author_id"`
	Category   string `gorm:"type:varchar(255)" json:"category"`
	Title      string `gorm:"type:varchar(255)" json:"title"`
	Decription string `gorm:"type:varchar(255)" json:"decription"`
}
