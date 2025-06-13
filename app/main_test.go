package main

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
)

type mockDynamoClient struct{}

func (m *mockDynamoClient) GetItem(ctx context.Context, input *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return &dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"secret_code": &types.AttributeValueMemberS{Value: "test-secret"},
		},
	}, nil
}

func TestGetSecret(t *testing.T) {
	mockClient := &mockDynamoClient{}
	code, err := getSecret(context.Background(), mockClient, "any-table")
	assert.NoError(t, err)
	assert.Equal(t, "test-secret", code)
}
