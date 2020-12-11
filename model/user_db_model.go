package model

type UserDB struct {
	Id       int `gorm:"primary_key; auto_increment; not_null"`
	Name     string
	Email    string
	Alamat   string
	Password string
}

func (e *UserDB) TableName() string {
	return "user"
}