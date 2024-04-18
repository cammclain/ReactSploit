package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthenticateWithMSF(credentials Credentials) (bool, error) {
	// Construct the request to msfrpc
	data, _ := json.Marshal(map[string]string{
		"username": credentials.Username,
		"password": credentials.Password,
	})
	resp, err := http.Post("http://localhost:55553/api/v1/auth/login", "application/json", bytes.NewBuffer(data))
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
