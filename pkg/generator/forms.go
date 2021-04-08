// Copyright 2018 the Service Broker Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"fmt"
	"log"
	"strings"

	"github.com/cloudfoundry-incubator/cloud-service-broker/pkg/toggles"
	"gopkg.in/yaml.v3"
)

// TileFormsSections holds the top level fields in tile.yml responsible for
// the forms.
// https://docs.pivotal.io/tiledev/2-2/tile-structure.html
type TileFormsSections struct {
	Forms            []Form `yaml:"forms"`
	ServicePlanForms []Form `yaml:"service_plan_forms,omitempty"`
}

// Form is a PCF Ops Manager compatible form definition used to generate forms.
// See https://docs.pivotal.io/tiledev/2-2/product-template-reference.html#form-properties
// for details about the fields.
type Form struct {
	Name        string         `yaml:"name"`
	Label       string         `yaml:"label"`
	Description string         `yaml:"description"`
	Optional    bool           `yaml:"optional,omitempty"` // optional, default false
	Properties  []FormProperty `yaml:"properties"`
}

// FormOption is an enumerated element for FormProperties that can be selected
// from. Name is the value and label is the human-readable display name.
type FormOption struct {
	Name  string `yaml:"name"`
	Label string `yaml:"label"`
}

// FormProperty holds a single form element in a PCF Ops manager form.
type FormProperty struct {
	Name         string       `yaml:"name"`
	Type         string       `yaml:"type"`
	Default      interface{}  `yaml:"default,omitempty"`
	Label        string       `yaml:"label,omitempty"`
	Description  string       `yaml:"description,omitempty"`
	Configurable bool         `yaml:"configurable,omitempty"` // optional, default false
	Options      []FormOption `yaml:"options,omitempty"`
	Optional     bool         `yaml:"optional,omitempty"` // optional, default false
}

// GenerateFormsString creates all the forms for the user to fill out in the PCF tile
// and returns it as a string.
func GenerateFormsString() string {
	response, err := yaml.Marshal(GenerateForms())
	if err != nil {
		log.Fatalf("Error marshaling YAML: %s", err)
	}

	return string(response)
}

// GenerateForms creates all the forms for the user to fill out in the PCF tile.
func GenerateForms() TileFormsSections {
	// Add new forms at the bottom of the list because the order is reflected
	// in the generated UI and we don't want to mix things up on users.
	return TileFormsSections{
		Forms: []Form{
			generateServiceAccountForm(),
			generateDatabaseForm(),
			generateBrokerpakForm(),
			generateFeatureFlagForm(),
		},

		ServicePlanForms: []Form{brokerpakConfigurationForm()},
	}
}

// generateDatabaseForm generates the form for configuring database settings.
func generateDatabaseForm() Form {
	return Form{
		Name:        "database_properties",
		Label:       "Database Properties",
		Description: "Connection details for the backing database for the service broker.",
		Properties: []FormProperty{
			{Name: "db_host", Type: "string", Label: "Database host", Configurable: true},
			{Name: "db_username", Type: "string", Label: "Database username", Optional: true, Configurable: true},
			{Name: "db_password", Type: "secret", Label: "Database password", Optional: true, Configurable: true},
			{Name: "db_port", Type: "string", Label: "Database port (defaults to 3306)", Default: "3306", Configurable: true},
			{Name: "db_name", Type: "string", Label: "Database name", Default: "servicebroker", Configurable: true},
			{Name: "ca_cert", Type: "text", Label: "Server CA cert", Optional: true, Configurable: true},
			{Name: "client_cert", Type: "text", Label: "Client cert", Optional: true, Configurable: true},
			{Name: "client_key", Type: "text", Label: "Client key", Optional: true, Configurable: true},
		},
	}
}

// generateServiceAccountForm generates the form for configuring the service
// account.
func generateServiceAccountForm() Form {
	return Form{
		Name:        "root_service_account",
		Label:       "Root Service Account",
		Description: "Please paste in the contents of the json keyfile (un-encoded) for your service account with owner credentials.",
		Properties: []FormProperty{
			{Name: "root_service_account_json", Type: "text", Label: "Root Service Account JSON", Configurable: true},
		},
	}
}

func generateFeatureFlagForm() Form {
	var formEntries []FormProperty

	for _, toggle := range toggles.Features.Toggles() {
		toggleEntry := FormProperty{
			Name:         strings.ToLower(toggle.EnvironmentVariable()),
			Type:         "boolean",
			Label:        toggle.Name,
			Configurable: true,
			Default:      fmt.Sprintf("%v", toggle.Default), // the tile deals with all values as strings so a default string is acceptable.
			Description:  singleLine(toggle.Description),
		}

		formEntries = append(formEntries, toggleEntry)
	}

	return Form{
		Name:        "features",
		Label:       "Feature Flags",
		Description: "Service broker feature flags.",
		Properties:  formEntries,
	}
}

func generateBrokerpakForm() Form {
	return Form{
		Name:  "brokerpaks",
		Label: "Brokerpaks",
		Description: `Brokerpaks are ways to extend the broker with custom services defined by Terraform templates.
A brokerpak is an archive comprised of a versioned Terraform binary and providers for one or more platform, a manifest, one or more service definitions, and source code.`,
		Properties: []FormProperty{
			{
				Name:         "gsb_brokerpak_config",
				Type:         "text",
				Label:        "Global Brokerpak Configuration",
				Description:  "A JSON map of configuration key/value pairs for all brokerpaks. If a variable isn't found in the specific brokerpak's configuration it's looked up here.",
				Default:      "{}",
				Optional:     false,
				Configurable: true,
			},
		},
	}
}

func brokerpakConfigurationForm() Form {
	return Form{
		Name:        "gsb_brokerpak_sources",
		Description: "Configure Brokerpaks",
		Label:       "Configure Brokerpaks",
		Optional:    true,
		Properties: []FormProperty{
			{
				Name:  "uri",
				Label: "Brokerpak URI",
				Type:  "string",
				Description: `The URI to load. Supported protocols are http, https, gs, and git.
				Cloud Storage (gs) URIs follow the gs://<bucket>/<path> convention and will be read using the service broker service account.

				You can validate the checksum of any file on download by appending a checksum query parameter to the URI in the format type:value.
				Valid checksum types are md5, sha1, sha256 and sha512. e.g. gs://foo/bar.brokerpak?checksum=md5:3063a2c62e82ef8614eee6745a7b6b59`,
				Optional:     false,
				Configurable: true,
			},
			{
				Name:         "service_prefix",
				Label:        "Service Prefix",
				Type:         "string",
				Description:  "A prefix to prepend to every service name. This will be exact, so you may want to include a trailing dash.",
				Optional:     true,
				Configurable: true,
			},
			{
				Name:         "excluded_services",
				Label:        "Excluded Services",
				Type:         "text",
				Description:  "A list of UUIDs of services to exclude, one per line.",
				Optional:     true,
				Configurable: true,
			},
			{
				Name:         "config",
				Label:        "Brokerpak Configuration",
				Type:         "text",
				Description:  "A JSON map of configuration key/value pairs for the brokerpak. If a variable isn't found here, it's looked up in the global config.",
				Default:      "{}",
				Configurable: true,
			},
			{
				Name:         "notes",
				Label:        "Notes",
				Type:         "text",
				Description:  "A place for your notes, not used by the broker.",
				Optional:     true,
				Configurable: true,
			},
		},
	}
}

func singleLine(text string) string {
	lines := strings.Split(text, "\n")

	var out []string
	for _, line := range lines {
		out = append(out, strings.TrimSpace(line))
	}

	return strings.Join(out, " ")
}
