package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	BrokerAddress                string `envconfig:"BROKER_ADDRESS" default:"amqp://127.0.0.1:5672"`
	PostmarkApiUrl               string `envconfig:"POSTMARK_API_URL" default:"https://api.postmarkapp.com/email/withTemplate"`
	PostmarkApiToken             string `envconfig:"POSTMARK_API_TOKEN" required:"true"`
	PostmarkEmailSender          string `envconfig:"POSTMARK_EMAIL_SENDER" required:"true"`
	PostmarkTemplateConfirmEmail string `envconfig:"POSTMARK_TEMPLATE_CONFIRM_EMAIL" required:"true"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)

	return cfg, err
}
