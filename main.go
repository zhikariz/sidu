package main

import (
	"log"
	"net/http"
	"sidu/auth"
	"sidu/document"
	"sidu/handler"
	"sidu/helper"
	"sidu/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	db := helper.SetupDB()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Repository
	userRepository := user.NewRepository(db)
	documentRepository := document.NewRepository(db)

	// Service
	authService := auth.NewService()
	userService := user.NewService(userRepository)
	documentService := document.NewService(documentRepository)

	// Handler
	userHandler := handler.NewUserHandler(userService, authService)
	documentHandler := handler.NewDocumentHandler(documentService)

	// Api Versioning
	api := router.Group("/api/v1")

	// Endpoint
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	api.GET("/documents", documentHandler.GetDocuments)
	api.GET("/documents/:id", documentHandler.GetDocument)
	api.POST("/documents", authMiddleware(authService, userService), documentHandler.CreateDocument)
	api.PUT("/documents/:id", authMiddleware(authService, userService), documentHandler.UpdateDocument)
	api.DELETE("/documents/:id", authMiddleware(authService, userService), documentHandler.DeleteDocument)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserById(userId)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
