package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name         string
	ClientID     string `gorm:"uniqueIndex"`
	ClientSecret string
	RedirectURI  string
}
