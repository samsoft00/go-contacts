package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type contact struct {
	ID				string		`json:"id"`
	FullName		string		`json:"fullname"`
	Email			string		`json:"email"`
	Subject			string		`json:"subject"`
	Message			string		`json:"message"`
	Createddate		string		`json:"created_at"`
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var contacts = []contact {}

// cor middleware
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
			return
        }
        c.Next()
    }
}

func generateRandID(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)	
}

//getContacts responds with the list of all albums as JSON
func getContacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contacts)
}

// postContacts adds an album from JSON received in the request body
func postContacts(x *gin.Context) {
	var newContact contact
	newContact.ID = generateRandID(10)
	newContact.Createddate = time.Now().String()

	if err := x.BindJSON(&newContact); err != nil {
		return
	}

	contacts = append(contacts, newContact)
	x.IndentedJSON(http.StatusCreated, newContact)
}

func getContactByID(c *gin.Context) {
	id := c.Param("id")

	//loop over the list of albums, looking for an album whose ID value matches the parameter
	for _, a := range contacts {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func printName(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"name": "Oyewole Abayomi Samuel",
		"role": "Software Engineer Manager",
		"Phone Number": "2347063317344",
		"email": "oyewoleabayomi@gmail.com",
		"Location": "127.0.0.1",
	})
}


func main(){
	router := gin.Default()
	router.Use(CORSMiddleware())
	// this print a name
	router.GET("/", printName)
	router.GET("/contacts", getContacts)
	router.GET("/contacts/:id", getContactByID)
	router.POST("/contacts", postContacts)

	port := os.Getenv("PORT")
	if port == "" { log.Fatal("$PORT must be set") }

	//nolint:errcheck //CODEI8:
	router.Run(":" + port)
}