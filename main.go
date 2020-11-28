package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	uh "TestKriya/users/handler"
	ur "TestKriya/users/repository"
	us "TestKriya/users/service"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("sslmode", "disable")
	connStr := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := gorm.Open("postgres", connStr)
	if err != nil {
		logrus.Fatal(err)
	}

	err = dbConn.DB().Ping()
	if err != nil {
		logrus.Error(err)
	}

	defer func() {
		err = dbConn.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	//dbConn.Debug().AutoMigrate(
	//	//&model.Users{},&model.Role{},
	//)

	router := mux.NewRouter().StrictSlash(true)

	userRepo := ur.CreateRepoImpl(dbConn)
	userService := us.CreateUserServiceImpl(userRepo)
	uh.CreateUserHandler(router, userService)

	fmt.Println("Starting web server at port : 8082")
	err = http.ListenAndServe(": "+"8082", router)
	if err != nil {
		logrus.Fatal(err)
	}
}
