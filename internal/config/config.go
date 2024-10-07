package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)
const (
	EnvLocal = "local"
	EnvProd  = "prod"
)
type Config struct {
	Env           string         `yaml:"env" env:"ENV" env-default:"local"`
	SqlConnection SqlConnection  `yaml:"database"`
	HTTPConfig    HTTPConfig     `yaml:"http_server"`
	AuthConfig    AuthConfig     `yaml:"auth"`
	JWTConfig     JWTConfig      `yaml:"jwt"`
}

type SqlConnection struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type HTTPConfig struct {
	Host               string        `yaml:"host"`
	Port               string        `yaml:"port"`
	ReadTimeout        time.Duration `yaml:"read_timeout"`
	WriteTimeout       time.Duration `yaml:"write_timeout"`
	MaxHeaderMegabytes int           `yaml:"max_header_megabytes"`
}

type AuthConfig struct {
	JWT          JWTConfig `yaml:"jwt"`
	PasswordSalt string     `yaml:"password_salt"`
}

type JWTConfig struct {
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl"`
	Secret          string        `yaml:"secret"`
}


func MustLoad(filePath string) *Config { 
	if filePath == "" {
		log.Fatal("CONFIG_PATH not set")
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist\n", filePath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		log.Fatalf("cannot read config %s", filePath)
	}
	return &cfg
}
