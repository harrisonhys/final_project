package entity

type Socialmedia struct {
	ID             int64  `gorm:"primary_key:auto_increment" json:"-"`
	Name           string `gorm:"type:varchar(255)" json:"-"`
	SocialMediaUrl string `gorm:"type:text" json:"-"`
	Userid         int64  `gorm:"type:varchar(255)" json:"-"`
	User           User   `gorm:"foreignkey:Userid;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
