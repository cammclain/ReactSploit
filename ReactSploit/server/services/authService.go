package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func load_config() {
	config := godotenv.Load()
	if config != nil {
		panic("Error loading .env file for auth service")
	}
	// fmt.Println("Loaded config")
	msfHost := os.Getenv("MSF_HOST")
	if msfHost == "" {
		panic("MSF_HOST environment variable not set")
	}
}

func AuthenticateWithMSF(credentials Credentials) (bool, error) {
	// Construct the request to msfrpc
	load_config()

	data, _ := json.Marshal(map[string]string{
		"username": credentials.Username,
		"password": credentials.Password,
	})
	msfHost := os.Getenv("MSF_HOST")
	resp, err := http.Post("http://"+msfHost+":55553/api/v1/auth/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Check if the status code is 200
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}
