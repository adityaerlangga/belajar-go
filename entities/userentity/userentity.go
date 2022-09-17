package userentity

type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Nama     string `gorm:"type:varchar(300)" json:"nama"`
	Username string `gorm:"type:varchar(300)" json:"username"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Password string `gorm:"type:varchar(300)" json:"password"`
}
