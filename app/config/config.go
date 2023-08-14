package config

type DatabaseDetails struct {
	URL      string
	Username string
	Password string
}

type Config struct {
	Requests int
	Database DatabaseDetails
}
