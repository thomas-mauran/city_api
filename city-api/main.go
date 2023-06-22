package main

import (
    "database/sql"
    "log"
    "net/http"
    "os"
    _ "github.com/lib/pq"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/joho/godotenv" // Import the PostgreSQL driver package
    "github.com/thomas-mauran/city_api/utils"
)
func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    cityApiAddr := os.Getenv("CITY_API_ADDR")
    cityApiPort := os.Getenv("CITY_API_PORT")
    cityApiDBUrl := os.Getenv("CITY_API_DB_URL")
    cityApiDbUser := os.Getenv("CITY_API_DB_USER")
    cityApiDbPwd := os.Getenv("CITY_API_DB_PWD")
    if cityApiDBUrl == "" || cityApiDbUser == "" || cityApiDbPwd == "" {
        log.Fatal("Missing some environment variables")
    }

    // Connect to the database
    db, err := sql.Open("postgres", cityApiDBUrl)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    print(cityApiAddr)
    print(cityApiPort)

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/_health", func(w http.ResponseWriter, r *http.Request) {
        utils.Response(w, r, 204, "welcome")
        return
    })
    r.Get("/city", func(w http.ResponseWriter, r *http.Request) {
        utils.Response(w, r, 200, "City get")
        return
    })
    r.Post("/city", func(w http.ResponseWriter, r *http.Request) {
        utils.Response(w, r, 200, "City post")
        return
    })
    http.ListenAndServe(":3000", r)
}