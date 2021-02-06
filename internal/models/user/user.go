package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type User struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Status      int       `json:"status"`
	LastLoginAt time.Time `json:"last_login_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserFull struct {
	Uid         int      `json:"uid"`
	Name        string   `json:"name"`
	Avatar      string   `json:"avatar"`
	Email       string   `json:"email"`
	Status      int      `json:"status"`
	LastLoginAt string   `json:"last_login_at"`
	CreatedAt   string   `json:"created_at"`
	Token       string   `json:"token"`
	Roles       []string `json:"roles"`
}

type Model struct {
	Db   *gorm.DB
	Name string
}

func NewModel(db *gorm.DB) *Model {
	return &Model{
		Db:   db,
		Name: "users",
	}
}

func (s *Model) Create(in User) (ret User, err error) {
	db := s.Db.Table(s.Name)
	err = db.Omit("last_login_at").Create(&in).Error
	if err != nil {
		return
	}
	ret = in
	return
}

func (s *Model) ExistEmail(email string) (ret User, err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("email = ?", email).First(&ret).Error
	if err != nil {
		return
	}
	return
}

func (s *Model) UpdateLastLoginAtWithId(id int) (err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("id = ?", id).Updates(map[string]interface{}{
		"last_login_at": time.Now(),
	}).Error
	if err != nil {
		return
	}
	return
}

func (s *Model) UpdateWithId(id int, values map[string]interface{}) (err error) {
	db := s.Db.Table(s.Name)
	sql := "update users set "
	conditions, vals := make([]string, 0), make([]interface{}, 0)
	for k, v := range values {
		if k == "password" {
			conditions = append(conditions, fmt.Sprintf("%v = password(?)", k))
		} else {
			conditions = append(conditions, fmt.Sprintf("%v = ?", k))
		}
		vals = append(vals, v)
	}
	sql = sql + strings.Join(conditions, ",") + " where id = ?"
	vals = append(vals, id)
	err = db.Exec(sql, vals...).Error
	if err != nil {
		return
	}
	return
}

func (s *Model) UpdateWithEmail(email string, values map[string]interface{}) (err error) {
	db := s.Db.Table(s.Name)
	sql := "update users set "
	conditions, vals := make([]string, 0), make([]interface{}, 0)
	for k, v := range values {
		if k == "password" {
			conditions = append(conditions, fmt.Sprintf("%v = password(?)", k))
		} else {
			conditions = append(conditions, fmt.Sprintf("%v = ?", k))
		}
		vals = append(vals, v)
	}
	sql = sql + strings.Join(conditions, ",") + " where email = ?"
	vals = append(vals, email)
	err = db.Exec(sql, vals...).Error
	if err != nil {
		return
	}
	return
}

func (s *Model) FindWithLogin(email, password string) (ret User, err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("`email` = ? and `password` = password(?) ", email, password).First(&ret).Error
	if err != nil {
		return
	}
	return
}
