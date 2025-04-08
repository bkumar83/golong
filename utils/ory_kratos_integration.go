// Ory Kratos integration service
package utils

import (
    "context"
    kratos "github.com/ory/kratos-client-go"
    "log"
    "errors"
)

var kratosClient *kratos.APIClient

// Initialize Kratos Client
func InitKratosClient(kratosURL string) {
    cfg := kratos.NewConfiguration()
    cfg.Servers = kratos.ServerConfigurations{{URL: kratosURL}}
    kratosClient = kratos.NewAPIClient(cfg)
    log.Println("Ory Kratos client initialized with URL:", kratosURL)
}

// Validate Kratos Session
func ValidateKratosSession(token string) (bool, error) {
    if kratosClient == nil {
        log.Println("Kratos client not initialized")
        return false, errors.New("Kratos client not initialized")
    }

    session, _, err := kratosClient.FrontendAPI.ToSession(context.Background()).XSessionToken(token).Execute()
    if err != nil {
        log.Println("Error validating Kratos session:", err)
        return false, err
    }

    if session.Active != nil && *session.Active {
        log.Println("Valid Kratos session for:", session.Identity.Id)
        return true, nil
    }

    log.Println("Invalid Kratos session")
    return false, errors.New("Invalid session")
}
