package accessreviewdefinitioninstancedefinition

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceDefinitionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceDefinitionClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceDefinitionClient{
		Client: client,
	}, nil
}
