package main

import (
	"bucketWise/cmd/api/handlers"
	"bucketWise/pkg/adapters/persistence/mongo"
	"bucketWise/pkg/services"
	"bucketWise/pkg/usecases"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
		//CategoriesRouteGroup.GET("/:id", categoryHandler.GetByIDCategory)
		//CategoriesRouteGroup.PUT("/:id", categoryHandler.UpdateCategory)
		//CategoriesRouteGroup.DELETE("/:id", categoryHandler.DeleteCategory)
	}

	log.Fatalln(ginEngine.Run(":8001"))
}
