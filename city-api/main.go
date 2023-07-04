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
	city "github.com/thomas-mauran/city_api/struct"
	"github.com/thomas-mauran/city_api/utils"
)


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

        var listOfCities []city.City

        for rows.Next() {
            var city city.City
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
    if _, err := w.Write(jsonData); err != nil {
        log.Println("Error writing response:", err)
        // Handle the error. You can choose to log the error, send an appropriate response, or take any other action.
    }
})
    

  // City POST
    r.Post("/city", func(w http.ResponseWriter, r *http.Request) {
        var cityObj city.City
        err := json.NewDecoder(r.Body).Decode(&cityObj)
        if err != nil {
            log.Println("Error decoding JSON:", err)
            utils.Response(w, r, http.StatusInternalServerError, "Internal Server Error")
            return
        }

        sqlStatement := `INSERT INTO city (department_code, insee_code, zip_code, name, lat, lon) VALUES ($1, $2, $3, $4, $5, $6)`
        _, errQuery := db.Exec(sqlStatement, cityObj.DepartmentCode, cityObj.InseeCode, cityObj.ZipCode, cityObj.Name, cityObj.Lat, cityObj.Lon)
        if errQuery != nil {
            log.Println("Error executing SQL query:", errQuery)
            utils.Response(w, r, http.StatusInternalServerError, "Internal Server Error")
            return
        }

        utils.Response(w, r, http.StatusCreated, "Posted!")
    })

	errServ := http.ListenAndServe(":3000", r)
    if errServ != nil {
        log.Fatal(errServ)
    }
}
