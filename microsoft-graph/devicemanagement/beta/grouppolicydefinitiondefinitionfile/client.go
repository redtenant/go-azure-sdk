package grouppolicydefinitiondefinitionfile

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionDefinitionFileClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionDefinitionFileClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionDefinitionFileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitiondefinitionfile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionDefinitionFileClient: %+v", err)
	}

	return &GroupPolicyDefinitionDefinitionFileClient{
		Client: client,
	}, nil
}
