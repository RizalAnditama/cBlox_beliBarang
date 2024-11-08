package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
Struct untuk menyimpan data baju.
harus diawali huruf besar, kalo ga malah dianggap private.
kalo di java mirip object
*/
type Baju struct {
	ID    int32  `json:"id"`
	Nama  string `json:"nama"`
	Jenis string `json:"jenis"`
	Warna string `json:"warna"`
	Merek string `json:"merek"`
}

// Dummy data slice baju
var shirts = []Baju{
	{ID: 1, Nama: "Kaos Top's Collection", Jenis: "Kaos", Warna: "Putih", Merek: "Uniqlo"},
	{ID: 2, Nama: "Kemeja Flanel", Jenis: "Kemeja", Warna: "Merah", Merek: "H&M"},
	{ID: 3, Nama: "Hoodie", Jenis: "Jaket", Warna: "Hitam", Merek: "Zara"},
}

// Fungsi untuk membeli baju berdasarkan ID yang dikirimkan via query atau form data
func beliBaju(c *gin.Context) {
	// Ambil ID dari query atau form
	idStr := c.DefaultPostForm("id", c.Query("id"))
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID baju tidak diberikan"})
		return
	}

	// Konversi ID menjadi integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID baju harus berupa angka"})
		return
	}

	// Cari baju berdasarkan ID
	for _, shirt := range shirts {
		if shirt.ID == int32(id) {
			c.JSON(http.StatusOK, shirt)
			return
		}
	}

	// Jika ID tidak ditemukan
	c.JSON(http.StatusNotFound, gin.H{"message": "Baju tidak ditemukan"})
}

func main() {
	router := gin.Default()

	// Endpoint POST untuk membeli baju
	router.POST("/beli_baju", beliBaju)

	// Jalankan server pada port 8080
	router.Run("localhost:8080")
}
