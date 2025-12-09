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

// @title BucketWise API
// @version 1.0
// @description The BucketWise API provides endpoints to manage financial categories and transactions.
// @contact.name Denis Amieva
// @contact.url https://github.com/damieva/bucketWise
// @license.name MIT
// @host localhost:8001
// @BasePath /
func main() {
	// Load environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Gin
	ginEngine := gin.Default()

	// Connect MongoDB
	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err.Error())
	}

	// ---------- Repositories ----------
	categoryRepo := mongo.CategoryRepo{Client: client}
	transactionRepo := mongo.TransactionRepo{Client: client}

	// ---------- Services ----------
	categorySrv := services.CategoryService{Repo: categoryRepo}
	transactionSrv := services.TransactionService{Repo: transactionRepo}

	// ---------- UseCases ----------
	categoryUC := usecases.CategoryUseCase{
		CategoryService:    categorySrv,
		TransactionService: transactionSrv,
	}
	transactionUC := usecases.TransactionUseCase{
		TransactionService: transactionSrv,
	}

	// ---------- Handlers ----------
	categoryHandler := handlers.CategoryHandler{CategoryUC: categoryUC}
	transactionHandler := handlers.TransactionHandler{TransactionUC: transactionUC}

	// ---------- Routes ----------
	CategoriesRouteGroup := ginEngine.Group("/categories")
	{
		CategoriesRouteGroup.POST("", categoryHandler.CreateCategory)
		CategoriesRouteGroup.GET("", categoryHandler.ListCategories)
		CategoriesRouteGroup.PUT("/:name", categoryHandler.UpdateCategory)
		CategoriesRouteGroup.DELETE("/", categoryHandler.DeleteCategories)
	}

	TransactionsRouteGroup := ginEngine.Group("/transactions")
	{
		TransactionsRouteGroup.POST("", transactionHandler.CreateTransaction)
		TransactionsRouteGroup.GET("", transactionHandler.ListTransactions)
		TransactionsRouteGroup.DELETE("", transactionHandler.DeleteTransactions)
	}

	// ---------- Swagger ----------
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ---------- Run ----------
	log.Fatalln(ginEngine.Run(":8001"))
}
