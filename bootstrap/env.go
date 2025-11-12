package bootstrap

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	MigrationsFolder       string `mapstructure:"MIGRATIONS_FOLDER"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Can't find file .env, using OS env instead.")
		return NewEnvWithoutFile(env)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}

func NewEnvWithoutFile(env Env) *Env {
	// Asignar valores desde las variables de entorno
	env.DBHost = getEnv("DB_HOST", "localhost")
	env.DBName = getEnv("DB_NAME", "")
	env.DBPass = getEnv("DB_PASS", "")
	env.DBPort = getEnv("DB_PORT", "5432")
	env.DBUser = getEnv("DB_USER", "")
	env.AccessTokenSecret = getEnv("ACCESS_TOKEN_SECRET", "")
	env.ServerAddress = getEnv("SERVER_ADDRESS", ":8080")
	env.MigrationsFolder = getEnv("MIGRATIONS_FOLDER", "./migrations")

	return &env
}

// Función auxiliar para obtener variables de entorno con un valor por defecto
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
