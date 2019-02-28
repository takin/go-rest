package akin

import (
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

var hashkey = beego.AppConfig.String("hashmackey")
var jwtSecretKey = beego.AppConfig.String("jwtsecretkey")

// AuthString ...
type AuthString struct {
	Password string
}

// AuthToken token model
type AuthToken struct {
	Token string
}

func generatePassword(text string) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)
	return
}

// CreateHashString method untuk men-genertae Hash dan mengembalikannya dalam bentuk string
func (t *AuthString) CreateHashString() (hashedPass string, err error) {
	hash, err := generatePassword(t.Password)
	if err != nil {
		hashedPass = ""
		return
	}
	hashedPass = string(hash)
	return
}

// CreateHashByte method untuk men-generate Hash dalam bentuk Byte
func (t *AuthString) CreateHashByte() (hash []byte, err error) {
	hash, err = generatePassword(t.Password)
	return
}

// ValidatePassword method untuk memvalidasi password
func (t *AuthString) ValidatePassword() (pass string, valid bool) {
	valid = false
	pass = ""
	if hash, err := generatePassword(t.Password); err == nil {
		if e := bcrypt.CompareHashAndPassword(hash, []byte(t.Password)); e != nil {
			valid = false
		} else {
			valid = true
			pass = string(hash)
		}
		return
	}
	return
}
