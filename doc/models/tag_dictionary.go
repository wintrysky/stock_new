package models

import (
	"github.com/guregu/null"
)

// TagDictionary table comment
type TagDictionary struct {
	// ID column comments
	ID int64 `gorm:"primary_key;column:id" json:"id"`
	// TagDesc description
	TagDesc string `gorm:"column:tag_desc" json:"tag_desc"`
	// TagName tag name
	TagName string `gorm:"column:tag_name" json:"tag_name"`
}

// TableName sets the insert table name for this struct type
func (t *TagDictionary) TableName() string {
	return "tag_dictionary"
}
