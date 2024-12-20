package androidforworkenrollmentprofile

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnrollmentProfileClient struct {
	Client *msgraph.Client
}

func NewAndroidForWorkEnrollmentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*AndroidForWorkEnrollmentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "androidforworkenrollmentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AndroidForWorkEnrollmentProfileClient: %+v", err)
	}

	return &AndroidForWorkEnrollmentProfileClient{
		Client: client,
	}, nil
}
