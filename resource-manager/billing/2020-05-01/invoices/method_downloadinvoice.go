package invoices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/client/pollers"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DownloadInvoiceOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DownloadURL
}

type DownloadInvoiceOperationOptions struct {
	DownloadToken *string
}

func DefaultDownloadInvoiceOperationOptions() DownloadInvoiceOperationOptions {
	return DownloadInvoiceOperationOptions{}
}

func (o DownloadInvoiceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DownloadInvoiceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DownloadInvoiceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DownloadToken != nil {
		out.Append("downloadToken", fmt.Sprintf("%v", *o.DownloadToken))
	}
	return &out
}

// DownloadInvoice ...
func (c InvoicesClient) DownloadInvoice(ctx context.Context, id BillingAccountInvoiceId, options DownloadInvoiceOperationOptions) (result DownloadInvoiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/download", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// DownloadInvoiceThenPoll performs DownloadInvoice then polls until it's completed
func (c InvoicesClient) DownloadInvoiceThenPoll(ctx context.Context, id BillingAccountInvoiceId, options DownloadInvoiceOperationOptions) error {
	result, err := c.DownloadInvoice(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing DownloadInvoice: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DownloadInvoice: %+v", err)
	}

	return nil
}
