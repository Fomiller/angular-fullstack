package main

import (
	"fmt"
	"net/http"

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
	router := gin.Default()
	router.GET("/api/heroes", heroHandler)
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
	// io.WriteString(res, string(Heroes))
	// HJ, _ := json.Marshal(Heroes)
	c.JSON(http.StatusOK, Heroes)
	// res.Write(HJ)
}
