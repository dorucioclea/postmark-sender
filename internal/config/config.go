package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	BrokerAddress    string `envconfig:"BROKER_ADDRESS" default:"amqp://127.0.0.1:5672"`
	PostmarkApiUrl   string `envconfig:"POSTMARK_API_URL" default:"https://api.postmarkapp.com/email/withTemplate"`
	PostmarkApiToken string `envconfig:"POSTMARK_API_TOKEN" required:"true"`

	PostmarkEmailFrom            string `envconfig:"POSTMARK_EMAIL_FROM" required:"true"`
	PostmarkEmailCc              string `envconfig:"POSTMARK_EMAIL_CC" default:""`
	PostmarkEmailBcc             string `envconfig:"POSTMARK_EMAIL_BCC" default:""`
	PostmarkEmailTrackOpens      bool   `envconfig:"POSTMARK_EMAIL_TRACK_OPENS" default:"false"`
	PostmarkEmailTrackTrackLinks string `envconfig:"POSTMARK_EMAIL_TRACK_LINKS" default:""`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)

	return cfg, err
}
