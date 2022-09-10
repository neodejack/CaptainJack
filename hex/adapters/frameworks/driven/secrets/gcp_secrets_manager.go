package secrets

import (
	"context"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// Constants

const (
	ProjectId string = "captainjack"
)

type AdapterGCP struct {
}

func NewAdapterGCP() *AdapterGCP {
	return &AdapterGCP{}
}

func (secrets AdapterGCP) GetTeleToken() ([]byte, error) {

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	defer client.Close()

	// Build the request.
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: "CaptainJack",
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
	}

	teleBotToken := result.Payload.Data

	return teleBotToken, err
}

