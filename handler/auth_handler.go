package handler

import (
	"bufio"
	"fmt"
	"os"
	"rea_games/entity"
	"rea_games/repository"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
    userRepo *repository.UserRepository
    scanner  *bufio.Scanner
}

func NewAuthHandler() *AuthHandler {
    return &AuthHandler{
        userRepo: repository.NewUserRepository(),
        scanner:  bufio.NewScanner(os.Stdin),
    }
}

func (h *AuthHandler) readInput(prompt string) string {
    fmt.Print(prompt)
    h.scanner.Scan()
    return strings.TrimSpace(h.scanner.Text())
}

func (h *AuthHandler) Register() (int, error) {
    var user entity.User
    
    user.Email = h.readInput("Enter email: ")
    password := h.readInput("Enter password: ")

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
    
    loginReq.Email = h.readInput("Enter email: ")
    loginReq.Password = h.readInput("Enter password: ")

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