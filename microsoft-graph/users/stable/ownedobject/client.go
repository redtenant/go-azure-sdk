package ownedobject

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OwnedObjectClient struct {
	Client *msgraph.Client
}

func NewOwnedObjectClientWithBaseURI(sdkApi sdkEnv.Api) (*OwnedObjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "ownedobject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OwnedObjectClient: %+v", err)
	}

	return &OwnedObjectClient{
		Client: client,
	}, nil
}
