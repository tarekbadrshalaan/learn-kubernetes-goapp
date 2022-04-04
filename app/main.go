package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("POSTGRES_HOST")     // "goapp-database-service"
	port     = os.Getenv("POSTGRES_PORT")     // 5432
	user     = os.Getenv("POSTGRES_USER")     // "postgres"
	password = os.Getenv("POSTGRES_PASSWORD") // "password"
	dbname   = os.Getenv("POSTGRES_DB")       // "goapp"
)

func main() {
	http.HandleFunc("/", handlermain)
	http.HandleFunc("/worker", handlerId)
	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe(":8080", nil)

}

func handlermain(w http.ResponseWriter, r *http.Request) {
	log.Println("log from Hello World")
	fmt.Fprintf(w, "Hello World!")
}

// http://localhost:8080/worker?id=123
func handlerId(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect to Database")
	defer db.Close()

	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		fmt.Fprintf(w, "Url Param 'id' is missing!")
		return
	}
	key := keys[0]
	sqlStatement := fmt.Sprintf(`INSERT INTO workers (worker) VALUES (%v)`, key)
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "id = %v", key)
}
