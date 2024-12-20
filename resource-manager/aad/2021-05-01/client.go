package v2021_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/aad/2021-05-01/domainservices"
	"github.com/redtenant/go-azure-sdk/resource-manager/aad/2021-05-01/oucontainer"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	DomainServices *domainservices.DomainServicesClient
	OuContainer    *oucontainer.OuContainerClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	domainServicesClient, err := domainservices.NewDomainServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DomainServices client: %+v", err)
	}
	configureFunc(domainServicesClient.Client)

	ouContainerClient, err := oucontainer.NewOuContainerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OuContainer client: %+v", err)
	}
	configureFunc(ouContainerClient.Client)

	return &Client{
		DomainServices: domainServicesClient,
		OuContainer:    ouContainerClient,
	}, nil
}
