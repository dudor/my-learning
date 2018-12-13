package model

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(128)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	Posts        []Post
	Followers    []User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

func (u *User) SetPassword(password string) {
	u.PasswordHash = GeneratePasswordHash(password)
}
func (u *User) CheckPassword(password string) bool {
	return u.PasswordHash == GeneratePasswordHash(password)
}
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Debug().Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
