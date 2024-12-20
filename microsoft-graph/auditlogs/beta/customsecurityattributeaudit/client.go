package customsecurityattributeaudit

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSecurityAttributeAuditClient struct {
	Client *msgraph.Client
}

func NewCustomSecurityAttributeAuditClientWithBaseURI(sdkApi sdkEnv.Api) (*CustomSecurityAttributeAuditClient, error) {
	client, err := msgraph.NewClient(sdkApi, "customsecurityattributeaudit", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomSecurityAttributeAuditClient: %+v", err)
	}

	return &CustomSecurityAttributeAuditClient{
		Client: client,
	}, nil
}
