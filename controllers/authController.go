package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"go-forum-thingy/database"
	"go-forum-thingy/models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const SecretKey = "SECRET-STRING"

func Register(c *fiber.Ctx) error {
	//var data map[string]string

	var data struct {
		name string
		email string
		password []byte
	}

	data.name = c.FormValue("name")
	data.email = c.FormValue("email")
	data.password = []byte(c.FormValue("password"))

	//TODO also check for empty password

	if (data.name == "" || data.email == "") {
		log.Info().Msg("Non valid form values received")
		c.Status(fiber.StatusNotAcceptable)

		return c.Render("register", fiber.Map{
		})
	}

	password, _ := bcrypt.GenerateFromPassword(data.password, 14)

user := models.User{
	Name: data.name,
	Email: data.email,
	Password: password,
}

	database.DB.Create(&user)

	//return c.JSON(user)
	return c.Redirect("/")
}


func Login(c *fiber.Ctx) error {
	// Make option to decide between username or email
	var data struct {
		email string
		password []byte
	}

	data.email = c.FormValue("email")
	data.password = []byte(c.FormValue("password"))

	var user models.User


	database.DB.Where("email = ?", data.email).First(&user)
	//If the query results in Id = 0 then an error is returned
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		log.Info().Msg("User not found")
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//Comparing passwords using the stored data in the models.User struct
	if err := bcrypt.CompareHashAndPassword(user.Password, data.password); err != nil {
		c.Status(fiber.StatusBadRequest)
		log.Info().Msg("Failed to authenticate")
		return c.JSON(fiber.Map{
			"message": "Failed to authenticate",
		})
	}

	//Setting jwt token for the current user ttl of 24 hours
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour *24).Unix(),
	})

	//Create the token
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		log.Info().Msg("Failed to create a jwt token for the current user")
	}

	//Setting cookie struct
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour *24),
		HTTPOnly: true,
	}

	//Storing cookie with the jwt in the response struct cookie field
	c.Cookie(&cookie)

	c.Status(fiber.StatusOK)
	return c.Redirect("/")

}

func Logout(c *fiber.Ctx) error {
	//Deleting the cookie works by setting the expirary date in the past
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	//Setting expired cookie and returning a message
	c.Cookie(&cookie)
	return c.Redirect("/")
}

func User(c *fiber.Ctx) error {
	//Fetching cookie from response
	cookie := c.Cookies("jwt")
	//Check if cookie/jwt token is valid
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized! invalid token!",
		})
	}
	//TODO: Enter description for this part
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	return c.JSON(user)
}