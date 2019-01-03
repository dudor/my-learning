package model

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(128)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
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
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return u.FollowSelf()
}
func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", Md5(email))
}
func UpdateUserByUsername(username string, contents map[string]interface{}) error {
	item, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(item).Update(contents).Error
}
func UpdateLastSeen(username string) error {
	fields := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserByUsername(username, fields)
}
func UpdateAboutMe(username, aboutme string) error {
	fields := map[string]interface{}{"about_me": aboutme}
	return UpdateUserByUsername(username, fields)
}
func (u *User) Follow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Append(u).Error
}
func (u *User) UnFollow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Delete(u).Error
}
func (u *User) FollowSelf() error {
	return db.Model(u).Association("Followers").Append(u).Error
}
func (u *User) FollowingIDs() []int {
	var ids []int
	log.Println(u.ID)
	rows, err := db.Table("follower").Where("follower_id = ?", u.ID).Select("user_id, follower_id").Rows()
	if err != nil {
		log.Print("FollowingIDs Errors", err)
		return ids
	}
	defer rows.Close()
	for rows.Next() {
		var id, follower_id int
		rows.Scan(&id, &follower_id)
		ids = append(ids, id)
	}
	log.Print("ids=", ids)
	return ids
}
func (u *User) FollowingCount() int {
	return len(u.FollowingIDs())
}
func (u *User) FollowersCount() int {
	return db.Model(u).Association("Followers").Count()
}
func (u *User) FollowingPosts() (*[]Post, error) {
	var posts []Post
	ids := u.FollowingIDs()
	if err := db.Preload("User").Order("timestamp desc").Where("user_id in (?)", ids).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}
func (u *User) IsFollowedByUser(username string) bool {
	other, _ := GetUserByUsername(username)
	ids := other.FollowingIDs()
	for _, v := range ids {
		if v == u.ID {
			return true
		}
	}
	return false
}
func (u *User) CreatePost(body string) error {
	post := Post{
		Body:   body,
		UserID: u.ID,
	}
	return db.Create(&post).Error
}
