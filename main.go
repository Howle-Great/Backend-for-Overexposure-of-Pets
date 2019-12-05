package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	// "github.com/spf13/viper"
)

func main() {
	GetConfig("")
	r := mux.NewRouter()
	router := ConfigurateHandlers(r)
	fmt.Printf("Сервер запущен на http://localhost:%v/\n", viper.Get("port"))
	http.ListenAndServe(":"+viper.GetString("port"), router)
}
