package registrationassignments

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistrationAssignmentsClient struct {
	Client *resourcemanager.Client
}

func NewRegistrationAssignmentsClientWithBaseURI(sdkApi sdkEnv.Api) (*RegistrationAssignmentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "registrationassignments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RegistrationAssignmentsClient: %+v", err)
	}

	return &RegistrationAssignmentsClient{
		Client: client,
	}, nil
}
