package handler

import (
	"fmt"
	"rea_games/entity"
	"rea_games/helper"
	"rea_games/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
    userRepo *repository.UserRepository
    inputter  *helper.Inputter
}

func NewAuthHandler() *AuthHandler {
    return &AuthHandler{
        userRepo: repository.NewUserRepository(),
        inputter:  helper.NewInputter(),
    }
}

func (h *AuthHandler) Register() (int, error) {
    var user entity.User
    
    user.Email = h.inputter.ReadInput("Enter email: ")
    password := h.inputter.ReadInput("Enter password: ")

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return 0, err
    }
    user.PasswordHash = string(hashedPassword)

    err = h.userRepo.CreateUser(&user)
    if err != nil {
        return 0, err
    }

    fmt.Println("User registered successfully!")
    return user.ID, nil
}

func (h *AuthHandler) Login() (int, error) {
    var loginReq entity.LoginRequest
    
    loginReq.Email = h.inputter.ReadInput("Enter email: ")
    loginReq.Password = h.inputter.ReadInput("Enter password: ")

    user, err := h.userRepo.GetUserByEmail(loginReq.Email)
    if err != nil {
        return 0, fmt.Errorf("invalid credentials")
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password))
    if err != nil {
        return 0, fmt.Errorf("invalid credentials")
    }

    fmt.Println("Login successful!")
    return user.ID, nil
}