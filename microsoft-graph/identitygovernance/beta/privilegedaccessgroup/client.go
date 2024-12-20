package privilegedaccessgroup

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupClient: %+v", err)
	}

	return &PrivilegedAccessGroupClient{
		Client: client,
	}, nil
}
