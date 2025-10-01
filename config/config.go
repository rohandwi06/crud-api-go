package config

// Sebuah struct yang berisi struktur dari .env
type Config struct {
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
}
