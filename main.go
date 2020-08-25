package main

import (
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
	// router.PUT("/api/heroes", updateHandler)
	router.GET("/api/heroes/:id", detailHandler)
	router.Run(":8080")
}

func heroHandler(c *gin.Context) {
	// io.WriteString(res, string(Heroes))
	// HJ, _ := json.Marshal(Heroes)
	c.JSON(http.StatusOK, Heroes)

	// res.Write(HJ)
}
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
	// res.Write(HJ)
}

// func updateHandler(c *gin.Context) {
// 	body := c.Request.Body
// 	hero := Hero{}
// 	json.NewDecoder(body).Decode(&hero)
// 	i := funk.Get(Heroes, hero.Id)
// 	fmt.Println(i)
// 	fmt.Println(hero)
// }
