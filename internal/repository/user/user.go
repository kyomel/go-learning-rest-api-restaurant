package user

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"rest-api-restaurant/internal/model"
	"time"

	"gorm.io/gorm"
)

type userRepo struct {
	db        *gorm.DB
	gcm       cipher.AEAD
	time      uint32
	memory    uint32
	threads   uint8
	keylen    uint32
	signKey   *rsa.PrivateKey
	accessExp time.Duration
}

func GetRepository(db *gorm.DB, secret string, time uint32, memory uint32, threads uint8, keylen uint32, signKey *rsa.PrivateKey, accessExp time.Duration) (Repository, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &userRepo{
		db:        db,
		gcm:       gcm,
		time:      time,
		memory:    memory,
		threads:   threads,
		keylen:    keylen,
		signKey:   signKey,
		accessExp: accessExp,
	}, nil
}

func (ur *userRepo) RegisterUser(userData model.User) (model.User, error) {
	if err := ur.db.Create(&userData).Error; err != nil {
		return model.User{}, err
	}
	return userData, nil
}

func (ur *userRepo) CheckRegistered(username string) (bool, error) {
	var userData model.User
	if err := ur.db.Where(model.User{Username: username}).First(&userData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	return userData.ID != "", nil
}

func (ur *userRepo) GetUserData(username string) (model.User, error) {
	var userData model.User
	if err := ur.db.Where(model.User{Username: username}).First(&userData).Error; err != nil {
		return userData, err
	}
	return userData, nil
}

func (ur *userRepo) VerifyLogin(username, password string, userData model.User) (bool, error) {
	if username != userData.Username {
		return false, nil
	}

	verified, err := ur.comparePassword(password, userData.Hash)
	if err != nil {
		return false, err
	}

	return verified, nil
}
