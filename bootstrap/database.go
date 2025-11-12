package bootstrap

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB inicializa la pool de conexiones de GORM a la BD
func InitDB(env *Env) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName)

	//Correr migraciones en /migrations
	autoMigrate(dsn, env)
	log.Println("Migrations run succesfully")

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), //Ajustar el log level segun lo necesario
	})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	//Obtener  *sql.DB para configurar la pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	//Configurar la pool
	sqlDB.SetMaxOpenConns(50)               //Numero maximo de conexiones abiertas
	sqlDB.SetMaxIdleConns(25)               //Conexiones idle que se mantienen en la pool
	sqlDB.SetConnMaxLifetime(1 * time.Hour) //Reciclar conexiones cada una hora

	//Verificar conexiones
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("GORM database connection pool initialized")
	}

}
