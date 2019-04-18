package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func multirouter() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/casestudies", caseStudies).Methods("GET")
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	//r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	//r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func main() {
	r := multirouter()
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Welcome to Psychic Design studies</h1>")
}

func caseStudies (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1> My Case Studies</h1>")
}

//bird_handlers.go file added below here

/*type Bird struct {
	Species string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {

	bird := Bird{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	birds = append(birds, bird)

	http.Redirect(w, r, "/assets/", http.StatusFound)
}*/