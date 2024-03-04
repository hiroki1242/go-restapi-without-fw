package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Movie struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Release int    `json:"release"`
}

var movies = []Movie{
	{1, "Avengers", 2012},
	{2, "Harry Potter", 2001},
	{3, "Star Wars", 1978},
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	movieId, err := strconv.Atoi(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	for _, movie := range movies {
		if movieId == movie.Id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	http.NotFound(w, r)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie Movie

	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// idが重複しないように修正
	for _, movie := range movies {
		if movie.Id == newMovie.Id {
			http.Error(w, "Id already exists.", http.StatusBadRequest)
			return
		}
	}

	movies = append(movies, newMovie)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	movieId, err := strconv.Atoi(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	var movieInfo Movie

	// bodyのjsonをデコード
	err = json.NewDecoder(r.Body).Decode(&movieInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// jsonのidとパスパラメーターのidの整合性を確認
	if movieId != movieInfo.Id {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	for i, movie := range movies {
		if movieId == movie.Id {
			movies[i] = movieInfo
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

	http.NotFound(w, r)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	movieId, err := strconv.Atoi(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	for i, movie := range movies {
		if movieId == movie.Id {
			movies = append(movies[:i], movies[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

	http.NotFound(w, r)
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getMovies(w, r)
	case "POST":
		createMovie(w, r)
	default:
		errMsg := fmt.Sprintf("%s is not supported HTTP Method", r.Method)
		http.Error(w, errMsg, http.StatusMethodNotAllowed)
	}
}

func movieHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getMovie(w, r)
	case "PUT":
		updateMovie(w, r)
	case "DELETE":
		deleteMovie(w, r)
	default:
		errMsg := fmt.Sprintf("%s is not supported HTTP Method", r.Method)
		http.Error(w, errMsg, http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/movies", moviesHandler)
	http.Handle("/movies/", http.StripPrefix("/movies/", http.HandlerFunc(movieHandler)))
	http.ListenAndServe(":8080", nil)
}
