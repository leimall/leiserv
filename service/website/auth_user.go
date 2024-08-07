package website

import (
	"errors"
	"leiserv/global"
	website "leiserv/models/website/types"
	"leiserv/utils"
	"time"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) SignUp(u website.ClientUser) (userInfo website.ClientUser, err error) {

	var user website.ClientUser
	if !errors.Is(global.MALL_DB.Where("username =?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInfo, errors.New("user already exists")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	u.UserId = utils.SnowflakeID()

	err = global.MALL_DB.Create(&u).Error

	return u, err
}

func (s *UserService) SignIn(u *website.ClientUser) (userInfo *website.ClientUser, err error) {
	if nil == global.MALL_DB {
		return nil, errors.New("database connection not initialized")
	}
	var user website.ClientUser
	err = global.MALL_DB.Where("email = ?", u.Email).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return userInfo, errors.New("invalid username or password")
		}
	}
	return &user, err
}

func (s *UserService) GetUserInfo(userId string) (userInfo website.ClientUser, err error) {
	var user website.ClientUser
	err = global.MALL_DB.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

func (s *UserService) UpdateUserInfo(u website.ClientUser) error {
	return global.MALL_DB.Model(&website.ClientUser{}).
		Select("update_at", "username", "header_img", "phone", "email").
		Where("user_id = ?", u.UserId).
		Updates(map[string]interface{}{
			"update_at":  time.Now(),
			"username":   u.Username,
			"header_img": u.HeaderImg,
			"phone":      u.Phone,
			"email":      u.Email,
		}).Error
}
