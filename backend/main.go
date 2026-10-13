package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ==========================================
// 1. STRUKTUR DATABASE (MODELS)
// ==========================================

// Tabel Whitelist Email (Daftar Klien VIP dari Tim Kamu)
type Whitelist struct {
	gorm.Model
	Email          string `gorm:"unique;not null" json:"email"`
	NamaPerusahaan string `json:"nama_perusahaan"`
	IsUsed         bool   `gorm:"default:false" json:"is_used"`
}

// Tabel Perusahaan (Tenant)
type Tenant struct {
	gorm.Model
	NamaPerusahaan string `json:"nama_perusahaan"`
	LogoUrl        string `json:"logo_url"`
	Alamat         string `json:"alamat"`
	Kontak         string `json:"kontak"`
}

// Tabel Pengguna
type User struct {
	gorm.Model
	TenantID uint   `json:"tenant_id"`
	Nama     string `json:"nama"`
	Email    string `gorm:"unique;not null" json:"email"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"` // "superadmin", "owner", atau "admin"
	Telepon  string `json:"telepon"`
	Alamat   string `json:"alamat"`
}

type Client struct {
	gorm.Model
	TenantID uint   `json:"tenant_id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	HP       string `json:"hp"`
	Alamat   string `json:"alamat"`
}

type Product struct {
	gorm.Model
	TenantID uint    `json:"tenant_id"`
	Nama     string  `json:"nama"`
	Harga    float64 `json:"harga"`
}

type Invoice struct {
	gorm.Model
	TenantID          uint          `json:"tenant_id"`
	NomorInvoice      string        `json:"nomor_invoice"`
	ClientID          uint          `json:"client_id"`
	Client            Client        `json:"client" gorm:"foreignKey:ClientID"`
	TanggalTerbit     string        `json:"tanggal_terbit"`
	TanggalJatuhTempo string        `json:"tanggal_jatuh_tempo"`
	Deskripsi         string        `json:"deskripsi"`
	TemaID            int           `json:"tema_id"`
	Status            string        `json:"status"` // "Lunas", "Belum Bayar", "Jatuh Tempo"
	TotalAmount       float64       `json:"total_amount"`
	UniqueToken       string        `gorm:"unique" json:"unique_token"`
	Items             []InvoiceItem `json:"items" gorm:"foreignKey:InvoiceID;constraint:OnDelete:CASCADE;"`
}

type InvoiceItem struct {
	gorm.Model
	InvoiceID uint    `json:"invoice_id"`
	ProductID *uint   `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	NamaItem  string  `json:"nama_item"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Subtotal  float64 `json:"subtotal"`
}

var db *gorm.DB
var jwtSecretKey = []byte("super_secret_jwt_key_kilat_invoice_2026")

// ==========================================
// 2. MAIN FUNCTION & DATABASE SETUP
// ==========================================

func main() {
	_ = godotenv.Load()

	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		jwtSecretKey = []byte(secret)
	}

	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "invoice_db")
	port := getEnv("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke PostgreSQL: %v", err)
	}

	log.Println("Berhasil terhubung ke PostgreSQL!")

	// Auto Migrate Tabel
	db.AutoMigrate(&Whitelist{}, &Tenant{}, &User{}, &Client{}, &Product{}, &Invoice{}, &InvoiceItem{})

	// Menjalankan Seeder Data Otomatis
	SeedData(db)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// ------------------------------------------
	// PUBLIC ENDPOINTS
	// ------------------------------------------
	r.POST("/api/auth/login", loginHandler)
	r.POST("/api/auth/register-owner", registerOwnerWithWhitelist)
	r.GET("/api/public/invoices/:token", getInvoiceByToken)

	// Endpoint Tim Kamu buat daftarin Email Whitelist
	r.POST("/api/superadmin/whitelist", addWhitelistEmail)

	// ------------------------------------------
	// PROTECTED ENDPOINTS (Dibutuhkan Login)
	// ------------------------------------------
	api := r.Group("/api")
	api.Use(AuthMiddleware())
	{
		api.GET("/user/profile", getUserProfile)
		api.GET("/dashboard/stats", getDashboardStats)

		// Khusus Role Owner: Bikin Akun Admin Baru untuk Perusahaannya
		api.POST("/owner/create-admin", createAdminByOwner)

		// Client Master
		api.GET("/clients", getClients)
		api.POST("/clients", createClient)
		api.PUT("/clients/:id", updateClient)
		api.DELETE("/clients/:id", deleteClient)

		// Product Master
		api.GET("/products", getProducts)
		api.POST("/products", createProduct)

		// Invoices
		api.GET("/invoices", getInvoices)
		api.POST("/invoices", createInvoice)
		api.PUT("/invoices/:id/status", updateInvoiceStatus)
	}

	serverPort := getEnv("PORT", "8080")
	log.Printf("Server berjalan di port %s", serverPort)
	r.Run(":" + serverPort)
}

// ==========================================
// 3. AUTO-SEEDER DATA AWAL
// ==========================================

func SeedData(db *gorm.DB) {
	var userCount int64
	db.Model(&User{}).Count(&userCount)

	if userCount == 0 {
		log.Println("Database masih kosong. Menjalankan Auto-Seeder...")

		tenant := Tenant{
			NamaPerusahaan: "PT Kilat",
			Alamat:         "Jl. Pemuda No. 12, Surabaya",
			Kontak:         "081234567890",
		}
		db.Create(&tenant)

		hashedPassword, _ := HashPassword("password123")

		// Akun Owner
		ownerUser := User{
			TenantID: tenant.ID,
			Nama:     "Rafli Pratama (Owner)",
			Email:    "owner@kilat.co.id",
			Username: "owner_kilat",
			Password: hashedPassword,
			Role:     "owner",
			Telepon:  "08123456789",
			Alamat:   "Jl. Pemuda No. 12, Surabaya",
		}
		db.Create(&ownerUser)

		// Akun Admin
		adminUser := User{
			TenantID: tenant.ID,
			Nama:     "Kaka Admin",
			Email:    "admin@kilat.co.id",
			Username: "admin_kilat",
			Password: hashedPassword,
			Role:     "admin",
			Telepon:  "08987654321",
			Alamat:   "Jl. Pemuda No. 12, Surabaya",
		}
		db.Create(&adminUser)

		// Whitelist Email Sample
		db.Create(&Whitelist{
			Email:          "klienbaru@gmail.com",
			NamaPerusahaan: "CV Mitra Jaya",
			IsUsed:         false,
		})

		// Sample Klien
		client1 := Client{TenantID: tenant.ID, Nama: "Studio Nusantara", Email: "halo@studionusantara.com", HP: "085611223344", Alamat: "Jl. Tunjungan No. 88, Surabaya"}
		client2 := Client{TenantID: tenant.ID, Nama: "Toko Berkah", Email: "tokoberkah@gmail.com", HP: "087855443322", Alamat: "Jl. Pahlawan No. 10, Sidoarjo"}
		db.Create(&client1)
		db.Create(&client2)

		// Sample Produk
		prod1 := Product{TenantID: tenant.ID, Nama: "Jasa Desain Brand & Logo", Harga: 5500000}
		prod2 := Product{TenantID: tenant.ID, Nama: "Pengembangan Website Company Profile", Harga: 12500000}
		db.Create(&prod1)
		db.Create(&prod2)

		log.Println("✅ Auto-Seeder Berhasil! Data awal telah dibuat.")
	}
}

// ==========================================
// 4. HELPER & MIDDLEWARE
// ==========================================

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateJWT(userID uint, tenantID uint, role string, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"tenant_id": tenantID,
		"role":      role,
		"username":  username,
		"exp":       time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Header otorisasi tidak ditemukan"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", uint(claims["user_id"].(float64)))
			c.Set("tenant_id", uint(claims["tenant_id"].(float64)))
			c.Set("role", claims["role"].(string))
			c.Set("username", claims["username"].(string))
		}
		c.Next()
	}
}

// ==========================================
// 5. HANDLERS
// ==========================================

type WhitelistInput struct {
	Email          string `json:"email" binding:"required"`
	NamaPerusahaan string `json:"nama_perusahaan" binding:"required"`
}

func addWhitelistEmail(c *gin.Context) {
	var input WhitelistInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	whitelist := Whitelist{
		Email:          input.Email,
		NamaPerusahaan: input.NamaPerusahaan,
		IsUsed:         false,
	}

	if err := db.Create(&whitelist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email ini sudah ada di whitelist!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Email berhasil di-whitelist!", "data": whitelist})
}

type RegisterOwnerInput struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Telepon  string `json:"telepon"`
}

func registerOwnerWithWhitelist(c *gin.Context) {
	var input RegisterOwnerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var wl Whitelist
	if err := db.Where("email = ? AND is_used = ?", input.Email, false).First(&wl).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Email kamu belum terdaftar di whitelist kami atau sudah digunakan."})
		return
	}

	tenant := Tenant{
		NamaPerusahaan: wl.NamaPerusahaan,
		Kontak:         input.Telepon,
	}
	db.Create(&tenant)

	hashedPassword, _ := HashPassword(input.Password)
	user := User{
		TenantID: tenant.ID,
		Nama:     input.Nama,
		Email:    input.Email,
		Username: input.Username,
		Password: hashedPassword,
		Role:     "owner",
		Telepon:  input.Telepon,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username atau email sudah dipakai."})
		return
	}

	wl.IsUsed = true
	db.Save(&wl)

	token, _ := GenerateJWT(user.ID, tenant.ID, user.Role, user.Username)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi Owner Berhasil!",
		"token":   token,
		"user":    user,
		"tenant":  tenant,
	})
}

type CreateAdminInput struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func createAdminByOwner(c *gin.Context) {
	role := c.GetString("role")
	tenantID := c.GetUint("tenant_id")

	if role != "owner" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Hanya Owner perusahaan yang bisa menambah Admin!"})
		return
	}

	var input CreateAdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := HashPassword(input.Password)
	user := User{
		TenantID: tenantID,
		Nama:     input.Nama,
		Email:    input.Email,
		Username: input.Username,
		Password: hashedPassword,
		Role:     "admin",
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username atau email sudah digunakan."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Akun Admin berhasil dibuat!", "user": user})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func loginHandler(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}

	if !CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}

	token, _ := GenerateJWT(user.ID, user.TenantID, user.Role, user.Username)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Berhasil!",
		"token":   token,
		"user": gin.H{
			"id":        user.ID,
			"tenant_id": user.TenantID,
			"nama":      user.Nama,
			"role":      user.Role,
			"username":  user.Username,
		},
	})
}

func getUserProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	var user User
	db.First(&user, userID)
	c.JSON(http.StatusOK, user)
}

func getClients(c *gin.Context) {
	tenantID := c.GetUint("tenant_id")
	var clients []Client
	db.Where("tenant_id = ?", tenantID).Order("created_at desc").Find(&clients)
	c.JSON(http.StatusOK, clients)
}

func createClient(c *gin.Context) {
	tenantID := c.GetUint("tenant_id")
	var client Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client.TenantID = tenantID
	db.Create(&client)
	c.JSON(http.StatusCreated, client)
}

func updateClient(c *gin.Context) {
	tenantID := c.GetUint("tenant_id")
	id := c.Param("id")
	var client Client
	if err := db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&client).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Klien tidak ditemukan"})
		return
	}
	c.ShouldBindJSON(&client)
	db.Save(&client)
	c.JSON(http.StatusOK, client)
}

func deleteClient(c *gin.Context) {
	tenantID := c.GetUint("tenant_id")
	id := c.Param("id")
	db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&Client{})
	c.JSON(http.StatusOK, gin.H{"message": "Klien dihapus"})
}

func getProducts(c *gin.Context) {
	tenantID := c.GetUint("tenant_id")
	var products []Product
	db.Where("tenant_id = ?", tenantID).Find(&products)
	c.JSON(http.StatusOK, products)
}

func createProduct(c *gin.Context) {
	tenantID := c.GetUint("tenant_id")
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.TenantID = tenantID
	db.Create(&product)
	c.JSON(http.StatusCreated, product)
}

type CreateInvoiceItemInput struct {
	ProductID *uint   `json:"product_id"`
	NamaItem  string  `json:"nama_item"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	UnitPrice float64 `json:"unit_price"`
}

