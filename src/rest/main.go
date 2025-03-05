package main

import (
	"encoding/json"
	"fmt"
	"io"
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
		AllowedOrigins: []string{"*"}, 									
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
	reqBody, _ := io.ReadAll(r.Body)

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
		{Id: "1", Title: "Exploring the Universe", Desc: "A journey through the galaxy", Content: "In this article, we explore the vastness of space, from black holes to galaxies."},
		{Id: "2", Title: "The Future of Technology", Desc: "Innovations shaping tomorrow", Content: "This piece discusses emerging technologies that are set to change our world."},
		{Id: "3", Title: "Healthy Living Tips", Desc: "Ways to improve your health", Content: "Here are some practical tips for leading a healthier lifestyle."},
		{Id: "4", Title: "History of Art", Desc: "From the Renaissance to modern times", Content: "An overview of the major art movements and their impact on culture."},
		{Id: "5", Title: "Understanding Climate Change", Desc: "The science behind global warming", Content: "This article delves into the causes and effects of climate change."},
		{Id: "6", Title: "Travel Destinations", Desc: "Top places to visit this year", Content: "Explore some of the most beautiful travel destinations around the globe."},
		{Id: "7", Title: "Mastering Personal Finance", Desc: "Tips for managing your money", Content: "Learn how to budget, save, and invest wisely for a secure financial future."},
		{Id: "8", Title: "The Art of Cooking", Desc: "Culinary skills to impress", Content: "This article provides essential cooking techniques for home chefs."},
		{Id: "9", Title: "Fitness Trends", Desc: "What's hot in the fitness world", Content: "A look at the latest trends in health and fitness."},
		{Id: "10", Title: "The Power of Mindfulness", Desc: "Finding peace in a chaotic world", Content: "Explore how mindfulness can enhance your mental well-being."},
	}

	handleRequests();
}