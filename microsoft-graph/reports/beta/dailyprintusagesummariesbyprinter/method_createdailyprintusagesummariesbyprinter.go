package dailyprintusagesummariesbyprinter

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDailyPrintUsageSummariesByPrinterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PrintUsageByPrinter
}

type CreateDailyPrintUsageSummariesByPrinterOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDailyPrintUsageSummariesByPrinterOperationOptions() CreateDailyPrintUsageSummariesByPrinterOperationOptions {
	return CreateDailyPrintUsageSummariesByPrinterOperationOptions{}
}

func (o CreateDailyPrintUsageSummariesByPrinterOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDailyPrintUsageSummariesByPrinterOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDailyPrintUsageSummariesByPrinterOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDailyPrintUsageSummariesByPrinter - Create new navigation property to dailyPrintUsageSummariesByPrinter for
// reports
func (c DailyPrintUsageSummariesByPrinterClient) CreateDailyPrintUsageSummariesByPrinter(ctx context.Context, input beta.PrintUsageByPrinter, options CreateDailyPrintUsageSummariesByPrinterOperationOptions) (result CreateDailyPrintUsageSummariesByPrinterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/reports/dailyPrintUsageSummariesByPrinter",
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.PrintUsageByPrinter
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
