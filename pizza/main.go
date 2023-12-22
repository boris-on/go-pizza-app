package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	connStr := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("SLLMODE"),
	)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	ID      int    `json:"id"`
	Login   string `json:"login"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
	Address string `json:"address"`
}

type UpdateUser struct {
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
	Address string `json:"address"`
}

type Dish struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Price       int         `json:"price"`
	Description string      `json:"description"`
	ImageURL    string      `json:"image_url"`
	Tags        interface{} `json:"tags"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Phone    int    `json:"phone"`
		Address  string `json:"address"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO user_info (login, password, email, phone, address) VALUES ($1, $2, $3, $4, $5)",
		user.Login, user.Password, user.Email, user.Phone, user.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	err = db.QueryRow("SELECT id, login, email, phone, address FROM user_info WHERE login = $1 AND password = $2",
		credentials.Login, credentials.Password).Scan(&user.ID, &user.Login, &user.Email, &user.Phone, &user.Address)
	if err != nil {
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	var query struct {
		Login string `json:"login"`
	}

	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	err = db.QueryRow("SELECT id, login, email, phone, address FROM user_info WHERE login = $1", query.Login).Scan(&user.ID, &user.Login, &user.Email, &user.Phone, &user.Address)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	jsonResponse(w, user)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	var update struct {
		Login   string `json:"login"`
		Email   string `json:"email"`
		Phone   int    `json:"phone"`
		Address string `json:"address"`
	}

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = updateUser(update.Login, UpdateUser{
		Email:   update.Email,
		Phone:   update.Phone,
		Address: update.Address,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getAllDishesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, price, description, image_url, tags FROM dish")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var dishes []Dish
	for rows.Next() {
		var dish Dish
		err := rows.Scan(&dish.ID, &dish.Name, &dish.Price, &dish.Description, &dish.ImageURL, &dish.Tags)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tagsBytes, ok := dish.Tags.([]byte)
		if ok {
			var tags []string
			err := json.Unmarshal(tagsBytes, &tags)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dish.Tags = tags
		}

		dishes = append(dishes, dish)
	}

	response := map[string]interface{}{"items": dishes}
	jsonResponse(w, response)
}

func updateUser(login string, update UpdateUser) error {
	_, err := db.Exec(
		"UPDATE user_info SET email=$1, phone=$2, address=$3 WHERE login=$4",
		update.Email, update.Phone, update.Address, login,
	)
	return err
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	initDB()

	r := mux.NewRouter()

	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/user/get", getUserHandler).Methods("POST")
	r.HandleFunc("/user/set", updateUserHandler).Methods("POST")
	r.HandleFunc("/getAllDishes", getAllDishesHandler).Methods("GET")

	portStr := os.Getenv("SERVER_CONTAINER_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Server listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
