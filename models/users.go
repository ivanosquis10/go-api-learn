package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Base contiene los campos comunes para las tablas.
type Base struct {
    ID        uuid.UUID      `gorm:"type:uuid;primary_key;"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// BeforeCreate establece un nuevo UUID antes de crear un registro.
func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
    // Genera un nuevo UUID para el ID.
    base.ID = uuid.New()

    return
}

type User struct {
	Base
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

// HashPassword es un hook de GORM para hashear la contraseña antes de crear un usuario.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    // Llama al hook de la Base para generar el UUID
	if err := u.Base.BeforeCreate(tx); err != nil {
		return err
	}
	
	// Hashea la contraseña antes de guardarla en la DB.
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return
}