type CreateInvoiceInput struct {
	NomorInvoice      string                   `json:"nomor_invoice"`
	ClientID          uint                     `json:"client_id" binding:"required"`
	TanggalTerbit     string                   `json:"tanggal_terbit" binding:"required"`
	TanggalJatuhTempo string                   `json:"tanggal_jatuh_tempo" binding:"required"`
	Deskripsi         string                   `json:"deskripsi"`
	TemaID            int                      `json:"tema_id"`
	Items             []CreateInvoiceItemInput `json:"items" binding:"required,min=1"`
}

func createInvoice(c *gin.Context) {
	tenantID := c.GetUint("tenant_id")
	var input CreateInvoiceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.NomorInvoice == "" {
		input.NomorInvoice = fmt.Sprintf("INV-%s-%03d", time.Now().Format("2006"), time.Now().Unix()%1000)
	}

	var grandTotal float64 = 0
	var invoiceItems []InvoiceItem

	for _, itemInput := range input.Items {
		unitPrice := itemInput.UnitPrice
		namaItem := itemInput.NamaItem

		if itemInput.ProductID != nil && *itemInput.ProductID > 0 {
			var prod Product
			if err := db.Where("id = ? AND tenant_id = ?", *itemInput.ProductID, tenantID).First(&prod).Error; err == nil {
				if unitPrice == 0 {
					unitPrice = prod.Harga
				}
				if namaItem == "" {
					namaItem = prod.Nama
				}
			}
		}

		subtotal := float64(itemInput.Quantity) * unitPrice
		grandTotal += subtotal

		invoiceItems = append(invoiceItems, InvoiceItem{
			ProductID: itemInput.ProductID,
			NamaItem:  namaItem,
			Quantity:  itemInput.Quantity,
			UnitPrice: unitPrice,
			Subtotal:  subtotal,
		})
	}

	invoice := Invoice{
		TenantID:          tenantID,
		NomorInvoice:      input.NomorInvoice,
		ClientID:          input.ClientID,
		TanggalTerbit:     input.TanggalTerbit,
		TanggalJatuhTempo: input.TanggalJatuhTempo,
		Deskripsi:         input.Deskripsi,
		TemaID:            input.TemaID,
		Status:            "Belum Bayar",
		TotalAmount:       grandTotal,
		UniqueToken:       uuid.New().String(),
		Items:             invoiceItems,
	}

	db.Create(&invoice)
	db.Preload("Client").Preload("Items").First(&invoice, invoice.ID)
	c.JSON(http.StatusCreated, invoice)
}

