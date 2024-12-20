package drivelastmodifiedbyuser

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveLastModifiedByUserClient: %+v", err)
	}

	return &DriveLastModifiedByUserClient{
		Client: client,
	}, nil
}
