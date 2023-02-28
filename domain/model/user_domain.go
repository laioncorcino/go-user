package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomain struct {
	id       string
	name     string
	age      int8
	email    string
	password string
}

func (user *UserDomain) GetID() string {
	return user.id
}

func (user *UserDomain) SetID(userId string) {
	user.id = userId
}

func (user *UserDomain) GetEmail() string {
	return user.email
}

func (user *UserDomain) SetEmail(email string) {
	user.email = email
}

func (user *UserDomain) GetPassword() string {
	return user.password
}

func (user *UserDomain) SetPassword(pass string) {
	user.password = pass
}

func (user *UserDomain) GetName() string {
	return user.name
}

func (user *UserDomain) SetName(name string) {
	user.name = name
}

func (user *UserDomain) GetAge() int8 {
	return user.age
}

func (user *UserDomain) SetAge(age int8) {
	user.age = age
}

func (user *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.password))
	user.password = hex.EncodeToString(hash.Sum(nil))
}
