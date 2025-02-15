package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int32  `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type StokBaju struct {
	ID    int32   `json:"id" gorm:"primaryKey"`
	Nama  string  `json:"nama"`
	Jenis string  `json:"jenis"`
	Warna string  `json:"warna"`
	Merek string  `json:"merek"`
	Size  string  `json:"size"`
	Harga float64 `json:"harga"`
}
type Keranjang struct {
	ID     int32 `json:"id" gorm:"primaryKey"`
	UserID int32 `json:"id_user"`
	BajuID int32 `json:"id_baju"`
	Jumlah int32 `json:"jumlah"`
}
type Order struct {
	ID     int32 `json:"id" gorm:"primaryKey"`
	UserID int32 `json:"user_id"`
	Total  int32 `json:"total"`
}

var db *gorm.DB
var err error

func initDB() {
	dbName := "sddp_beli_baju"
	dsn := fmt.Sprintf("root:@tcp(localhost:3306)/%s", dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	users := []User{
		{ID: 1, Username: "rizalandit", Password: "sofan"},
		{ID: 2, Username: "root", Password: "root"},
	}

	stokBaju := []StokBaju{
		{ID: 1, Nama: "Kaos Top's Collection", Jenis: "Kaos", Warna: "Putih", Merek: "Uniqlo", Size: "M", Harga: 120000},
		{ID: 2, Nama: "Kemeja Flanel", Jenis: "Kemeja", Warna: "Merah", Merek: "H&M", Size: "L", Harga: 150000},
		{ID: 3, Nama: "Hoodie", Jenis: "Jaket", Warna: "Hitam", Merek: "Zara", Size: "XL", Harga: 300000},
		{ID: 4, Nama: "Sweater Rajut", Jenis: "Sweater", Warna: "Biru", Merek: "Pull&Bear", Size: "M", Harga: 200000},
		{ID: 5, Nama: "Jaket Bomber", Jenis: "Jaket", Warna: "Hijau", Merek: "Levi's", Size: "L", Harga: 250000},
		{ID: 6, Nama: "Kaos Polo", Jenis: "Kaos", Warna: "Abu-abu", Merek: "Ralph Lauren", Size: "S", Harga: 180000},
		{ID: 7, Nama: "Kemeja Denim", Jenis: "Kemeja", Warna: "Biru Tua", Merek: "Levi's", Size: "M", Harga: 230000},
		{ID: 8, Nama: "Turtleneck", Jenis: "Sweater", Warna: "Putih", Merek: "Zara", Size: "L", Harga: 170000},
		{ID: 9, Nama: "Tank Top", Jenis: "Kaos", Warna: "Hitam", Merek: "Adidas", Size: "S", Harga: 90000},
		{ID: 10, Nama: "Jaket Kulit", Jenis: "Jaket", Warna: "Cokelat", Merek: "Zara", Size: "XL", Harga: 500000},
		{ID: 11, Nama: "Hoodie Zipper", Jenis: "Jaket", Warna: "Merah", Merek: "Nike", Size: "L", Harga: 280000},
		{ID: 12, Nama: "Sweatshirt", Jenis: "Sweater", Warna: "Kuning", Merek: "Uniqlo", Size: "M", Harga: 150000},
		{ID: 13, Nama: "Kaos Lengan Panjang", Jenis: "Kaos", Warna: "Biru Muda", Merek: "Uniqlo", Size: "S", Harga: 130000},
		{ID: 14, Nama: "Jaket Jeans", Jenis: "Jaket", Warna: "Biru", Merek: "Levi's", Size: "L", Harga: 270000},
		{ID: 15, Nama: "Vest", Jenis: "Rompi", Warna: "Cokelat", Merek: "H&M", Size: "M", Harga: 110000},
		{ID: 16, Nama: "Kaos V-Neck", Jenis: "Kaos", Warna: "Hitam", Merek: "Uniqlo", Size: "L", Harga: 90000},
		{ID: 17, Nama: "Kemeja Batik", Jenis: "Kemeja", Warna: "Biru", Merek: "Danar Hadi", Size: "XL", Harga: 220000},
		{ID: 18, Nama: "Sweater Hoodie", Jenis: "Sweater", Warna: "Merah Maroon", Merek: "Pull&Bear", Size: "M", Harga: 190000},
		{ID: 19, Nama: "Cardigan", Jenis: "Sweater", Warna: "Abu-abu", Merek: "Zara", Size: "L", Harga: 210000},
		{ID: 20, Nama: "Kaos Raglan", Jenis: "Kaos", Warna: "Putih dan Hitam", Merek: "H&M", Size: "S", Harga: 95000},
	}

	db.AutoMigrate(&Keranjang{}, &Order{}, &StokBaju{}, &User{})

	for _, baju := range stokBaju {
		db.FirstOrCreate(&baju, StokBaju{ID: baju.ID})
	}

	for _, user := range users {
		db.FirstOrCreate(&user, User{ID: user.ID})
	}
}

func authenticateUser(c *gin.Context) {
	// Get the Authorization header

	authHeader := c.GetHeader("Authorization")

	if !strings.HasPrefix(authHeader, "Basic ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		c.Abort()
		return
	}

	// Decode the base64 encoded username:password
	payload, err := base64.StdEncoding.DecodeString(authHeader[6:])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
		c.Abort()
		return
	}

	// Split username and password
	userPass := strings.SplitN(string(payload), ":", 2)
	if len(userPass) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
		c.Abort()
		return
	}

	username := userPass[0]
	password := userPass[1]
	// Check credentials in the database
	var user User
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		c.Abort()
		return
	}

	// Store user ID in the context for further use
	c.Set("userID", user.ID)
	c.Next()
}

func getStokBaju(c *gin.Context) {
	var stokBaju []StokBaju
	db.Find(&stokBaju)
	c.JSON(http.StatusOK, stokBaju)
}

func getBajuByID(c *gin.Context) {
	id := c.Param("id")
	var baju StokBaju
	if err := db.First(&baju, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, baju)
}

func addBajuToKeranjang(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID not found"})
	}
	var keranjang Keranjang
	if err := c.ShouldBindJSON(&keranjang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	keranjang.UserID = int32(userID.(int32))
	db.Create(&keranjang)
	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func getKeranjang(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID not found"})
	}
	var keranjangs []Keranjang
	db.Where("user_id = ?", userID).Find(&keranjangs)
	c.JSON(http.StatusOK, keranjangs)
}

func checkoutKeranjang(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID not found"})
	}
	var keranjangs []Keranjang
	db.Where("user_id = ?", userID).Find(&keranjangs)

	// Calculate total
	var total int32
	for _, keranjang := range keranjangs {
		var stokBaju StokBaju
		db.First(&stokBaju, keranjang.BajuID)
		total += int32(stokBaju.Harga) * keranjang.Jumlah
	}

	// Create Order
	order := Order{UserID: int32(userID.(int32)), Total: total}
	db.Create(&order)

	// Clear Keranjang
	db.Where("user_id = ?", userID).Delete(&Keranjang{})

	c.JSON(http.StatusOK, gin.H{"message": "Checkout successful", "order_id": order.ID})
}

func getOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID not found"})
	}
	var orders []Order
	db.Where("user_id = ?", userID).Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func getOrderByID(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID not found"})
	}
	id := c.Param("id")

	var order Order
	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func filterBaju(c *gin.Context) {
    var stokBaju []StokBaju

    // Get filter parameters from query
    jenis := c.Query("jenis")
    warna := c.Query("warna")
    merek := c.Query("merek")
    minHarga := c.Query("min_harga")
    maxHarga := c.Query("max_harga")

    // Build query with conditions
    query := db.Model(&StokBaju{})
    if jenis != "" {
        query = query.Where("jenis = ?", jenis)
    }
    if warna != "" {
        query = query.Where("warna = ?", warna)
    }
    if merek != "" {
        query = query.Where("merek = ?", merek)
    }
    if minHarga != "" {
        query = query.Where("harga >= ?", minHarga)
    }
    if maxHarga != "" {
        query = query.Where("harga <= ?", maxHarga)
    }

    // Execute query
    if err := query.Find(&stokBaju).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
        return
    }

    // Return results
    c.JSON(http.StatusOK, stokBaju)
}


func main() {
	initDB()
	router := gin.Default()

	// Item routes
	router.GET("/baju", getStokBaju)       //
	router.GET("/baju/:id", getBajuByID)   //
	router.GET("/baju/filter", filterBaju) //

	// Authenticated routes
	authRoutes := router.Group("/")
	authRoutes.Use(authenticateUser)
	{
		authRoutes.POST("/keranjang", addBajuToKeranjang) //
		authRoutes.GET("/keranjang", getKeranjang)        //
		authRoutes.POST("/checkout", checkoutKeranjang)   //
		authRoutes.GET("/pesanan", getOrders)             //
		authRoutes.GET("/pesanan/:id", getOrderByID)      //
	}

	router.Run(":8080")

}
