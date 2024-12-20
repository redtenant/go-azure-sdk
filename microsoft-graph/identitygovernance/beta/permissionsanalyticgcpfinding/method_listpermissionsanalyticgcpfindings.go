package permissionsanalyticgcpfinding

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPermissionsAnalyticGcpFindingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Finding
}

type ListPermissionsAnalyticGcpFindingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Finding
}

type ListPermissionsAnalyticGcpFindingsOperationOptions struct {
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

func DefaultListPermissionsAnalyticGcpFindingsOperationOptions() ListPermissionsAnalyticGcpFindingsOperationOptions {
	return ListPermissionsAnalyticGcpFindingsOperationOptions{}
}

func (o ListPermissionsAnalyticGcpFindingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPermissionsAnalyticGcpFindingsOperationOptions) ToOData() *odata.Query {
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

func (o ListPermissionsAnalyticGcpFindingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPermissionsAnalyticGcpFindingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPermissionsAnalyticGcpFindingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPermissionsAnalyticGcpFindings - Get findings from identityGovernance. The output of the permissions usage data
// analysis performed by Permissions Management to assess risk with identities and resources.
func (c PermissionsAnalyticGcpFindingClient) ListPermissionsAnalyticGcpFindings(ctx context.Context, options ListPermissionsAnalyticGcpFindingsOperationOptions) (result ListPermissionsAnalyticGcpFindingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPermissionsAnalyticGcpFindingsCustomPager{},
		Path:          "/identityGovernance/permissionsAnalytics/gcp/findings",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.Finding, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalFindingImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.Finding (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListPermissionsAnalyticGcpFindingsComplete retrieves all the results into a single object
func (c PermissionsAnalyticGcpFindingClient) ListPermissionsAnalyticGcpFindingsComplete(ctx context.Context, options ListPermissionsAnalyticGcpFindingsOperationOptions) (ListPermissionsAnalyticGcpFindingsCompleteResult, error) {
	return c.ListPermissionsAnalyticGcpFindingsCompleteMatchingPredicate(ctx, options, FindingOperationPredicate{})
}

// ListPermissionsAnalyticGcpFindingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionsAnalyticGcpFindingClient) ListPermissionsAnalyticGcpFindingsCompleteMatchingPredicate(ctx context.Context, options ListPermissionsAnalyticGcpFindingsOperationOptions, predicate FindingOperationPredicate) (result ListPermissionsAnalyticGcpFindingsCompleteResult, err error) {
	items := make([]beta.Finding, 0)

	resp, err := c.ListPermissionsAnalyticGcpFindings(ctx, options)
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

	result = ListPermissionsAnalyticGcpFindingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
