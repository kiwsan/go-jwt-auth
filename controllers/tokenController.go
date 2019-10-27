package controllers

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kiwsan/go-jwt-auth/data"
	"github.com/kiwsan/go-jwt-auth/entities"
	"github.com/kiwsan/go-jwt-auth/models"
	"github.com/kiwsan/go-jwt-auth/services"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func RefreshAccessTokenPostHandler(c echo.Context) error {

	refreshToken := c.Param("refreshToken")
	req := new(models.RefreshTokenRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	client, err := data.ClientDb()
	if err != nil {
		return err
	}

	rt := new(entities.RefreshToken)
	collection := client.Database("identityService").Collection("refreshTokens")

	// validate token
	err = collection.FindOne(context.TODO(), bson.D{{"token", req.Token}}).Decode(&rt)
	if err != nil {
		return c.String(http.StatusBadRequest, "Refresh token was not found.")
	}

	// check token had invoked
	if rt.RevokedAt.After(time.Time{}) {
		return c.String(http.StatusBadRequest, "Refresh token: "+req.Token+" was revoked.")
	}

	// validate user
	user := new(entities.User)
	err = collection.FindOne(context.TODO(), bson.D{{"email", rt.Email}}).Decode(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "Email was not found.")
	}

	// create a new token
	jwtToken, err := services.RefreshAccessToken(req, refreshToken)
	if err != nil {
		return err
	}

	// return jwtToken
	return c.JSON(http.StatusOK, jwtToken)
}

func RevokeAccessTokenHandler(c echo.Context) error {

	return c.String(http.StatusAccepted, "Accepted")
}

func RevokeRefreshTokenPostHandler(c echo.Context) error {

	refreshToken := c.Param("refreshToken")
	req := new(models.RevokeRefreshTokenRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	fmt.Println(refreshToken)

	client, err := data.ClientDb()
	if err != nil {
		return err
	}

	rt := new(entities.RefreshToken)
	collection := client.Database("identityService").Collection("refreshTokens")

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	// validate token
	err = collection.FindOne(context.TODO(), bson.D{{"token", req.Token}}).Decode(&rt)
	if err != nil || rt.Email != email {
		return c.String(http.StatusBadRequest, "Refresh token was not found.")
	}

	// update revoke date
	revokedAt := time.Now().UTC()

	filter := bson.M{"token": rt.Token}
	update := bson.D{{"$set",
		bson.D{
			{"revokedat", revokedAt},
		},
	}}

	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)

	return c.String(http.StatusAccepted, "Refresh token has been revoked.")
}
