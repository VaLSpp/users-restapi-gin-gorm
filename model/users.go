package model

type Users struct {
	ID           uint         `json:"id" form:"id" gorm:"type:int; primaryKey; autoIncrement"`
	Username     string      `json:"username" form:"username" gorm:"not null; type:varchar(50); unique_index;"`
	PasswordHash string      `json:"password_hash" form:"password_hash" gorm:"not null; type:varchar(255); unique_index;"`
	Email        string      `json:"email" form:"email" gorm:"not null; type:varchar(255); unique_index;"`
	RoleId       uint         `json:"role_id" form:"role_id" gorm:"type:int; unique_index"`
	UserDetails  UserDetails `gorm:"foreignKey:UserID"`
}
