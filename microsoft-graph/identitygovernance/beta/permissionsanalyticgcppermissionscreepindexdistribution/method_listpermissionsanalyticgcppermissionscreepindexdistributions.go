package permissionsanalyticgcppermissionscreepindexdistribution

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PermissionsCreepIndexDistribution
}

type ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PermissionsCreepIndexDistribution
}

type ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationOptions() ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationOptions {
	return ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationOptions{}
}

func (o ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionsAnalyticGcpPermissionsCreepIndexDistributions - Get permissionsCreepIndexDistributions from
// identityGovernance. Represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart
// shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium,
// high).
func (c PermissionsAnalyticGcpPermissionsCreepIndexDistributionClient) ListPermissionsAnalyticGcpPermissionsCreepIndexDistributions(ctx context.Context, options ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationOptions) (result ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsCustomPager{},
		Path:          "/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions",
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]beta.PermissionsCreepIndexDistribution `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsComplete retrieves all the results into a single object
func (c PermissionsAnalyticGcpPermissionsCreepIndexDistributionClient) ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsComplete(ctx context.Context, options ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationOptions) (ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsCompleteResult, error) {
	return c.ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsCompleteMatchingPredicate(ctx, options, PermissionsCreepIndexDistributionOperationPredicate{})
}

// ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionsAnalyticGcpPermissionsCreepIndexDistributionClient) ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsCompleteMatchingPredicate(ctx context.Context, options ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsOperationOptions, predicate PermissionsCreepIndexDistributionOperationPredicate) (result ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsCompleteResult, err error) {
	items := make([]beta.PermissionsCreepIndexDistribution, 0)

	resp, err := c.ListPermissionsAnalyticGcpPermissionsCreepIndexDistributions(ctx, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListPermissionsAnalyticGcpPermissionsCreepIndexDistributionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
