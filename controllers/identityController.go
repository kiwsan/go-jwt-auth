package controllers

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kiwsan/go-jwt-auth/data"
	"github.com/kiwsan/go-jwt-auth/entities"
	"github.com/kiwsan/go-jwt-auth/models"
	"github.com/kiwsan/go-jwt-auth/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func MeGetHandler(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	return c.String(http.StatusOK, "Welcome "+username+"!")
}

func LoginPostHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	client, err := data.ClientDb()
	if err != nil {
		return err
	}

	user := new(entities.User)
	collection := client.Database("identityService").Collection("user")

	// Check in your db if the user exists or not
	err = collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid username.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid password.")
	}

	// Get claims from database

	// Create token
	jwtToken, err := utils.CreateToken(username) // Add claims
	if err != nil {
		return err
	}

	// return jwtToken
	return c.JSON(http.StatusOK, jwtToken)
}

func RegisterPostHandler(c echo.Context) error {

	req := new(models.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}

	client, err := data.ClientDb()
	if err != nil {
		return err
	}

	collection := client.Database("identityService").Collection("user")

	user, err := entities.NewUser(req.Username, req.Password)
	if err != nil {
		return err
	}

	if !entities.IsValidPassword(req.Password, req.ConfirmPassword) {
		return c.JSON(http.StatusBadRequest, "Invalid confirm password.")
	}

	err = collection.FindOne(context.TODO(), bson.D{{"username", req.Username}}).Decode(&user)
	if err == nil {
		return c.String(http.StatusBadRequest, "Username already Exists!")
	}

	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	fmt.Println("Inserted a Single Document: ", result.InsertedID)

	return c.JSON(http.StatusCreated, req)
}
