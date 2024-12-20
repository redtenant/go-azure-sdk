package backupandexport

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupAndExportClient struct {
	Client *resourcemanager.Client
}

func NewBackupAndExportClientWithBaseURI(sdkApi sdkEnv.Api) (*BackupAndExportClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "backupandexport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BackupAndExportClient: %+v", err)
	}

	return &BackupAndExportClient{
		Client: client,
	}, nil
}
