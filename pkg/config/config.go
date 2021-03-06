///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package config

// NO TESTS

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Global contains global configuration variables
var Global Config

// EmptyRegistryAuth == echo -n '{"username":"","password":"","email":""}' | base64
var EmptyRegistryAuth = "eyJ1c2VybmFtZSI6IiIsInBhc3N3b3JkIjoiIiwiZW1haWwiOiIifQ=="

// Identity defines the identity manager specific config
type Identity struct {
	OIDCProvider string   `json:"oidcProvider"`
	ClientID     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
	RedirectURL  string   `json:"redirectUrl"`
	Scopes       []string `json:"scopes"`
}

// Openwhisk defines the OpenWhisk faas specific config
type Openwhisk struct {
	AuthToken string `json:"authToken"`
	Host      string `json:"host"`
}

// OpenFaas defines the OpenFaaS faas specific config
type OpenFaas struct {
	Gateway       string `json:"gateway"`
	K8sConfig     string `json:"k8sConfig"`
	FuncNamespace string `json:"funcNamespace"`
}

// Riff defines the Riff faas specific config
type Riff struct {
	Gateway       string `json:"gateway"`
	K8sConfig     string `json:"k8sConfig"`
	FuncNamespace string `json:"funcNamespace"`
}

// Function defines the function manager specific config
type Function struct {
	Openwhisk        `json:"openwhisk"`
	OpenFaas         `json:"openFaas"`
	Riff             `json:"riff"`
	Faas             string `json:"faas"`
	TemplateDir      string `json:"templateDir"`
	ResyncPeriod     int    `json:"resyncPeriod"`
	FileImageManager string `json:"fileImageManager"`
}

// Registry defines the image registry specific config
type Registry struct {
	RegistryURI  string `json:"uri"`
	RegistryAuth string `json:"auth"`
}

// Config defines global configurations used in Dispatch
type Config struct {
	Identity       `json:"identity"`
	Function       `json:"function"`
	Registry       `json:"registry"`
	OrganizationID string `json:"organizationID"`
}

var defaultConfig = Config{
	Function: Function{
		Faas:         "openfaas",
		TemplateDir:  "images/function-manager/templates",
		ResyncPeriod: 10,
	},
	OrganizationID: "dispatch",
}

// LoadConfiguration loads configurations from a local json file
func LoadConfiguration(file string) Config {
	configFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
	config, err := loadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func loadConfig(reader io.Reader) (Config, error) {
	jsonParser := json.NewDecoder(reader)
	err := jsonParser.Decode(&defaultConfig)
	return defaultConfig, err
}
