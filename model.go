package main

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//MYTYPE -> new datatype
type MYTYPE uuid.UUID

//GormDataType -> sets type to binary(16)
func (my MYTYPE) GormDataType() string {
	return "binary(16)"
}

// Scan --> From DB
func (my *MYTYPE) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = MYTYPE(parseByte)
	return err
}

// Value -> TO DB
func (my MYTYPE) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}

// Post is the model for the posts table.
type Post struct {
	ID      MYTYPE `gorm:"primary_key;"`
	Name    string `gorm:"not null"`
	Comment Comment
}

// BeforeCreate ->
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	p.ID = MYTYPE(id)
	return err
}

// Comment is the model for the comments table.
type Comment struct {
	ID     uuid.UUID `gorm:"type:binary(16);primary_key;default:(UUID_TO_BIN(UUID()));"`
	Name   string    `gorm:"size:128;not null;"`
	PostID MYTYPE
}
