package handlers

import (
	"time"

	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/padinky/imperial-fleet/helper"
	"github.com/padinky/imperial-fleet/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User model.User
type Session model.Session

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		db: db,
	}
}

func (ah *AuthHandler) GetUser(sessionid uuid.UUID) (User, error) {
	db := ah.db
	query := Session{SessionID: sessionid}
	found := Session{}
	err := db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return User{}, err
	}
	user := User{}
	usrQuery := User{ID: found.UserRef}
	err = db.First(&user, &usrQuery).Error
	if err == gorm.ErrRecordNotFound {
		return User{}, err
	}
	user.SessionID = sessionid
	return user, nil
}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	db := ah.db
	json := new(LoginRequest)
	if err := c.BodyParser(json); err != nil {
		return helper.ResponseBadRequest(c, "Invalid JSON")
	}

	found := User{}
	query := User{Email: json.Email}
	err := db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return helper.ResponseNotFound(c, "Email not found")
	}
	if !comparePasswords(found.Password, []byte(json.Password)) {
		return helper.ResponseUnauthorized(c, "Invalid Password")
	}
	session := Session{UserRef: found.ID, Expires: SessionExpires(), SessionID: uuid.New()}
	db.Create(&session)
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Expires:  SessionExpires(),
		Value:    session.SessionID.String(),
		HTTPOnly: true,
	})
	return helper.ResponseSuccessWithCodeAndData(c, session)
}

func (ah *AuthHandler) Logout(c *fiber.Ctx) error {
	db := ah.db
	user := c.Locals("user").(User)
	session := Session{}
	query := Session{SessionID: user.SessionID}
	err := db.First(&session, &query).Error
	if err == gorm.ErrRecordNotFound {
		return helper.ResponseNotFound(c, "session not found")
	}
	db.Delete(&session)
	c.ClearCookie("session_id")
	return helper.ResponseSuccessOnly(c)
}

func (ah *AuthHandler) CreateUser(c *fiber.Ctx) error {
	type CreateUserRequest struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Email    string `json:"email"`
	}

	db := ah.db
	json := new(CreateUserRequest)
	if err := c.BodyParser(json); err != nil {
		return helper.ResponseBadRequest(c, "Invalid JSON")
	}
	password := hashAndSalt([]byte(json.Password))
	err := checkmail.ValidateFormat(json.Email)
	if err != nil {
		return helper.ResponseBadRequest(c, "Invalid Email Address")
	}
	new := User{
		Name:     json.Name,
		Password: password,
		Email:    json.Email,
	}
	found := User{}
	query := User{Email: json.Email}
	err = db.First(&found, &query).Error
	if err != gorm.ErrRecordNotFound {
		return helper.ResponseBadRequest(c, "Email already registered")
	}
	db.Create(&new)
	// session := Session{UserRef: new.ID, SessionID: uuid.New()}
	// err = db.Create(&session).Error
	// if err != nil {
	// 	return helper.ResponseError(c, "Error Occured")
	// }
	// c.Cookie(&fiber.Cookie{
	// 	Name:     "session_id",
	// 	Expires:  time.Now().Add(5 * 24 * time.Hour),
	// 	Value:    session.SessionID.String(),
	// 	HTTPOnly: true,
	// })

	return helper.ResponseSuccessOnly(c)
}

func (ah *AuthHandler) GetUserInfo(c *fiber.Ctx) error {
	user := c.Locals("user").(User)
	return helper.ResponseSuccessWithCodeAndData(c, user)
}

func (ah *AuthHandler) DeleteUser(c *fiber.Ctx) error {
	type DeleteUserRequest struct {
		Password string `json:"password"`
	}
	db := ah.db
	json := new(DeleteUserRequest)
	user := c.Locals("user").(User)
	if err := c.BodyParser(json); err != nil {
		return helper.ResponseBadRequest(c, "Invalid JSON")
	}
	if !comparePasswords(user.Password, []byte(json.Password)) {
		return helper.ResponseUnauthorized(c, "Invalid Password")
	}
	db.Model(&user).Association("Sessions").Delete()
	db.Model(&user).Association("Products").Delete()
	db.Delete(&user)
	c.ClearCookie("session_id")
	return helper.ResponseSuccessWithCodeAndData(c, nil)
}

func (ah *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	type ChangePasswordRequest struct {
		Password    string `json:"password"`
		NewPassword string `json:"new_password"`
	}
	db := ah.db
	user := c.Locals("user").(User)
	json := new(ChangePasswordRequest)
	if err := c.BodyParser(json); err != nil {
		return helper.ResponseBadRequest(c, "Invalid JSON")
	}
	if !comparePasswords(user.Password, []byte(json.Password)) {
		return helper.ResponseUnauthorized(c, "Invalid Password")
	}
	user.Password = hashAndSalt([]byte(json.NewPassword))
	db.Save(&user)

	return helper.ResponseSuccessWithCodeAndData(c, nil)
}

func hashAndSalt(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return string(hash)
}
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}

// Universal date the Session Will Expire
func SessionExpires() time.Time {
	return time.Now().Add(5 * 24 * time.Hour)
}
