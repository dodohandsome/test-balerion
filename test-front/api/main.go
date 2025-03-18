package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Card struct {
	ID          string    `json:"id"`
	Tag         string    `json:"tag"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ListCard struct {
	cards []Card
	mu    sync.Mutex
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var cards = &ListCard{}

func (c *ListCard) AddCard(card Card) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cards = append(c.cards, card)
}

func (c *ListCard) GetCards(role string) []Card {
	var cards []Card
	if role == "ADMIN" {
		cards = c.cards
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].CreatedAt.After(cards[j].CreatedAt)
		})
	} else if role == "USER" {
		for _, card := range c.cards {
			if card.Tag == "USER" {
				cards = append(cards, card)
			}
		}
	}
	return cards
}

func (c *ListCard) FindCardsByRole(role string) []Card {
	c.mu.Lock()
	defer c.mu.Unlock()
	var cards []Card
	for _, card := range c.cards {
		if card.Tag == role {
			cards = append(cards, card)
		}
	}
	return cards
}

func authMiddleware(c *gin.Context) {
	tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	if err != nil && token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}

	c.Set("role", claims["role"])
	c.Next()
}

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.POST("/login", loginHandler)
	authorized := router.Group("/cards", authMiddleware)
	{
		authorized.GET("", getCards)
		authorized.POST("", createCard)
	}

	log.Fatal(router.Run(":3000"))
}

func getCards(c *gin.Context) {
	role := c.GetString("role")
	c.JSON(http.StatusOK, cards.GetCards(role))
}

func createCard(c *gin.Context) {
	role := c.GetString("role")

	var newCard Card
	if err := c.BindJSON(&newCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	prefix := "MEMO"
	if role == "ADMIN" {
		prefix = "ADMIN"
	}
	newCard.ID = fmt.Sprintf("%s-%d", prefix, len(cards.FindCardsByRole(role))+1)
	newCard.Tag = role
	newCard.CreatedAt = time.Now()
	cards.AddCard(newCard)
	c.JSON(http.StatusCreated, newCard)
}

func loginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	var tokenString string
	if req.Username == "admin1" {
		tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWRtaW4xIiwiZW1haWwiOiJhZG1pbi5lbWFpbEBnbWFpbC5jb20iLCJyb2xlIjoiQURNSU4ifQ.91VaQcMDdRWOj849ddLZO7pR_qjl_DpHdaaYCYfakkg"
	} else if req.Username == "user1" {
		tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlcjEiLCJlbWFpbCI6InVzZXIuZW1haWxAZ21haWwuY29tIiwicm9sZSI6IlVTRVIifQ.IgQln56kjBGc66IAjRMjeJtscM2u--Uz5Ul01r1f874"
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if err != nil && token.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot parse token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot parse claims"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"profile": claims,
	})
}
