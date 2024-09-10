package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/Niladri2003/server-monitor/server/app/models"
	"github.com/Niladri2003/server-monitor/server/pkg/utils"
	"github.com/Niladri2003/server-monitor/server/platform/cache"
	"github.com/Niladri2003/server-monitor/server/platform/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Create a new validator instance
var validate = validator.New()

func UserSignUp(c *fiber.Ctx) error {
	//Create a new user auth struct
	signUp := &models.User{}

	//Checking received data from JSON body.
	if err := c.BodyParser(signUp); err != nil {
		//Return status 400 and error message
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}
	// Validate sign-up fields
	if err := validate.Struct(signUp); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := utils.ValidatorErrors(validationErrors) // Format errors if needed
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   errorMessages,
		})
	}
	fmt.Printf("Parsed sign-up data: %+v\n", signUp)
	fmt.Println("Email=>", signUp.Email)

	////Create a new validator for User model.
	//validate := utils.NewValidator()
	//
	////Validate sign up fields.
	//if err := validate.Struct(signUp); err != nil {
	//	// Log validation errors for debugging
	//	fmt.Printf("Validation errors: %+v\n", utils.ValidatorErrors(err))
	//	//Return if some fields are not valid
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"error": true,
	//		"msg":   utils.ValidatorErrors(err),
	//	})
	//}
	//Create database connection
	collection := database.GetDbCollection("users")
	existingUser := &models.User{}
	err := collection.FindOne(context.Background(), bson.M{"email": signUp.Email}).Decode(existingUser)
	if err == nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "user already exists",
		})
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		// Other error occurred
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "cannot query database",
		})
	}

	//Create a new user struct
	user := models.User{
		ID:        primitive.NewObjectID(),
		FirstName: signUp.FirstName,
		LastName:  signUp.LastName,
		Email:     signUp.Email,
		Password:  utils.GeneratePassword(signUp.Password),
		UserRole:  signUp.UserRole,
		Verified:  true,
	}

	//Insert user in database
	collection = database.GetDbCollection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "cannont insert user inside db",
		})

	}

	//recipientName := signUp.FirstName + " " + signUp.LastName
	//recipientEmail := signUp.Email
	//subject := "Thanks for registering with us!"
	//mailTemplatePath := "templates/welcome_email.html"
	//
	//emailData := struct {
	//	Name      string
	//	LoginLink string
	//}{
	//	Name:      recipientName,
	//	LoginLink: "https://painting-ecommerce.vercel.app/signin",
	//}

	// Send confirmation email asynchronously
	//go func() {
	//	err := utils.SendEmailUsingMailgun(recipientEmail, subject, mailTemplatePath, emailData)
	//	if err != nil {
	//		fmt.Printf("Failed to send confirmation email: %v \n", err)
	//	} else {
	//		fmt.Println("Confirmation email sent successfully!")
	//	}
	//}()

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "User registration successful",
		"user":  user,
	})
}

func UserSignIn(c *fiber.Ctx) error {
	//Create a new user auth struct
	signIn := &models.SignIn{}

	if err := c.BodyParser(signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Validate sign-in fields
	if err := validate.Struct(signIn); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := utils.ValidatorErrors(validationErrors) // Format errors if needed
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   errorMessages,
		})
	}

	//create a database connection
	var user models.User
	collection := database.GetDbCollection("users")
	err := collection.FindOne(context.Background(), bson.M{"email": signIn.Email}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid email address",
		})
	}

	// Compare given user password with stored in found user.
	//compareUserPassword := utils.ComparePasswords(foundedUser.PasswordHash, signIn.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signIn.Password))
	if err != nil {
		//Return, if password is not incorrect
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Wrong Password",
		})
	}
	// Get role credentials from founded user.
	//userRole := foundedUser.UserRole.String()
	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(user.ID, user.UserRole)
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	//Define user ID.
	userId := user.ID.String()

	// Create a new Redis connection
	connRedis, err := cache.RedisConnection()
	if err != nil {
		//fmt.Println("Redis Error", err)
		//Return status 500 and Redis connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//Save refresh token to Redis.
	errSaveToRedis := connRedis.Set(context.Background(), userId, tokens.Refresh, 0).Err()
	fmt.Println(errSaveToRedis)
	if errSaveToRedis != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Token storing error",
		})
	}
	user.Password = ""

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          "login successful",
		"user_details": user,
		"tokens": fiber.Map{
			"access":       tokens.Access,
			"refreshToken": tokens.Refresh,
		},
	})
}
