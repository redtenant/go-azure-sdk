package domainnamereference

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainNameReferenceClient struct {
	Client *msgraph.Client
}

func NewDomainNameReferenceClientWithBaseURI(sdkApi sdkEnv.Api) (*DomainNameReferenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "domainnamereference", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DomainNameReferenceClient: %+v", err)
	}

	return &DomainNameReferenceClient{
		Client: client,
	}, nil
}
