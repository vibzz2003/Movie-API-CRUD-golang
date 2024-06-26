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

type Movie struct { //creating struct types for defining the attributes associated to the struct class
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct { //same for the director struct
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie //slice of type movie

func getMovies(w http.ResponseWriter, r *http.Request) { //get movies function w->responsewriter and r->pointer pointing to the http request by the user
	w.Header().Set("Content-Type", "application/json") //setting the content type into json format
	json.NewEncoder(w).Encode(movies)                  //encoding it into type movies that we have created
}

func deleteMovie(w http.ResponseWriter, r *http.Request) { //delet movies function
	w.Header().Set("Content-Type", "application/json") //setting the response writing in json format
	params := mux.Vars(r)                              //getting the value from the request parameter
	for index, item := range movies {                  //looping into movies
		if item.ID == params["id"] { //when parameter matches the required id
			movies = append(movies[:index], movies[index+1:]...) //appending all the movies after that to its place thus deleting it
			break
		}
	}
	json.NewEncoder(w).Encode(movies) //listing all the left movies
}

func getMovie(w http.ResponseWriter, r *http.Request) { //get movie by specific movie id
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) { //creating movie
	w.Header().Set("Content-Type", "application/json") //setting response writing into json format
	var movie Movie                                    //creating a var of type struct movie
	_ = json.NewDecoder(r.Body).Decode(&movie)         //decoding the request body into the type movie struct
	movie.ID = strconv.Itoa(rand.Intn(100000000))      //assigning a new number from random import
	movies = append(movies, movie)                     //appending the movie to the slice movie
	json.NewEncoder(w).Encode(movie)                   //showing thr movies
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//we will set the json content type
	//access to the params
	//loop over movies for finding the movie id
	//delete the resultant movie
	//add new movie in its place and its id in the body of the postman
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "2233456", Title: "Spiderman", Director: &Director{FirstName: "Sam", LastName: "Raimi"}})
	movies = append(movies, Movie{ID: "2", Isbn: "2233457", Title: "GOTG2", Director: &Director{FirstName: "James", LastName: "Gunn"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port: 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
