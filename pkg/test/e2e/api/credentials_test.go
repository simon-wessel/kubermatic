//go:build e2e

/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"context"
	"sort"
	"testing"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils"

	"k8s.io/apimachinery/pkg/api/equality"
)

func TestListCredentials(t *testing.T) {
	tests := []struct {
		name         string
		provider     string
		datacenter   string
		expectedList []string
	}{
		{
			name:         "test, get DigitalOcean credential names",
			provider:     "digitalocean",
			expectedList: []string{"e2e-digitalocean"},
		},
		{
			name:         "test, get Azure credential names",
			provider:     "azure",
			expectedList: []string{"e2e-azure"},
		},
		{
			name:         "test, get OpenStack credential names",
			provider:     "openstack",
			expectedList: []string{"e2e-openstack"},
		},
		{
			name:         "test, get GCP credential names",
			provider:     "gcp",
			expectedList: []string{"e2e-gcp"},
		},
		{
			name:         "test, get GCP credential names for the specific datacenter",
			provider:     "gcp",
			datacenter:   "gcp-westeurope",
			expectedList: []string{"e2e-gcp", "e2e-gcp-datacenter"},
		},
	}

	ctx := context.Background()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			masterToken, err := utils.RetrieveMasterToken(ctx)
			if err != nil {
				t.Fatalf("failed to get master token: %v", err)
			}

			testClient := utils.NewTestClient(masterToken, t)
			credentialList, err := testClient.ListCredentials(tc.provider, tc.datacenter)
			if err != nil {
				t.Fatalf("failed to get credential names for provider %s: %v", tc.provider, err)
			}
			sort.Strings(tc.expectedList)
			sort.Strings(credentialList)
			if !equality.Semantic.DeepEqual(tc.expectedList, credentialList) {
				t.Fatalf("expected: %v, got %v", tc.expectedList, credentialList)
			}
		})
	}
}

func TestAzureSizesWithCredentials(t *testing.T) {
	tests := []struct {
		name           string
		credentialName string
		location       string
	}{
		{
			name:           "test, get Azure VM sizes",
			credentialName: "e2e-azure",
			location:       "westeurope",
		},
	}
	ctx := context.Background()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			masterToken, err := utils.RetrieveMasterToken(ctx)
			if err != nil {
				t.Fatalf("failed to get master token: %v", err)
			}

			testClient := utils.NewTestClient(masterToken, t)
			if _, err := testClient.ListAzureSizes(tc.credentialName, tc.location); err != nil {
				t.Fatalf("failed to get Azure size list: %v", err)
			}
		})
	}
}

func TestDOSizesWithCredentials(t *testing.T) {
	tests := []struct {
		name           string
		credentialName string
	}{
		{
			name:           "test, get DigitalOcean VM sizes",
			credentialName: "e2e-digitalocean",
		},
	}
	ctx := context.Background()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			masterToken, err := utils.RetrieveMasterToken(ctx)
			if err != nil {
				t.Fatalf("failed to get master token: %v", err)
			}

			testClient := utils.NewTestClient(masterToken, t)
			if _, err := testClient.ListDOSizes(tc.credentialName); err != nil {
				t.Fatalf("failed to get DO size list: %v", err)
			}
		})
	}
}
