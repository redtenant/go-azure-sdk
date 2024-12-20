package windowsfeatureupdateprofileassignment

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

type ListWindowsFeatureUpdateProfileAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsFeatureUpdateProfileAssignment
}

type ListWindowsFeatureUpdateProfileAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsFeatureUpdateProfileAssignment
}

type ListWindowsFeatureUpdateProfileAssignmentsOperationOptions struct {
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

func DefaultListWindowsFeatureUpdateProfileAssignmentsOperationOptions() ListWindowsFeatureUpdateProfileAssignmentsOperationOptions {
	return ListWindowsFeatureUpdateProfileAssignmentsOperationOptions{}
}

func (o ListWindowsFeatureUpdateProfileAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListWindowsFeatureUpdateProfileAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListWindowsFeatureUpdateProfileAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListWindowsFeatureUpdateProfileAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsFeatureUpdateProfileAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsFeatureUpdateProfileAssignments - Get assignments from deviceManagement. The list of group assignments of
// the profile.
func (c WindowsFeatureUpdateProfileAssignmentClient) ListWindowsFeatureUpdateProfileAssignments(ctx context.Context, id beta.DeviceManagementWindowsFeatureUpdateProfileId, options ListWindowsFeatureUpdateProfileAssignmentsOperationOptions) (result ListWindowsFeatureUpdateProfileAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListWindowsFeatureUpdateProfileAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.WindowsFeatureUpdateProfileAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsFeatureUpdateProfileAssignmentsComplete retrieves all the results into a single object
func (c WindowsFeatureUpdateProfileAssignmentClient) ListWindowsFeatureUpdateProfileAssignmentsComplete(ctx context.Context, id beta.DeviceManagementWindowsFeatureUpdateProfileId, options ListWindowsFeatureUpdateProfileAssignmentsOperationOptions) (ListWindowsFeatureUpdateProfileAssignmentsCompleteResult, error) {
	return c.ListWindowsFeatureUpdateProfileAssignmentsCompleteMatchingPredicate(ctx, id, options, WindowsFeatureUpdateProfileAssignmentOperationPredicate{})
}

// ListWindowsFeatureUpdateProfileAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsFeatureUpdateProfileAssignmentClient) ListWindowsFeatureUpdateProfileAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementWindowsFeatureUpdateProfileId, options ListWindowsFeatureUpdateProfileAssignmentsOperationOptions, predicate WindowsFeatureUpdateProfileAssignmentOperationPredicate) (result ListWindowsFeatureUpdateProfileAssignmentsCompleteResult, err error) {
	items := make([]beta.WindowsFeatureUpdateProfileAssignment, 0)

	resp, err := c.ListWindowsFeatureUpdateProfileAssignments(ctx, id, options)
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

	result = ListWindowsFeatureUpdateProfileAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
