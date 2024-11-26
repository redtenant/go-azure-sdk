package informationprotectionbitlocker

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionBitlockerClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionBitlockerClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionBitlockerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectionbitlocker", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionBitlockerClient: %+v", err)
	}

	return &InformationProtectionBitlockerClient{
		Client: client,
	}, nil
}
