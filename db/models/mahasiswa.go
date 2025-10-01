package model

import (
	"time"

	"gorm.io/gorm"
)

type Mahasiswa struct {
	// Definisikan semua field secara manual dengan urutan yang diinginkan
	ID        uint           `json:"ID"`
	Nama      string         `json:"nama"`
	Nim       string         `json:"nim"`
	Prodi     string         `json:"prodi"`
	Kelas     string         `json:"kelas"`
	CreatedAt time.Time      `json:"CreatedAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `json:"DeletedAt" gorm:"index"` // tambahkan gorm:"index"
}
