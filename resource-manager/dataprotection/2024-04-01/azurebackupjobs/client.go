package azurebackupjobs

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureBackupJobsClient struct {
	Client *resourcemanager.Client
}

func NewAzureBackupJobsClientWithBaseURI(sdkApi sdkEnv.Api) (*AzureBackupJobsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "azurebackupjobs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AzureBackupJobsClient: %+v", err)
	}

	return &AzureBackupJobsClient{
		Client: client,
	}, nil
}
