package serverdevopsaudit

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerDevOpsAuditClient struct {
	Client *resourcemanager.Client
}

func NewServerDevOpsAuditClientWithBaseURI(sdkApi sdkEnv.Api) (*ServerDevOpsAuditClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "serverdevopsaudit", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerDevOpsAuditClient: %+v", err)
	}

	return &ServerDevOpsAuditClient{
		Client: client,
	}, nil
}
