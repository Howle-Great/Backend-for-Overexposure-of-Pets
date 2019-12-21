package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	// "github.com/spf13/viper"

	_ "github.com/lib/pq"
)

type PostgrConnect struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func (con *PostgrConnect) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		con.host, con.port, con.user, con.password, con.dbname)
}

func main() {
	GetConfig("")

	dbConnection := PostgrConnect{
		"localhost",
		5432,
		"postgres",
		"example",
		"postgres",
	}

	db, err := sql.Open("postgres", dbConnection.ConnectionString())

	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	itemsRepo := ItemRepository{
		DB: db,
	}
	auth := AuthManager{}
	handler := &Handler{
		Items: itemsRepo,
		Auth:  auth,
	}

	r := mux.NewRouter()
	router := ConfigurateHandlers(r, handler)

	fmt.Printf("Сервер запущен на http://localhost:%v/\n", viper.Get("port"))
	http.ListenAndServe(":"+viper.GetString("port"), router)
}
