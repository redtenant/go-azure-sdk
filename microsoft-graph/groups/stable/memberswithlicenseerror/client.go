package memberswithlicenseerror

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembersWithLicenseErrorClient struct {
	Client *msgraph.Client
}

func NewMembersWithLicenseErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*MembersWithLicenseErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "memberswithlicenseerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MembersWithLicenseErrorClient: %+v", err)
	}

	return &MembersWithLicenseErrorClient{
		Client: client,
	}, nil
}
