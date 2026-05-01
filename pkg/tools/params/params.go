// Copyright 2025 Google LLC
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

// Package params provide common tool parameter types.
package params

import "fmt"

// Project contains the GCP project ID.
type Project struct {
	ProjectID string `json:"project_id" jsonschema:"Required. GCP project ID."`
}

// ProjectIDPath returns the GCP project resource path.
func (p *Project) ProjectIDPath() string {
	return fmt.Sprintf("projects/%s", p.ProjectID)
}

// Location contains the project and GKE cluster location.
type Location struct {
	Project
	Location string `json:"location" jsonschema:"Required. GKE cluster location."`
}

// LocationPath returns the GKE location resource path.
func (l *Location) LocationPath() string {
	return fmt.Sprintf("%s/locations/%s", l.ProjectIDPath(), l.Location)
}

// LocationOptional contains the project and an optional GKE cluster location.
type LocationOptional struct {
	Project
	Location string `json:"location,omitempty" jsonschema:"Optional. GKE cluster location."`
}

// LocationPath returns the GKE location resource path, using "-" for all locations if empty.
func (l *LocationOptional) LocationPath() string {
	if l.Location == "" {
		return fmt.Sprintf("%s/locations/-", l.ProjectIDPath())
	}
	return fmt.Sprintf("%s/locations/%s", l.ProjectIDPath(), l.Location)
}

// Cluster contains the location and GKE cluster name.
type Cluster struct {
	Location
	ClusterName string `json:"cluster_name" jsonschema:"Required. GKE cluster name."`
}

// ClusterPath returns the full GKE cluster resource path.
func (c *Cluster) ClusterPath() string {
	return fmt.Sprintf("%s/clusters/%s", c.LocationPath(), c.ClusterName)
}
