package monthlyprintusagesummariesbyprinter

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonthlyPrintUsageSummariesByPrinterClient struct {
	Client *msgraph.Client
}

func NewMonthlyPrintUsageSummariesByPrinterClientWithBaseURI(sdkApi sdkEnv.Api) (*MonthlyPrintUsageSummariesByPrinterClient, error) {
	client, err := msgraph.NewClient(sdkApi, "monthlyprintusagesummariesbyprinter", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MonthlyPrintUsageSummariesByPrinterClient: %+v", err)
	}

	return &MonthlyPrintUsageSummariesByPrinterClient{
		Client: client,
	}, nil
}
