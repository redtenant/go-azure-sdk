package mobiledevicemanagementpolicyincludedgroup

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

type ListMobileDeviceManagementPolicyIncludedGroupRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListMobileDeviceManagementPolicyIncludedGroupRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions struct {
	Count     *bool
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Skip      *int64
	Top       *int64
}

func DefaultListMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions() ListMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions {
	return ListMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions{}
}

func (o ListMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
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
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMobileDeviceManagementPolicyIncludedGroupRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileDeviceManagementPolicyIncludedGroupRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileDeviceManagementPolicyIncludedGroupRefs - List includedGroups. Get the list of groups that are included in
// a mobile device management policy.
func (c MobileDeviceManagementPolicyIncludedGroupClient) ListMobileDeviceManagementPolicyIncludedGroupRefs(ctx context.Context, id beta.PolicyMobileDeviceManagementPolicyId, options ListMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions) (result ListMobileDeviceManagementPolicyIncludedGroupRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMobileDeviceManagementPolicyIncludedGroupRefsCustomPager{},
		Path:          fmt.Sprintf("%s/includedGroups/$ref", id.ID()),
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

	temp := make([]beta.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListMobileDeviceManagementPolicyIncludedGroupRefsComplete retrieves all the results into a single object
func (c MobileDeviceManagementPolicyIncludedGroupClient) ListMobileDeviceManagementPolicyIncludedGroupRefsComplete(ctx context.Context, id beta.PolicyMobileDeviceManagementPolicyId, options ListMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions) (ListMobileDeviceManagementPolicyIncludedGroupRefsCompleteResult, error) {
	return c.ListMobileDeviceManagementPolicyIncludedGroupRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListMobileDeviceManagementPolicyIncludedGroupRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileDeviceManagementPolicyIncludedGroupClient) ListMobileDeviceManagementPolicyIncludedGroupRefsCompleteMatchingPredicate(ctx context.Context, id beta.PolicyMobileDeviceManagementPolicyId, options ListMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListMobileDeviceManagementPolicyIncludedGroupRefsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListMobileDeviceManagementPolicyIncludedGroupRefs(ctx, id, options)
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

	result = ListMobileDeviceManagementPolicyIncludedGroupRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
