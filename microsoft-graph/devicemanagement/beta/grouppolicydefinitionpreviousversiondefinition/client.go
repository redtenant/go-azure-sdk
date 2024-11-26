package grouppolicydefinitionpreviousversiondefinition

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionPreviousVersionDefinitionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionPreviousVersionDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionPreviousVersionDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionpreviousversiondefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionPreviousVersionDefinitionClient: %+v", err)
	}

	return &GroupPolicyDefinitionPreviousVersionDefinitionClient{
		Client: client,
	}, nil
}
