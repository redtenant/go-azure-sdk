package backuprestore

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupRestoreClient struct {
	Client *resourcemanager.Client
}

func NewBackupRestoreClientWithBaseURI(sdkApi sdkEnv.Api) (*BackupRestoreClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "backuprestore", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BackupRestoreClient: %+v", err)
	}

	return &BackupRestoreClient{
		Client: client,
	}, nil
}
