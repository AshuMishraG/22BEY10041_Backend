package api

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "net/http"
    "yourproject/pkg/auth"
)

func RegisterHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var user auth.User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        err := auth.Register(db, user.Email, user.Password)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "User registered"})
    }
}

func LoginHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var user auth.User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        token, err := auth.Login(db, user.Email, user.Password)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"token": token})
    }
}