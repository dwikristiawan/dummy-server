package config

type Jwt struct {
	SecretKey    string `envconfig:"JWT_SECRETE_KEY" required:"true"`
	RefreshKey   string `envconfig:"JWT_REFRESH_KEY" required:"true"`
	Expiration   string `envconfig:"JWT_EXPIRATION" required:"true"`
	ReExpiration string `envconfig:"JWT_RE_EXPIRATION" required:"true"`
}