func getInvoices(c *gin.Context) {
	tenantID := c.GetUint("tenant_id")
	var invoices []Invoice
	db.Where("tenant_id = ?", tenantID).Preload("Client").Preload("Items").Order("created_at desc").Find(&invoices)
	c.JSON(http.StatusOK, invoices)
}

func getInvoiceByToken(c *gin.Context) {
	token := c.Param("token")
	var invoice Invoice
	if err := db.Preload("Client").Preload("Items").Where("unique_token = ?", token).First(&invoice).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func updateInvoiceStatus(c *gin.Context) {
	tenantID := c.GetUint("tenant_id")
	id := c.Param("id")
	var body struct {
		Status string `json:"status" binding:"required"`
	}
	c.ShouldBindJSON(&body)

	var invoice Invoice
	if err := db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&invoice).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice tidak ditemukan"})
		return
	}

	invoice.Status = body.Status
	db.Save(&invoice)
	c.JSON(http.StatusOK, invoice)
}

func getDashboardStats(c *gin.Context) {
	tenantID := c.GetUint("tenant_id")

	var totalPendapatan float64
	var totalMenunggu float64
	var countMenunggu int64
	var totalJatuhTempo float64
	var countJatuhTempo int64

	db.Model(&Invoice{}).Where("tenant_id = ? AND status = ?", tenantID, "Lunas").Select("COALESCE(SUM(total_amount), 0)").Scan(&totalPendapatan)
	db.Model(&Invoice{}).Where("tenant_id = ? AND status = ?", tenantID, "Belum Bayar").Select("COALESCE(SUM(total_amount), 0)").Scan(&totalMenunggu)
	db.Model(&Invoice{}).Where("tenant_id = ? AND status = ?", tenantID, "Belum Bayar").Count(&countMenunggu)
	db.Model(&Invoice{}).Where("tenant_id = ? AND status = ?", tenantID, "Jatuh Tempo").Select("COALESCE(SUM(total_amount), 0)").Scan(&totalJatuhTempo)
	db.Model(&Invoice{}).Where("tenant_id = ? AND status = ?", tenantID, "Jatuh Tempo").Count(&countJatuhTempo)

	c.JSON(http.StatusOK, gin.H{
		"total_pendapatan": totalPendapatan,
		"menunggu_pembayaran": gin.H{"total": totalMenunggu, "count": countMenunggu},
		"jatuh_tempo":         gin.H{"total": totalJatuhTempo, "count": countJatuhTempo},
	})
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}