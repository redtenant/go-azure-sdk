package devicemanagementroledefinitioninheritspermissionsfrom

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

type ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleDefinition
}

type ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleDefinition
}

type ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationOptions struct {
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

func DefaultListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationOptions() ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationOptions {
	return ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationOptions{}
}

func (o ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementRoleDefinitionInheritsPermissionsFroms - Get inheritsPermissionsFrom from roleManagement.
// Read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in
// roles support this attribute.
func (c DeviceManagementRoleDefinitionInheritsPermissionsFromClient) ListDeviceManagementRoleDefinitionInheritsPermissionsFroms(ctx context.Context, id beta.RoleManagementDeviceManagementRoleDefinitionId, options ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationOptions) (result ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCustomPager{},
		Path:          fmt.Sprintf("%s/inheritsPermissionsFrom", id.ID()),
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
		Values *[]beta.UnifiedRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceManagementRoleDefinitionInheritsPermissionsFromsComplete retrieves all the results into a single object
func (c DeviceManagementRoleDefinitionInheritsPermissionsFromClient) ListDeviceManagementRoleDefinitionInheritsPermissionsFromsComplete(ctx context.Context, id beta.RoleManagementDeviceManagementRoleDefinitionId, options ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationOptions) (ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteResult, error) {
	return c.ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx, id, options, UnifiedRoleDefinitionOperationPredicate{})
}

// ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementRoleDefinitionInheritsPermissionsFromClient) ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementDeviceManagementRoleDefinitionId, options ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationOptions, predicate UnifiedRoleDefinitionOperationPredicate) (result ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleDefinition, 0)

	resp, err := c.ListDeviceManagementRoleDefinitionInheritsPermissionsFroms(ctx, id, options)
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

	result = ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
