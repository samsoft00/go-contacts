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

func generateRandId(n int) string {
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
	newContact.ID = generateRandId(10)
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


func main(){
	router := gin.Default()
	router.GET("/contacts", getContacts)
	router.GET("/contacts/:id", getContactByID)
	router.POST("/contacts", postContacts)

	port := os.Getenv("PORT")
	if port == "" { log.Fatal("$PORT must be set") }

	router.Run(":" + port)
}