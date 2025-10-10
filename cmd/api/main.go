// @title BucketWise API
// @version 1.0
// @description ## Overview
// @description The BucketWise API provides endpoints to manage financial categories and transactions.
// @description It allows users to create, list, update, and delete categories, helping organize spending into personalized budgets.
// @description
// @description ## Features
// @description - Create, read, update, and delete categories
// @description - Categorize and manage transactions
// @description - RESTful interface using JSON
// @description
// @description ## Contact
// @description **Author:** Denis Amieva
// @description **Repository:** [github.com/damieva/bucketWise](https://github.com/damieva/bucketWise)
// @termsOfService http://swagger.io/terms/
// @contact.name Denis Amieva
// @contact.url https://github.com/damieva/bucketWise
// @license.name MIT
// @host localhost:8001
// @BasePath /

package main

import (
	"bucketWise/cmd/api/handlers"
	"bucketWise/pkg/adapters/persistence/mongo"
	"bucketWise/pkg/services"
	"bucketWise/pkg/usecases"
	"log"
	"os"

	_ "bucketWise/cmd/api/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()
	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err.Error())
	}

	categoryRepo := mongo.CategoryRepo{
		Client: client,
	}

	categorySrv := services.CategoryService{
		Repo: categoryRepo,
	}

	categoryUC := usecases.CategoryUseCase{
		CategoryService: categorySrv,
	}

	categoryHandler := handlers.CategoryHandler{
		CategoryUC: categoryUC,
	}

	CategoriesRouteGroup := ginEngine.Group("/categories")
	{
		CategoriesRouteGroup.POST("", categoryHandler.CreateCategory)
		CategoriesRouteGroup.GET("", categoryHandler.ListAllCategories)
		CategoriesRouteGroup.GET("/:name", categoryHandler.GetCategoryByName)
		CategoriesRouteGroup.PUT("/:name", categoryHandler.UpdateCategory)
		CategoriesRouteGroup.DELETE("/:name", categoryHandler.DeleteCategory)
	}

	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatalln(ginEngine.Run(":8001"))
}
