package deviceenrollmentconfigurationassignment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDeviceEnrollmentConfigurationAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.EnrollmentConfigurationAssignment
}

type ListDeviceEnrollmentConfigurationAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.EnrollmentConfigurationAssignment
}

type ListDeviceEnrollmentConfigurationAssignmentsOperationOptions struct {
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

func DefaultListDeviceEnrollmentConfigurationAssignmentsOperationOptions() ListDeviceEnrollmentConfigurationAssignmentsOperationOptions {
	return ListDeviceEnrollmentConfigurationAssignmentsOperationOptions{}
}

func (o ListDeviceEnrollmentConfigurationAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceEnrollmentConfigurationAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceEnrollmentConfigurationAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceEnrollmentConfigurationAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceEnrollmentConfigurationAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceEnrollmentConfigurationAssignments - List enrollmentConfigurationAssignments. List properties and
// relationships of the enrollmentConfigurationAssignment objects.
func (c DeviceEnrollmentConfigurationAssignmentClient) ListDeviceEnrollmentConfigurationAssignments(ctx context.Context, id stable.DeviceManagementDeviceEnrollmentConfigurationId, options ListDeviceEnrollmentConfigurationAssignmentsOperationOptions) (result ListDeviceEnrollmentConfigurationAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceEnrollmentConfigurationAssignmentsCustomPager{},
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
		Values *[]stable.EnrollmentConfigurationAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceEnrollmentConfigurationAssignmentsComplete retrieves all the results into a single object
func (c DeviceEnrollmentConfigurationAssignmentClient) ListDeviceEnrollmentConfigurationAssignmentsComplete(ctx context.Context, id stable.DeviceManagementDeviceEnrollmentConfigurationId, options ListDeviceEnrollmentConfigurationAssignmentsOperationOptions) (ListDeviceEnrollmentConfigurationAssignmentsCompleteResult, error) {
	return c.ListDeviceEnrollmentConfigurationAssignmentsCompleteMatchingPredicate(ctx, id, options, EnrollmentConfigurationAssignmentOperationPredicate{})
}

// ListDeviceEnrollmentConfigurationAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceEnrollmentConfigurationAssignmentClient) ListDeviceEnrollmentConfigurationAssignmentsCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementDeviceEnrollmentConfigurationId, options ListDeviceEnrollmentConfigurationAssignmentsOperationOptions, predicate EnrollmentConfigurationAssignmentOperationPredicate) (result ListDeviceEnrollmentConfigurationAssignmentsCompleteResult, err error) {
	items := make([]stable.EnrollmentConfigurationAssignment, 0)

	resp, err := c.ListDeviceEnrollmentConfigurationAssignments(ctx, id, options)
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

	result = ListDeviceEnrollmentConfigurationAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
