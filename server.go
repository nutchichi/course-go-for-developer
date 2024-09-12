package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nutchichi/course-go-for-developer/flight"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func initDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
	)

	db, err := sql.Open(viper.GetString("db.driver"), psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	initTimeZone()
	initConfig()

	db := initDatabase()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection database successfully")

	fr := flight.NewFlightRepositoryDB(db)
	fs := flight.NewFlightService(fr)
	fh := flight.NewFlightHandler(fs)

	r := gin.Default()

	r.GET("/ping", fh.PingHandler)
	r.GET("/flights/:id", fh.GetFlightByIDHandler)
	r.GET("/flights", fh.GetFlightsHandler)
	r.POST("/flights/create", fh.CreateFlightHandler)
	r.PUT("/flights/:id", fh.UpdateFlightHandler)
	r.DELETE("/flights/:id", fh.DeleteFlightHandler)
	r.Run(fmt.Sprintf(":%v", viper.GetInt("app.port")))
}
