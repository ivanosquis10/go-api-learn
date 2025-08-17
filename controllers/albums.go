package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type album struct {
    ID     string  `json:"id" binding:"required"`
    Title  string  `json:"title" binding:"required"`
    Artist string  `json:"artist" binding:"required"`
    Price  float64 `json:"price" binding:"required"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func HandleGetParams(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El parametro name no pueder ser nulo",
		})
	}


	c.JSON(200, gin.H{
		name: name,
	})
}

func HandleGetAllAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func HandleCreateAlbum(c *gin.Context) {
	var newAlbum album

	 // Call BindJSON to bind the received JSON to newAlbum.
	 if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Values are required"})
		return
	 }

	albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func HandleGetOneAlbum(c *gin.Context) {
	idParam := c.Param("id")

	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El parametro 'id' no puede ser nulo o vacío",
		})

        return 
	}

	// Validación 2: Intentar convertir la cadena a un número entero.
	_, err := strconv.Atoi(idParam)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número entero válido"})
        return
	}

	// #1 Opcion
	for _, album := range albums {
		if album.ID == idParam {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
}