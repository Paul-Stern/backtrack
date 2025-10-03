package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	NameID uint
	Name   Name
}

type Link struct {
	gorm.Model
	ParentID uint
	ChildID  uint
}

type Name struct {
	gorm.Model
	Name string
}
