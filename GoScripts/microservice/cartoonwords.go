package main

// libraries needed
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Structs needed

// Cartoon is a container for the Cartoon data to read in our JSON file
type Cartoon struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	DCDL        string `json:"DCDL_IDs"`
	Date        string `json:"Date"`
}

// JSONResponse defines the JSON response of our API
type JSONResponse struct {
	Status  string         `json:"status"`
	Request string         `json:"request"`
	Result  map[string]int `json:"dictionary"`
}

// Actual programme

// Global variable cartoonRepo is constructed with the readCartoons() function
var cartoonRepo = readCartoons()

// our main function, initialises the router
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dictionary/{id:[0-9]+}", GetWords).Methods("GET")
	router.HandleFunc("/dictionary", GetWordsExt).Methods("POST")
	log.Println("Listening at :8000" + "...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// init function for our "database"
func readCartoons() map[int]Cartoon {
	cartoonRepo := make(map[int]Cartoon)
	jsonFile, err := os.Open("cartoons.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	cartoonStruct := []Cartoon{}
	json.Unmarshal([]byte(byteValue), &cartoonStruct)
	for i := range cartoonStruct {
		cartoonRepo[cartoonStruct[i].ID] = cartoonStruct[i]
	}
	return cartoonRepo
}

// helping function that counts words (white-space separated tokens) in a given string
func findWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	reg, _ := regexp.Compile("[^a-zA-Z]+")
	m := make(map[string]int)
	for _, value := range words {
		key := reg.ReplaceAllString(value, "")
		if key != "" {
			m[key] = m[key] + 1
		}
	}
	return m
}

// GetWords returns a dictionary for the requested ID
func GetWords(w http.ResponseWriter, r *http.Request) {
	jsonresponse := JSONResponse{}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	cartoon, ok := cartoonRepo[id]
	if !ok {
		jsonresponse.Status = "exception: not found"
		jsonresponse.Request = "/dictionary/" + vars["id"]
		resultJSON, _ := json.Marshal(jsonresponse)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintln(w, string(resultJSON))
	} else {
		jsonresponse.Status = "ok"
		jsonresponse.Request = "/dictionary/" + vars["id"]
		jsonresponse.Result = findWords(cartoon.Description)
		resultJSON, _ := json.Marshal(jsonresponse)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintln(w, string(resultJSON))
	}
}

// GetWordsExt returns a dictionary for POST text
func GetWordsExt(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	jsonresponse := JSONResponse{}
	jsonresponse.Status = "ok"
	jsonresponse.Request = "/dictionary"
	jsonresponse.Result = findWords(text)
	resultJSON, _ := json.Marshal(jsonresponse)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintln(w, string(resultJSON))
}
