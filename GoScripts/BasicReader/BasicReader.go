package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	// "encoding/csv"
	// "encoding/gob"
	// "github.com/boltdb/bolt"
	"github.com/ThomasK81/gocite"
	"github.com/gorilla/mux"
)

// 1. server-config
// It's always better to give people to change stuff. Instead of hard-coding the port etc., parse a config file into this struct
// The config file could be JSON or XML or anything else really, but those two formats I recommend
type serverConfig struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	Source string `json:"cex_source"`
}

// reading in the info from the config.json into a serverConfig object
var confvar = loadConfiguration("./config.json")

func loadConfiguration(file string) serverConfig {
	var config serverConfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

// 2. HTML templating
// Page stores all the information to be fed into our HTML GUI
type Page struct {
	ID   string
	Text template.HTML
}

// and here are somne templates, check them out
var templates = template.Must(template.ParseFiles("templates/view.html"))

// we need a function that put the data from the Page object into the actual HTML page
func renderTemplate(w http.ResponseWriter, tmpl string, p Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// we can test it here

func testTemplate(w http.ResponseWriter, r *http.Request) {
	testpage := Page{
		ID:   "urn:cts:tests:test1.test1:1",
		Text: template.HTML("This is a testing of the template")}

	renderTemplate(w, "view", testpage)
}

// we want to be able to have a corpus
// a corpus is a collection of documents / works
// we can use the gocite library to mimic a backend for this
// but firdst we need to be able to simply read from the interweb
// getContent does this
func getContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

var backend = []gocite.Work{}

func loaddata() {
	corpus := []gocite.Work{}
	data, _ := getContent(confvar.Source)
	str := string(data)
	// find the section that contains the textdata and split the string so it only contains that section
	ctsdata := strings.Split(str, "#!ctsdata")[1]
	ctsdata = strings.Split(ctsdata, "#!")[0]
	// fix multiline errors in the data
	re := regexp.MustCompile("(?m)[\r\n]*^//.*$")
	ctsdata = re.ReplaceAllString(ctsdata, "")

	reader := csv.NewReader(strings.NewReader(ctsdata))
	reader.Comma = '#'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	var texturns, text []string

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println(line)
			log.Fatal(error)
		}
		switch {
		// case to be expected
		case len(line) == 2:
			texturns = append(texturns, line[0])
			text = append(text, line[1])
		// in case the lazy quote does not work properly
		case len(line) > 2:
			texturns = append(texturns, line[0])
			var textstring string
			for j := 1; j < len(line); j++ {
				textstring = textstring + line[j]
			}
			text = append(text, textstring)
		// fatal data error
		case len(line) < 2:
			log.Fatal("Wrong line:", line)
		}
	}

	workid := ""
	newwork := gocite.Work{}
	count := 0
	for i, v := range texturns {
		texturn := gocite.SplitCTS(v)
		workurn := strings.Join([]string{texturn.Base, texturn.Protocol, texturn.Namespace, texturn.Work}, ":")
		if workurn != workid {
			if workid == "" {
				workid = workurn
				newwork.WorkID = workurn
				newPassage := gocite.Passage{PassageID: workurn + ":" + texturn.Passage, Text: gocite.EncText{CEX: text[i]}, Index: count}
				newwork.Passages = append(newwork.Passages, newPassage)
			} else {
				workid = workurn
				corpus = append(corpus, newwork)
				newwork = gocite.Work{}
				newwork.WorkID = workurn
				count = 0
				newPassage := gocite.Passage{PassageID: workurn + ":" + texturn.Passage, Text: gocite.EncText{CEX: text[i]}, Index: count}
				newwork.Passages = append(newwork.Passages, newPassage)
			}
		} else {
			newPassage := gocite.Passage{PassageID: workurn + ":" + texturn.Passage, Text: gocite.EncText{CEX: text[i]}, Index: count}
			newwork.Passages = append(newwork.Passages, newPassage)
		}
		count = count + 1
	}
	corpus = append(corpus, newwork)
	log.Println("corpus successfully read")
	backend = corpus
	return
}

func main() {
	router := mux.NewRouter()
	// first thing we are going to do is to open a folder for external reading, so clients can read static files from there
	cexHandler := http.StripPrefix("/cex/", http.FileServer(http.Dir("./cex/")))
	router.PathPrefix("/cex/").Handler(cexHandler)
	// next, we build and test a HTML Template
	router.HandleFunc("/testTemplate", testTemplate)
	// build the backend
	router.HandleFunc("/loadDB", loadDB)
	// now let's built an actual basic text interface
	router.HandleFunc("/passageView/{urn}", passageView)
	// with which services could you improve this?
	log.Println("Listening at" + confvar.Port + "...")
	log.Fatal(http.ListenAndServe(confvar.Port, router))
}

func loadDB(w http.ResponseWriter, r *http.Request) {
	loaddata()
	io.WriteString(w, "Success")
}

func passageView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v := vars["urn"]
	texturn := gocite.SplitCTS(v)
	workurn := strings.Join([]string{texturn.Base, texturn.Protocol, texturn.Namespace, texturn.Work}, ":")
	found := false
	passage := gocite.Passage{}
	for i := range backend {
		if backend[i].WorkID == workurn {
			for j := range backend[i].Passages {
				if backend[i].Passages[j].PassageID == v {
					passage = backend[i].Passages[j]
					found = true
					break
				}
			}
		}
	}
	if !found {
		io.WriteString(w, "Passage not found")
	} else {
		page := Page{
			ID:   passage.PassageID,
			Text: template.HTML(passage.Text.CEX)}
		renderTemplate(w, "view", page)
	}
}
