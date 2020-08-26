package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Hero struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var Heroes []Hero

func init() {
	Heroes = []Hero{
		{Id: 11, Name: "Dr Nice"},
		{Id: 12, Name: "Narco"},
		{Id: 13, Name: "Bombasto"},
		{Id: 14, Name: "Celeritas"},
		{Id: 15, Name: "Magneta"},
		{Id: 16, Name: "RubberMan"},
		{Id: 17, Name: "Dynama"},
		{Id: 18, Name: "Dr IQ"},
		{Id: 19, Name: "Magma"},
		{Id: 20, Name: "Tornado"},
	}
	fmt.Println(Heroes)
}

func main() {
	// create gin router
	router := gin.Default()
	// serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./angular-tour-of-heroes/dist/angular-tour-of-heroes", true)))

	router.GET("/api/heroes", heroHandler)
	router.POST("/api/heroes", addHandler)
	router.PUT("/api/heroes", updateHandler)
	router.GET("/api/heroes/:id", detailHandler)
	// router.DELETE("/api/heroes/:id", deleteHandler)
	router.Run(":8080")
}

// return all heroes
func heroHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Heroes)
}

func addHandler(c *gin.Context) {
	body := c.Request.Body
	// create hero strucy to decode to
	newHero := Hero{}
	// decode body
	json.NewDecoder(body).Decode(&newHero)
	newHero.Id = len(Heroes) + 1
	fmt.Println("New hero", newHero)
	Heroes = append(Heroes, newHero)
	fmt.Println("all heros", Heroes)
	c.JSON(http.StatusOK, newHero)
}

// return a single hero
// could use index position?
func detailHandler(c *gin.Context) {
	p := c.Param("id")
	id, _ := strconv.Atoi(p)
	hero := Hero{}
	fmt.Println(id)
	for _, v := range Heroes {
		if id == v.Id {
			hero.Id = v.Id
			hero.Name = v.Name
		}
	}
	c.JSON(http.StatusOK, hero)
}

// // Delete handler
// func deleteHandler(c *gin.Context) {
// 	p := c.Param("id")
// 	matchId, _ := strconv.Atoi(p)
// 	for i, _ := range Heroes {
// 		if Heroes[i].Id == matchId {
// 			Heroes = RemoveIndex(Heroes, i)
// 		}
// 	}
// }

// update hero name
func updateHandler(c *gin.Context) {
	// parse request body
	body := c.Request.Body
	// create hero strucy to decode to
	hero := Hero{}
	// decode body
	json.NewDecoder(body).Decode(&hero)
	// replace name of hero that matches post hero
	for i, _ := range Heroes {
		if Heroes[i].Id == hero.Id {
			Heroes[i].Name = hero.Name
		}
	}
}

func RemoveIndex(s []Hero, index int) []Hero {
	return append(s[:index], s[index+1:]...)
}
