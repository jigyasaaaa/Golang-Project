// main.go
package main

import (
	"net/http"
	"strconv"

	"gofr.dev/pkg/gofr"
)

// User represents a user entity.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Doe"},
}

func main() {
	app := gofr.New()

	// Get all users
	app.GET("/users", func(ctx *gofr.Context) (interface{}, error) {
		err := NewHTTPError(http.StatusNotFound, "Resource not found")
		return nil, err
	})

	// Get a single user by ID
	app.GET("/users/:id", func(ctx *gofr.Context) (interface{}, error) {
		userIDStr := ctx.Param("id")

		// Convert userIDStr to int
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return nil, NewHTTPError(http.StatusBadRequest, "Invalid user ID")
		}

		// Implement logic to find the user by ID
		// For simplicity, we just convert the parameter to an integer
		// and return the corresponding user
		// In a real project, you would typically query a database.
		// Here, we use a simple linear search for demonstration purposes.
		for _, user := range users {
			if user.ID == userID {
				return user, nil
			}
		}
		return nil, NewHTTPError(http.StatusNotFound, "User not found")
	})

	// Create a new user
	app.POST("/users", func(ctx *gofr.Context) (interface{}, error) {
		var newUser User
		if err := ctx.BindJSON(&newUser); err != nil {
			return nil, NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
		}

		// Assign a new ID and add the user to the list
		newUser.ID = len(users) + 1
		users = append(users, newUser)
		return newUser, nil
	})

	// Start the server
	app.Start()
}
