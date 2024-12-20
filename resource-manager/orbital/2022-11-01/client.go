package v2022_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/orbital/2022-11-01/contact"
	"github.com/redtenant/go-azure-sdk/resource-manager/orbital/2022-11-01/contactprofile"
	"github.com/redtenant/go-azure-sdk/resource-manager/orbital/2022-11-01/groundstation"
	"github.com/redtenant/go-azure-sdk/resource-manager/orbital/2022-11-01/spacecraft"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	Contact        *contact.ContactClient
	ContactProfile *contactprofile.ContactProfileClient
	GroundStation  *groundstation.GroundStationClient
	Spacecraft     *spacecraft.SpacecraftClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	contactClient, err := contact.NewContactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Contact client: %+v", err)
	}
	configureFunc(contactClient.Client)

	contactProfileClient, err := contactprofile.NewContactProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContactProfile client: %+v", err)
	}
	configureFunc(contactProfileClient.Client)

	groundStationClient, err := groundstation.NewGroundStationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroundStation client: %+v", err)
	}
	configureFunc(groundStationClient.Client)

	spacecraftClient, err := spacecraft.NewSpacecraftClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Spacecraft client: %+v", err)
	}
	configureFunc(spacecraftClient.Client)

	return &Client{
		Contact:        contactClient,
		ContactProfile: contactProfileClient,
		GroundStation:  groundStationClient,
		Spacecraft:     spacecraftClient,
	}, nil
}
