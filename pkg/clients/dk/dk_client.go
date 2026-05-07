// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package dk provides the Developer Knowledge API client.
package dk

import (
	"context"
	"fmt"
)

// DeveloperKnowledgeClient defines the interface for interacting with the Developer Knowledge API.
type DeveloperKnowledgeClient interface {
	GetDocuments(ctx context.Context, documentIDs []string) (string, error)
	AnswerQuery(ctx context.Context, query string) (string, error)
	SearchDocuments(ctx context.Context, query string) (string, error)
}

// RealDeveloperKnowledgeClient is the actual implementation (stubbed for now).
type RealDeveloperKnowledgeClient struct {
	// Add configuration fields here (e.g., API key, base URL)
}

// NewRealDeveloperKnowledgeClient creates a new real client instance.
func NewRealDeveloperKnowledgeClient() *RealDeveloperKnowledgeClient {
	return &RealDeveloperKnowledgeClient{}
}

// GetDocuments fetches specific documents by their IDs.
func (c *RealDeveloperKnowledgeClient) GetDocuments(_ context.Context, _ []string) (string, error) {
	return "", fmt.Errorf("GetDocuments not implemented")
}

// AnswerQuery answers a query based on the knowledge base.
func (c *RealDeveloperKnowledgeClient) AnswerQuery(_ context.Context, _ string) (string, error) {
	return "", fmt.Errorf("AnswerQuery not implemented")
}

// SearchDocuments searches for documents related to a query.
func (c *RealDeveloperKnowledgeClient) SearchDocuments(_ context.Context, _ string) (string, error) {
	return "", fmt.Errorf("SearchDocuments not implemented")
}
