package rolemanagementalertalert

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

type ListRoleManagementAlertAlertsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleManagementAlert
}

type ListRoleManagementAlertAlertsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleManagementAlert
}

type ListRoleManagementAlertAlertsOperationOptions struct {
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

func DefaultListRoleManagementAlertAlertsOperationOptions() ListRoleManagementAlertAlertsOperationOptions {
	return ListRoleManagementAlertAlertsOperationOptions{}
}

func (o ListRoleManagementAlertAlertsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRoleManagementAlertAlertsOperationOptions) ToOData() *odata.Query {
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

func (o ListRoleManagementAlertAlertsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRoleManagementAlertAlertsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleManagementAlertAlertsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleManagementAlertAlerts - List alerts. Get a list of the unifiedRoleManagementAlert objects and their
// properties.
func (c RoleManagementAlertAlertClient) ListRoleManagementAlertAlerts(ctx context.Context, options ListRoleManagementAlertAlertsOperationOptions) (result ListRoleManagementAlertAlertsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRoleManagementAlertAlertsCustomPager{},
		Path:          "/identityGovernance/roleManagementAlerts/alerts",
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
		Values *[]beta.UnifiedRoleManagementAlert `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementAlertAlertsComplete retrieves all the results into a single object
func (c RoleManagementAlertAlertClient) ListRoleManagementAlertAlertsComplete(ctx context.Context, options ListRoleManagementAlertAlertsOperationOptions) (ListRoleManagementAlertAlertsCompleteResult, error) {
	return c.ListRoleManagementAlertAlertsCompleteMatchingPredicate(ctx, options, UnifiedRoleManagementAlertOperationPredicate{})
}

// ListRoleManagementAlertAlertsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementAlertAlertClient) ListRoleManagementAlertAlertsCompleteMatchingPredicate(ctx context.Context, options ListRoleManagementAlertAlertsOperationOptions, predicate UnifiedRoleManagementAlertOperationPredicate) (result ListRoleManagementAlertAlertsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleManagementAlert, 0)

	resp, err := c.ListRoleManagementAlertAlerts(ctx, options)
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

	result = ListRoleManagementAlertAlertsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
