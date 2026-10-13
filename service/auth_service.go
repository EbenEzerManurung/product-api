package service

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"

    "product-api/config"
    "product-api/model"
    "product-api/repository"
)

type AuthService struct {
    userRepo *repository.UserRepository
    cfg      *config.Config
}

func NewAuthService(repo *repository.UserRepository, cfg *config.Config) *AuthService {
    return &AuthService{userRepo: repo, cfg: cfg}
}

type RegisterInput struct {
    Username             string `json:"username" binding:"required,min=3"`
    Password             string `json:"password" binding:"required,min=6"`
    PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password"`
}

type LoginInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
    AccessToken  string `json:"authentication_token"`
    RefreshToken string `json:"refresh_token"`
}

func (s *AuthService) Register(input RegisterInput) (*model.User, error) {
    existing, _ := s.userRepo.FindByUsername(input.Username)
    if existing != nil {
        return nil, errors.New("username sudah terdaftar")
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &model.User{
        Username: input.Username,
        Password: string(hashed),
    }

    if err := s.userRepo.Create(user); err != nil {
        return nil, err
    }
    return user, nil
}

func (s *AuthService) Login(input LoginInput) (*TokenResponse, error) {
    user, err := s.userRepo.FindByUsername(input.Username)
    if err != nil {
        return nil, errors.New("kredensial tidak valid")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        return nil, errors.New("kredensial tidak valid")
    }

    accessToken, err := s.generateToken(user.ID, user.Username, s.cfg.AccessExpiry, "access")
    if err != nil {
        return nil, err
    }

    refreshToken, err := s.generateToken(user.ID, user.Username, s.cfg.RefreshExpiry, "refresh")
    if err != nil {
        return nil, err
    }

    return &TokenResponse{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
    }, nil
}

func (s *AuthService) generateToken(userID uint, username string, expiry time.Duration, tokenType string) (string, error) {
    claims := jwt.MapClaims{
        "sub":      userID,
        "username": username,
        "type":     tokenType,
        "exp":      time.Now().Add(expiry).Unix(),
        "iat":      time.Now().Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.cfg.JWTSecret))
}

func (s *AuthService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("metode signing tidak sesuai")
        }
        return []byte(s.cfg.JWTSecret), nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, errors.New("token tidak valid")
}