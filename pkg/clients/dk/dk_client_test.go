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

package dk

import (
	"context"
	"fmt"
)

// MockDeveloperKnowledgeClient is a mock implementation for testing.
type MockDeveloperKnowledgeClient struct{}

func (m *MockDeveloperKnowledgeClient) GetDocuments(_ context.Context, documentIDs []string) (string, error) {
	return fmt.Sprintf("Mock documents for IDs: %v", documentIDs), nil
}

func (m *MockDeveloperKnowledgeClient) AnswerQuery(_ context.Context, query string) (string, error) {
	return fmt.Sprintf("Mock answer for query: %s", query), nil
}

func (m *MockDeveloperKnowledgeClient) SearchDocuments(_ context.Context, query string) (string, error) {
	return fmt.Sprintf("Mock search results for query: %s", query), nil
}
