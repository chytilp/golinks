package main

import (
	"fmt"
	"log"
	"flag"

	//"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	//"github.com/chytilp/golinks/controller"
	"github.com/chytilp/golinks/database"
	//"github.com/chytilp/golinks/middleware"
	"github.com/chytilp/golinks/model"
	"github.com/chytilp/golinks/data"
)

func main() {
    var automigrateOn bool
    var fillDataOn bool

    flag.BoolVar(&automigrateOn, "automigrate", false, "Run automigrate")
	flag.BoolVar(&fillDataOn, "fill", false, "Fill db with test data")

    flag.Parse()

	loadEnv()
	database.Connect()
	if automigrateOn {
		fmt.Println("Run automigrating database")
		createDatabase()
	}

	if fillDataOn {
		fmt.Println("Run filling data to database")
		fillData()
	}
	
	//serveApplication()
	fmt.Println("Done")
}

func createDatabase() {
	database.Database.SetupJoinTable(&model.User{}, "Links", &model.UserLink{})
	database.Database.AutoMigrate(&model.Category{})
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Link{})
	database.Database.AutoMigrate(&model.Role{})
	database.Database.AutoMigrate(&model.UserLink{})
}

func fillData() {
	err := data.FillData()
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Filling database with test data successfully finished.")
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	/*router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")*/
}
