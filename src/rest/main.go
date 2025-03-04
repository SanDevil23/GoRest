package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// declaring the structure for the article
type Article struct {
	Id 		string `json:"id"`
	Title 	string `json:"title"`
	Desc 	string `json:"desc"`
	Content string `json:"content"`
}

// array of Article for storing articles
var articles []Article;

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

func handleRequests() {

	// handler
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/delete/{id}", deleteArticle).Methods("DELETE")

	// Set up CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow the frontend running on localhost:8080
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	

	// Apply CORS middleware to the router
	handler := c.Handler(myRouter)

	log.Fatal(http.ListenAndServe(":3000", handler))
	fmt.Println("Server successfully started on the port: 3000" )
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	// fmt.Fprintln(w, "Key: " + key)

	// Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, article := range articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)

	var art Article
	json.Unmarshal(reqBody, &art)
	// fmt.Fprintf(w, "%+v", string(reqBody))

	// update our global variable of articles
	// to include the new article

	articles = append(articles, art)
	json.NewEncoder(w).Encode(art)
}


func deleteArticle(w http.ResponseWriter, r *http.Request){
	// parsing the path parameters to read the incoming request
	vars := mux.Vars(r)

	// here we are extracting the 'id' of the article passed in the path
	id := vars["id"]

	for index, article := range articles {
		if article.Id == id {
			articles = append(articles[:index], articles[index+1:]...)
		}
	}	
	fmt.Println("The article is deleted", id)
	fmt.Fprintln(w, "The Article has been deleted successfully")
}

func main() {

	articles = []Article{
		Article{Id:"1", Title: "Hello", Desc: "Protcol Buffers", Content: "Article Content"},
		Article{Id:"2", Title: "Hello2", Desc: "Protcol Buffers2", Content: "Article Content2"},
	}
	handleRequests();
}