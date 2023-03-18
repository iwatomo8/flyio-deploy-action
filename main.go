package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//　webページに移動
func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./public/home.html")
}

func initRouting(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		serveHome(c.Response(), c.Request())
		return nil
	})
}

func main() {
	fmt.Println("Start main func.")
	godotenv.Load()

	port := os.Getenv("PORT")
	fmt.Println("PORT: " + port)
	if port == "" {
		fmt.Println("$PORT must be set")
		port = "8080"
	}

	fmt.Println("Start echo.")
	e := echo.New()
	e.Use(middleware.CORS())

	initRouting(e)

	fmt.Println("End main func.")
	e.Logger.Fatal(e.Start(":" + port))
}
