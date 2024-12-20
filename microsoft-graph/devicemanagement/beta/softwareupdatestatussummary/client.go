package softwareupdatestatussummary

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareUpdateStatusSummaryClient struct {
	Client *msgraph.Client
}

func NewSoftwareUpdateStatusSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*SoftwareUpdateStatusSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "softwareupdatestatussummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SoftwareUpdateStatusSummaryClient: %+v", err)
	}

	return &SoftwareUpdateStatusSummaryClient{
		Client: client,
	}, nil
}
