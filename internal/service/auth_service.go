package service

import (
	"errors"
	"github.com/DryundeL/blog-graphql/internal/config"
	"github.com/DryundeL/blog-graphql/internal/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{db: db, cfg: cfg}
}

func (s *AuthService) Register(username, password, email string) (*models.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{Username: username, Password: string(hash), Email: email}
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(username, password string) (string, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"userID": user.ID})
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

func (s *AuthService) GetUserIDFromToken(tokenStr string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	return uint(claims["userID"].(float64)), nil
}
