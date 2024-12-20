package androidforworkappconfigurationschema

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkAppConfigurationSchemaClient struct {
	Client *msgraph.Client
}

func NewAndroidForWorkAppConfigurationSchemaClientWithBaseURI(sdkApi sdkEnv.Api) (*AndroidForWorkAppConfigurationSchemaClient, error) {
	client, err := msgraph.NewClient(sdkApi, "androidforworkappconfigurationschema", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AndroidForWorkAppConfigurationSchemaClient: %+v", err)
	}

	return &AndroidForWorkAppConfigurationSchemaClient{
		Client: client,
	}, nil
}
