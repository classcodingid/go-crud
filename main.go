package main

import (
	"fmt"
	"go-crud/controllers"
	"go-crud/database"
	"go-crud/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin" // router HTTP
	"gorm.io/gorm"             // orm
)

var DB *gorm.DB

func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	//with gorilla mux
	// router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	//with gorilla mux
	// RegisterProductRoutes(router)

	//// Initialize Router
	//with go gin
	router := initRouter()
	router.Run(":8080")

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

// framework go gin
// api/user/register
// api/token
// api/secured/ping
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		//secure middleware
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}

// with gorilla mux
// func RegisterProductRoutes(router *mux.Router) {
// 	router.HandleFunc("/api/v1/products", controllers.GetProducts).Methods("GET")
// 	router.HandleFunc("/api/v1/products/{id}", controllers.GetProductById).Methods("GET")
// 	router.HandleFunc("/api/v1/products", controllers.CreateProduct).Methods("POST")
// 	router.HandleFunc("/api/v1/products/{id}", controllers.UpdateProduct).Methods("PUT")
// 	router.HandleFunc("/api/v1/products/{id}", controllers.DeleteProduct).Methods("DELETE")
// }
