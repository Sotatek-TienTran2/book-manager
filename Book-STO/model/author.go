package model

type Author struct {
	IdAuthor   int    `json:"idAuthor,omitempty" gorm:"column:id;type:serial;primaryKey" `
	Name       string `json:"name,omitempty" gorm:"column:name;type:varchar(100);not null"`
	NativeLand string `json:"NativeLand,omitempty" gorm:"column:native_land;type:varchar(100);not null"`
	Username   string `json:"username,omitempty" gorm:"column:username;type:varchar(100);not null"`
	Password   string `json:"password,omitempty" gorm:"column:password;type:varchar(100);not null"`
}
