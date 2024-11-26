package reservationtransactions

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationTransactionsClient struct {
	Client *resourcemanager.Client
}

func NewReservationTransactionsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReservationTransactionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "reservationtransactions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReservationTransactionsClient: %+v", err)
	}

	return &ReservationTransactionsClient{
		Client: client,
	}, nil
}
