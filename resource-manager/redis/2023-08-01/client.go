package v2023_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/redis/2023-08-01/aad"
	"github.com/redtenant/go-azure-sdk/resource-manager/redis/2023-08-01/firewallrules"
	"github.com/redtenant/go-azure-sdk/resource-manager/redis/2023-08-01/patchschedules"
	"github.com/redtenant/go-azure-sdk/resource-manager/redis/2023-08-01/privateendpointconnections"
	"github.com/redtenant/go-azure-sdk/resource-manager/redis/2023-08-01/privatelinkresources"
	"github.com/redtenant/go-azure-sdk/resource-manager/redis/2023-08-01/redis"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	AAD                        *aad.AADClient
	FirewallRules              *firewallrules.FirewallRulesClient
	PatchSchedules             *patchschedules.PatchSchedulesClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	Redis                      *redis.RedisClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	aADClient, err := aad.NewAADClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AAD client: %+v", err)
	}
	configureFunc(aADClient.Client)

	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FirewallRules client: %+v", err)
	}
	configureFunc(firewallRulesClient.Client)

	patchSchedulesClient, err := patchschedules.NewPatchSchedulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PatchSchedules client: %+v", err)
	}
	configureFunc(patchSchedulesClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	redisClient, err := redis.NewRedisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Redis client: %+v", err)
	}
	configureFunc(redisClient.Client)

	return &Client{
		AAD:                        aADClient,
		FirewallRules:              firewallRulesClient,
		PatchSchedules:             patchSchedulesClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		Redis:                      redisClient,
	}, nil
}
