package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
)

type SecretResponse struct {
	SecretCode string `json:"secret_code"`
}

type HealthResponse struct {
	Status    string `json:"status"`
	Container string `json:"container"`
	Project   string `json:"project"`
}

type DynamoDBAPI interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

func getSecret(ctx context.Context, client DynamoDBAPI, tableName string) (string, error) {
	resp, err := client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			"code_name": &types.AttributeValueMemberS{Value: "thedoctor"},
		},
	})
	if err != nil {
		return "", err
	}
	code := resp.Item["secret_code"].(*types.AttributeValueMemberS).Value
	return code, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	r := gin.Default()
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)
	tableName := os.Getenv("DYNAMODB_TABLE")

	r.GET("/secret", func(c *gin.Context) {
		code, err := getSecret(ctx, dynamoClient, tableName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, SecretResponse{SecretCode: code})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, HealthResponse{
			Status:    "healthy",
			Project: getEnv("PROJECT_URL", "https://github.com/PatentTest/devops-challenge"),
			Container: getEnv("CONTAINER_URL", "https://hub.docker.com/r/PatentTest/devops-challenge"),
		})
	})

	r.Run("0.0.0.0:5000")
}
