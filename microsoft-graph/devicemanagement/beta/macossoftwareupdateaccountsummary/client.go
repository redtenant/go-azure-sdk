package macossoftwareupdateaccountsummary

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateAccountSummaryClient struct {
	Client *msgraph.Client
}

func NewMacOSSoftwareUpdateAccountSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*MacOSSoftwareUpdateAccountSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "macossoftwareupdateaccountsummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MacOSSoftwareUpdateAccountSummaryClient: %+v", err)
	}

	return &MacOSSoftwareUpdateAccountSummaryClient{
		Client: client,
	}, nil
}
