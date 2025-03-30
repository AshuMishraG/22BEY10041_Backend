package auth

import (
    "database/sql"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "time"
)

type User struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func Register(db *sql.DB, email, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    _, err = db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, hashedPassword)
    return err
}

func Login(db *sql.DB, email, password string) (string, error) {
    var userID int
    var storedPassword string
    err := db.QueryRow("SELECT id, password FROM users WHERE email = $1", email).Scan(&userID, &storedPassword)
    if err != nil {
        return "", err
    }
    err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
    if err != nil {
        return "", err
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    tokenString, err := token.SignedString([]byte("your-secret-key"))
    return tokenString, err
}