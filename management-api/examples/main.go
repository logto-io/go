package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/logto-io/go/management-api/client"
	"github.com/logto-io/go/management-api/logto"
)

const (
	endpoint   = "http://localhost:3001"
	appId      = "<app_id>"
	appSecret  = "<app_secret>"
	testUserId = "<user_id>"
)

func main() {
	client := client.NewClient(&client.Config{
		AppId:     appId,
		AppSecret: appSecret,
		Endpoint:  endpoint,
		Debug:     true,
	}, nil)

	// Example: Get user information
	req := client.UsersAPI.GetUser(client.Context, testUserId)
	result, resp, err := req.Execute()
	if err != nil {
		log.Fatalf("API call failed: %v", err)
	}

	fmt.Printf("Response status code: %d\n", resp.StatusCode)
	jsonBody, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("Response body: %s\n", string(jsonBody))

	{
		// Update user email
		newEmail := fmt.Sprintf("%s@change.me", result.Id)

		req := client.UsersAPI.UpdateUser(client.Context, testUserId)
		var updateUserRequestPayload logto.UpdateUserRequest
		updateUserRequestPayload.SetPrimaryEmail(logto.UpdateUserRequestPrimaryEmail{String: &newEmail})
		req = req.UpdateUserRequest(updateUserRequestPayload)
		result, resp, err := req.Execute()
		if err != nil {
			log.Fatalf("API call failed: %v", err)
		}
		fmt.Printf("Response status code: %d\n", resp.StatusCode)
		jsonBody, _ := json.MarshalIndent(result, "", "  ")
		fmt.Printf("Response body: %s\n", string(jsonBody))
	}
}
