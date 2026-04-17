package main

import (
	"bucketWise/pkg/adapters/input/http/api"
	"bucketWise/pkg/adapters/input/http/web"
	"bucketWise/pkg/adapters/output/persistence/mongo"
	"bucketWise/pkg/services"
	"bucketWise/pkg/usecases"
	"log"
	"os"

	_ "bucketWise/cmd/app/docs"

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
// @host localhost:8080
// @BasePath /
func main() {
	// Load environment
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize Gin
	ginEngine := gin.Default()
	ginEngine.HTMLRender = web.NewRenderer()
	ginEngine.Static("/static", "./web/static")

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
	categoryHandler := api.CategoryHandler{CategoryUC: categoryUC}
	transactionHandler := api.TransactionHandler{TransactionUC: transactionUC}
	categoryWebHandler := web.CategoryWebHandler{CategoryUC: categoryUC}

	// ---------- API Routes (/api) ----------
	apiGroup := ginEngine.Group("/api")
	{
		categoriesAPI := apiGroup.Group("/categories")
		categoriesAPI.POST("", categoryHandler.CreateCategory)
		categoriesAPI.GET("", categoryHandler.ListCategories)
		categoriesAPI.PUT("/:name", categoryHandler.UpdateCategory)
		categoriesAPI.DELETE("", categoryHandler.DeleteCategories)

		transactionsAPI := apiGroup.Group("/transactions")
		transactionsAPI.POST("", transactionHandler.CreateTransaction)
		transactionsAPI.GET("", transactionHandler.ListTransactions)
		transactionsAPI.DELETE("", transactionHandler.DeleteTransactions)
	}

	// ---------- Web Routes ----------
	categoriesWeb := ginEngine.Group("/categories")
	{
		categoriesWeb.GET("", categoryWebHandler.Index)
		categoriesWeb.POST("", categoryWebHandler.Create)
		categoriesWeb.DELETE("/:id", categoryWebHandler.Delete)
	}

	// ---------- Health ----------
	ginEngine.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})

	// ---------- Swagger ----------
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ---------- Run ----------
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatalln(ginEngine.Run(":" + port))
}
