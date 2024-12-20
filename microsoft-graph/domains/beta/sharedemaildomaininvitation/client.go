package sharedemaildomaininvitation

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedEmailDomainInvitationClient struct {
	Client *msgraph.Client
}

func NewSharedEmailDomainInvitationClientWithBaseURI(sdkApi sdkEnv.Api) (*SharedEmailDomainInvitationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sharedemaildomaininvitation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SharedEmailDomainInvitationClient: %+v", err)
	}

	return &SharedEmailDomainInvitationClient{
		Client: client,
	}, nil
}
