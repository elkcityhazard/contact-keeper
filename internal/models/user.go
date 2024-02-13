package models

import (
	"reflect"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserScope int

type UserStatus int

const (
	Default UserScope = iota
	Admin
	Owner
)

const (
	InActive UserStatus = iota
	Active
)

const (
	passwordHashRounds = 10
)

type User struct {
	ID         int64
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
	Password   Password  `json:"-"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Version    int       `json:"version"`
	Scope      int       `json:"scope"`
	UserStatus int       `json:"user_status"`
}

type Password struct {
	Hash      []byte `json:"-"`
	PlainText string `json:"-"`
}

func (p *Password) GenerateHashedPassword() {

	hash, err := bcrypt.GenerateFromPassword([]byte(p.PlainText), passwordHashRounds)

	if err != nil {
		panic(err)
	}

	p.Hash = hash
}

func (p *Password) ValidatePassword(plainText string) bool {
	return bcrypt.CompareHashAndPassword(p.Hash, []byte(plainText)) == nil
}

type Contact struct {
	ID             int64          `json:"id"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	AddressLine1   string         `json:"addressLine1"`
	AddressLine2   string         `json:"addressLine2,omitempty"` // omitempty if the field is optional
	City           string         `json:"city"`
	State          string         `json:"state"`
	Country        string         `json:"country"`
	Zipcode        string         `json:"zipcode"`
	EmailAddresses []EmailAddress `json:"emailAddresses"`
	PhoneNumbers   []PhoneNumber  `json:"phoneNumbers"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	Version        int            `json:"version"`
	Activated      bool           `json:"activated"`
}

// Assuming EmailAddress and PhoneNumber are already defined with JSON struct tags

type EmailAddress struct {
	ID        int64     `json:"-"`
	Address   string    `json:"emailAddress"`
	ContactID int64     `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Version   int       `json:"version"`
	Activate  bool      `json:"isActivate"`
}

type PhoneNumber struct {
	ID        int64     `json:"-"`
	Number    string    `json:"phoneNumber"`
	ContactID int64     `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Version   int       `json:"version"`
	Activate  bool      `json:"isActivate"`
}

// Required checks if the specified keys in the User struct are set and not empty.
// It takes a variable number of string keys as input and returns a boolean.
// Note: the keys are case-sensitive
func (u *User) Required(keys ...string) bool {

	// reflect the value of the pointer struct

	val := reflect.ValueOf(u).Elem()

	for _, key := range keys {
		fieldValue := val.FieldByName(key)

		if !fieldValue.IsValid() {
			return false
		}

		if fieldValue.Interface() == reflect.Zero(fieldValue.Type()).Interface() {
			return false
		}

		if fieldValue.Kind() == reflect.String && fieldValue.String() == "" {
			return false
		}
	}

	return true

}
