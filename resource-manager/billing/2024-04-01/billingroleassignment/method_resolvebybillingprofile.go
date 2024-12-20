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

type ResolveByBillingProfileOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRoleAssignment
}

type ResolveByBillingProfileCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleAssignment
}

type ResolveByBillingProfileOperationOptions struct {
	Filter                   *string
	ResolveScopeDisplayNames *bool
}

func DefaultResolveByBillingProfileOperationOptions() ResolveByBillingProfileOperationOptions {
	return ResolveByBillingProfileOperationOptions{}
}

func (o ResolveByBillingProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResolveByBillingProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ResolveByBillingProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.ResolveScopeDisplayNames != nil {
		out.Append("resolveScopeDisplayNames", fmt.Sprintf("%v", *o.ResolveScopeDisplayNames))
	}
	return &out
}

type ResolveByBillingProfileCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResolveByBillingProfileCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResolveByBillingProfile ...
func (c BillingRoleAssignmentClient) ResolveByBillingProfile(ctx context.Context, id BillingProfileId, options ResolveByBillingProfileOperationOptions) (result ResolveByBillingProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ResolveByBillingProfileCustomPager{},
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

// ResolveByBillingProfileThenPoll performs ResolveByBillingProfile then polls until it's completed
func (c BillingRoleAssignmentClient) ResolveByBillingProfileThenPoll(ctx context.Context, id BillingProfileId, options ResolveByBillingProfileOperationOptions) error {
	result, err := c.ResolveByBillingProfile(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing ResolveByBillingProfile: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ResolveByBillingProfile: %+v", err)
	}

	return nil
}
