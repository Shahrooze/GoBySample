package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var movies []Movie

func main() {
	movies = append(movies,
		Movie{
			Id:    "1",
			Isbn:  "TS@##$",
			Title: "ECO",
			Director: &Director{
				FirstName: "SHAHROOZ",
				LastName:  "Jafari",
			}},
	)
	movies = append(movies,
		Movie{
			Id:    "2",
			Isbn:  "DCV#$",
			Title: "RFC",
			Director: &Director{
				FirstName: "ALI",
				LastName:  "Jafari",
			}},
	)
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Starting server at port 8081 \n")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			//به خط پایین دقت کنید.
			//یک لیست جدید درست میکند
			//به وسیله
			//movies[:index]
			//از ایندکس صفر تا سر آیتمی که باید حذف شود را
			//انتخاب میکند.
			//همچنین به وسیله
			//movies[index+1:]
			//از بعد از آیتمی که باید حذف شود را تا انتهای لیست
			//درون movie
			//میریزد
			//...
			//باعث می شود یک کپی از کل لیست ایجاد شده درون مموری ایجاد شود/
			//چون
			//slice
			//رفرنس تایپ هست
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
