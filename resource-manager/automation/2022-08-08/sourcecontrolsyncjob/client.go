package sourcecontrolsyncjob

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlSyncJobClient struct {
	Client *resourcemanager.Client
}

func NewSourceControlSyncJobClientWithBaseURI(sdkApi sdkEnv.Api) (*SourceControlSyncJobClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sourcecontrolsyncjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SourceControlSyncJobClient: %+v", err)
	}

	return &SourceControlSyncJobClient{
		Client: client,
	}, nil
}
