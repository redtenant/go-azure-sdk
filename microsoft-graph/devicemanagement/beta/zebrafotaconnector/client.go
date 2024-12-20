package zebrafotaconnector

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaConnectorClient struct {
	Client *msgraph.Client
}

func NewZebraFotaConnectorClientWithBaseURI(sdkApi sdkEnv.Api) (*ZebraFotaConnectorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "zebrafotaconnector", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ZebraFotaConnectorClient: %+v", err)
	}

	return &ZebraFotaConnectorClient{
		Client: client,
	}, nil
}
