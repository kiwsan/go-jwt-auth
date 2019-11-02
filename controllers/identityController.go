package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kiwsan/go-jwt-auth/data"
	"github.com/kiwsan/go-jwt-auth/entities"
	"github.com/kiwsan/go-jwt-auth/messages/events"
	"github.com/kiwsan/go-jwt-auth/models"
	"github.com/kiwsan/go-jwt-auth/utils"
	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func errorHandler(e error) error {
	fmt.Println(e)

	return e
}

func MeGetHandler(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	return c.String(http.StatusOK, "Welcome "+email+"!")
}

func LoginPostHandler(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	client, err := data.DbConnection()
	if err != nil {
		return errorHandler(err)
	}

	user := new(entities.User)
	collection := client.Database("identityService").Collection("users")

	// Check in your db if the user exists or not
	err = collection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)

	if err != nil {
		return c.String(http.StatusBadRequest, "Email was not found.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid password.")
	}

	// Create token
	jwtToken, err := utils.CreateToken(email) // Add claims
	if err != nil {
		return errorHandler(err)
	}

	// Add token to database
	refreshToken, err := entities.NewRefreshToken(user.Id.Hex(), email, jwtToken.AccessToken)
	if err != nil {
		return errorHandler(err)
	}

	refreshTokens := client.Database("identityService").Collection("refreshTokens")
	result, err := refreshTokens.InsertOne(context.TODO(), refreshToken)
	if err != nil {
		return errorHandler(err)
	}

	fmt.Println("Inserted a Single Document: ", result.InsertedID)

	// return jwtToken
	return c.JSON(http.StatusOK, jwtToken)
}

func RegisterPostHandler(c echo.Context) error {

	req := new(models.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}

	if req.Email == "" {
		return c.JSON(http.StatusBadRequest, "Invalid email")
	}

	if req.Password == "" {
		return c.JSON(http.StatusBadRequest, "Invalid password")
	}

	client, err := data.DbConnection()
	if err != nil {
		return errorHandler(err)
	}

	collection := client.Database("identityService").Collection("users")

	user, err := entities.NewUser(req.Email, req.Password)
	if err != nil {
		return errorHandler(err)
	}

	if !entities.IsValidPassword(req.Password, req.ConfirmPassword) {
		return c.JSON(http.StatusBadRequest, "Invalid confirm password.")
	}

	err = collection.FindOne(context.TODO(), bson.D{{"email", req.Email}}).Decode(&user)
	if err == nil {
		return c.String(http.StatusBadRequest, "Email already Exists!")
	}

	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return errorHandler(err)
	}

	//https://stackoverflow.com/questions/49933249/mongo-go-driver-get-objectid-from-insert-result
	if uid, ok := result.InsertedID.(primitive.ObjectID); ok {
		//add account
		conn, ch, q, err := data.RabbitMqConnection()
		if err != nil {
			return errorHandler(err)
		}

		registered := events.Registered{UserId: uid.Hex(), Email: user.Email, Role: user.Role}
		body, err := json.Marshal(registered)
		if err != nil {
			return errorHandler(err)
		}

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})

		if err != nil {
			return errorHandler(err)
		}

		defer conn.Close()
		defer ch.Close()
	} else {
		return c.JSON(http.StatusBadRequest, "Cannot convert uid.")
	}

	return c.JSON(http.StatusCreated, "The user has been created.")
}
