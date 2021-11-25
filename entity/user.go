package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;unique;index" json:"id"`
	Nama      string    `gorm:"type:varchar(255)" json:"nama"`
	Username  string    `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	Password  string    `gorm:"->;<-;not null" json:"-"`
	Token     string    `gorm:"-" json:"token,omitempty"`
	IsActive  bool      `gorm:"not null; column:is_active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Prepare() error {
	u.ID = uuid.NewV4()
	u.IsActive = true
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
