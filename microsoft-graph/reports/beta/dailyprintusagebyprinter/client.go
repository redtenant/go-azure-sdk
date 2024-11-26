package dailyprintusagebyprinter

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DailyPrintUsageByPrinterClient struct {
	Client *msgraph.Client
}

func NewDailyPrintUsageByPrinterClientWithBaseURI(sdkApi sdkEnv.Api) (*DailyPrintUsageByPrinterClient, error) {
	client, err := msgraph.NewClient(sdkApi, "dailyprintusagebyprinter", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DailyPrintUsageByPrinterClient: %+v", err)
	}

	return &DailyPrintUsageByPrinterClient{
		Client: client,
	}, nil
}
