package entity

type Photo struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"-"`
	Title    string `gorm:"type:varchar(255)" json:"-"`
	Caption  string `gorm:"type:varchar(255)" json:"-"`
	PhotoUrl string `gorm:"type:varchar(255)" json:"-"`
	UserID   int64  `gorm:"type:bigint" json:"-"`
	User     User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
