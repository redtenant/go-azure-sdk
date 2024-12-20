package resourceaccessprofileassignment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceAccessProfileAssignmentClient struct {
	Client *msgraph.Client
}

func NewResourceAccessProfileAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*ResourceAccessProfileAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "resourceaccessprofileassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ResourceAccessProfileAssignmentClient: %+v", err)
	}

	return &ResourceAccessProfileAssignmentClient{
		Client: client,
	}, nil
}
