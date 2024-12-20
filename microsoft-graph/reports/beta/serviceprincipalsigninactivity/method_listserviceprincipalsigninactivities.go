package serviceprincipalsigninactivity

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

type ListServicePrincipalSignInActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServicePrincipalSignInActivity
}

type ListServicePrincipalSignInActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServicePrincipalSignInActivity
}

type ListServicePrincipalSignInActivitiesOperationOptions struct {
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

func DefaultListServicePrincipalSignInActivitiesOperationOptions() ListServicePrincipalSignInActivitiesOperationOptions {
	return ListServicePrincipalSignInActivitiesOperationOptions{}
}

func (o ListServicePrincipalSignInActivitiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListServicePrincipalSignInActivitiesOperationOptions) ToOData() *odata.Query {
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

func (o ListServicePrincipalSignInActivitiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListServicePrincipalSignInActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListServicePrincipalSignInActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListServicePrincipalSignInActivities - List servicePrincipalSignInActivities. Get a list of
// servicePrincipalSignInActivity objects that contains sign-in activity information for service principals in a
// Microsoft Entra tenant. You can use a service principal as a client or resource. A service principal supports
// delegated or app-only authentication context.
func (c ServicePrincipalSignInActivityClient) ListServicePrincipalSignInActivities(ctx context.Context, options ListServicePrincipalSignInActivitiesOperationOptions) (result ListServicePrincipalSignInActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListServicePrincipalSignInActivitiesCustomPager{},
		Path:          "/reports/servicePrincipalSignInActivities",
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
		Values *[]beta.ServicePrincipalSignInActivity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServicePrincipalSignInActivitiesComplete retrieves all the results into a single object
func (c ServicePrincipalSignInActivityClient) ListServicePrincipalSignInActivitiesComplete(ctx context.Context, options ListServicePrincipalSignInActivitiesOperationOptions) (ListServicePrincipalSignInActivitiesCompleteResult, error) {
	return c.ListServicePrincipalSignInActivitiesCompleteMatchingPredicate(ctx, options, ServicePrincipalSignInActivityOperationPredicate{})
}

// ListServicePrincipalSignInActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalSignInActivityClient) ListServicePrincipalSignInActivitiesCompleteMatchingPredicate(ctx context.Context, options ListServicePrincipalSignInActivitiesOperationOptions, predicate ServicePrincipalSignInActivityOperationPredicate) (result ListServicePrincipalSignInActivitiesCompleteResult, err error) {
	items := make([]beta.ServicePrincipalSignInActivity, 0)

	resp, err := c.ListServicePrincipalSignInActivities(ctx, options)
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

	result = ListServicePrincipalSignInActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
