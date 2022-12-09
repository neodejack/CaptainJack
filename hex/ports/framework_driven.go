package ports

type SecretsPorts interface {
	GetTeleToken() string
	GetWebAppUrl() string
}
