package configs

import (
	"strconv"
)

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func NewEmail(config Config) EmailConfig {
	port, _ := strconv.Atoi(config.Get("SMTP_PORT"))
	return EmailConfig{
		Host:     config.Get("SMTP_HOST"),
		Port:     port,
		Username: config.Get("SMTP_USER"),
		Password: config.Get("SMTP_PASS"),
		From:     config.Get("SMTP_FROM"),
	}
}
