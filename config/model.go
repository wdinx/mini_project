package config

type Config struct {
	Database Database
	Midtrans Midtrans
}

type Database struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

type Midtrans struct {
	Key    string
	IsProd bool
}
