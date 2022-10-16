package entity

type Comment struct {
	ID      int64  `gorm:"primary_key:auto_increment" json:"-"`
	UserId  int64  `gorm:"type:bigint" json:"-"`
	PhotoId int64  `gorm:"type:bigint" json:"-"`
	Message string `gorm:"type:varchar(255)" json:"-"`
	User    User   `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
