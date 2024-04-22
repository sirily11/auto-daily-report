package config

type RunMode string

const (
	ModeDevelopment RunMode = "development"
	ModeProduction  RunMode = "production"
)

type Config struct {
	Run Run `mapstructure:"run"`
}

type Run struct {
	Mode RunMode
}
