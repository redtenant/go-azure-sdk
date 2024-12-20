package workspacemanagedsqlserverserverencryptionprotector

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerServerEncryptionProtectorClient struct {
	Client *resourcemanager.Client
}

func NewWorkspaceManagedSqlServerServerEncryptionProtectorClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkspaceManagedSqlServerServerEncryptionProtectorClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "workspacemanagedsqlserverserverencryptionprotector", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceManagedSqlServerServerEncryptionProtectorClient: %+v", err)
	}

	return &WorkspaceManagedSqlServerServerEncryptionProtectorClient{
		Client: client,
	}, nil
}
