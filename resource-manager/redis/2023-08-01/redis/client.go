package redis

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RedisClient struct {
	Client *resourcemanager.Client
}

func NewRedisClientWithBaseURI(sdkApi sdkEnv.Api) (*RedisClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "redis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RedisClient: %+v", err)
	}

	return &RedisClient{
		Client: client,
	}, nil
}
