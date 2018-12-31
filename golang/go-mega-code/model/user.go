package model

import (
	"fmt"
	"time"
)

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(128)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	Posts        []Post
	Followers    []User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
	LastSeen     *time.Time
	AboutMe      string `gorm:"type:varchar(150)"`
	Avatar       string `gorm:"type:varchar(200)"`
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
func AddUser(username, password, email string) error {
	u := User{Username: username, Email: email}
	u.SetPassword(password)
	u.SetAvatar(email)
	return db.Create(&u).Error
}
func (u *User)SetAvatar(email string)  {
	u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon",Md5(email))
}
func UpdateUserByUsername(username string,contents map[string]interface{}) error {
	item,err:= GetUserByUsername(username)
	if err!=nil{
		return err
	}
	return db.Model(item).Update(contents).Error
}
func UpdateLastSeen(username string )error  {
	fields:= map[string]interface{}{"last_seen":time.Now()}
	return UpdateUserByUsername(username,fields)
}
func UpdateAboutMe(username,aboutme string) error {
	fields := map[string]interface{}{"about_me":aboutme}
	return UpdateUserByUsername(username,fields)
}