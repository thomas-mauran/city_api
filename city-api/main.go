package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/thomas-mauran/city_api/utils"
)

type City struct {
    ID             int     `json:"id"`
    DepartmentCode string  `json:"department_code"`
    InseeCode      string  `json:"insee_code"`
    ZipCode        string  `json:"zip_code"`
    Name           string  `json:"name"`
    Lat            float64 `json:"lat"`
    Lon            float64 `json:"lon"`
}

func main() {
	cityApiAddr := os.Getenv("CITY_API_ADDR")
	cityApiPort := os.Getenv("CITY_API_PORT")
	cityApiDBUrl := os.Getenv("CITY_API_DB_URL")
	cityApiDbUser := os.Getenv("CITY_API_DB_USER")
	cityApiDbPwd := os.Getenv("CITY_API_DB_PWD")

    print("cityApiAddr:", cityApiAddr)
    print("cityApiPort:", cityApiPort)

	if cityApiDBUrl == "" || cityApiDbUser == "" || cityApiDbPwd == "" {
		log.Fatal("Missing some environment variables")
	}

	// Connect to the database
	db, err := sql.Open("postgres", cityApiDBUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

    // Health check
	r.Get("/_health", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			log.Println("Unable to ping the database:", err)
			response := fmt.Sprintf("Unable to ping the database: %v", err)
			utils.Response(w, r, 204, response)
			return
		}
		log.Println("Everything is good!")
		utils.Response(w, r, 204, "Everything is good!")
		return
	})


    // City GET
    r.Get("/city", func(w http.ResponseWriter, r *http.Request) {
        sqlStatement := `SELECT * FROM city`
        rows, err := db.Query(sqlStatement)
        if err != nil {
            log.Println("Error querying the database:", err)
            utils.Response(w, r, http.StatusInternalServerError, "Internal Server Error")
            return
        }
        defer rows.Close()

        var listOfCities []City

        for rows.Next() {
            var city City
            if err := rows.Scan(&city.ID, &city.DepartmentCode, &city.InseeCode, &city.ZipCode, &city.Name, &city.Lat, &city.Lon); err != nil {
                log.Println("Error scanning row:", err)
                utils.Response(w, r, http.StatusInternalServerError, "Internal Server Error")
                return
            }
            listOfCities = append(listOfCities, city)
        }
        if err := rows.Err(); err != nil {
            log.Println("Error iterating rows:", err)
            utils.Response(w, r, http.StatusInternalServerError, "Internal Server Error")
            return
        }

        jsonData, err := json.Marshal(listOfCities)
        if err != nil {
            log.Println("Error marshaling JSON:", err)
            utils.Response(w, r, http.StatusInternalServerError, "Internal Server Error")
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(jsonData)
    })

    

    // City POST
	r.Post("/city", func(w http.ResponseWriter, r *http.Request) {
		utils.Response(w, r, 200, "City post")
		return
	})
	errServ := http.ListenAndServe(":3000", r)
    if errServ != nil {
        log.Fatal(errServ)
    }
}
