package billingroleassignment

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

type ResolveByInvoiceSectionOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRoleAssignment
}

type ResolveByInvoiceSectionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleAssignment
}

type ResolveByInvoiceSectionOperationOptions struct {
	Filter                   *string
	ResolveScopeDisplayNames *bool
}

func DefaultResolveByInvoiceSectionOperationOptions() ResolveByInvoiceSectionOperationOptions {
	return ResolveByInvoiceSectionOperationOptions{}
}

func (o ResolveByInvoiceSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResolveByInvoiceSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ResolveByInvoiceSectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.ResolveScopeDisplayNames != nil {
		out.Append("resolveScopeDisplayNames", fmt.Sprintf("%v", *o.ResolveScopeDisplayNames))
	}
	return &out
}

type ResolveByInvoiceSectionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResolveByInvoiceSectionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResolveByInvoiceSection ...
func (c BillingRoleAssignmentClient) ResolveByInvoiceSection(ctx context.Context, id InvoiceSectionId, options ResolveByInvoiceSectionOperationOptions) (result ResolveByInvoiceSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ResolveByInvoiceSectionCustomPager{},
		Path:          fmt.Sprintf("%s/resolveBillingRoleAssignments", id.ID()),
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

// ResolveByInvoiceSectionThenPoll performs ResolveByInvoiceSection then polls until it's completed
func (c BillingRoleAssignmentClient) ResolveByInvoiceSectionThenPoll(ctx context.Context, id InvoiceSectionId, options ResolveByInvoiceSectionOperationOptions) error {
	result, err := c.ResolveByInvoiceSection(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing ResolveByInvoiceSection: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ResolveByInvoiceSection: %+v", err)
	}

	return nil
}
