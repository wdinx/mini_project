package config

type Config struct {
	Database           Database
	Midtrans           Midtrans
	DigitalOceanSpaces DigitalOceanSpaces
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

type DigitalOceanSpaces struct {
	AccessToken string
	SecretKey   string
	Region      string
	Name        string
	Endpoint    string
}
