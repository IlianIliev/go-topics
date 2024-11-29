package interfaces

import (
	"encoding/json"
	"net/http"
	"to-topics/interfaces/services/users"
)

type createUserInput struct {
	Name     string
	Email    string
	Password string
}

func createUserHandler(usersService users.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var input createUserInput

		// Unmarshal the request body into the createUserInput struct
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validate input
		if input.Name == "" || input.Email == "" || input.Password == "" {
			http.Error(w, "name, email and password are required", http.StatusBadRequest)
			return
		}

		// Save the user to the database
		err = usersService.CreateUser(users.CreateUserInput{
			Name:     input.Name,
			Email:    input.Email,
			Password: input.Password,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send confirmation email

		w.WriteHeader(http.StatusCreated)
	}
}
