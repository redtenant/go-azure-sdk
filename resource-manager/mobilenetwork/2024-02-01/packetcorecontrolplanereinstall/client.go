package packetcorecontrolplanereinstall

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreControlPlaneReinstallClient struct {
	Client *resourcemanager.Client
}

func NewPacketCoreControlPlaneReinstallClientWithBaseURI(sdkApi sdkEnv.Api) (*PacketCoreControlPlaneReinstallClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "packetcorecontrolplanereinstall", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PacketCoreControlPlaneReinstallClient: %+v", err)
	}

	return &PacketCoreControlPlaneReinstallClient{
		Client: client,
	}, nil
}
