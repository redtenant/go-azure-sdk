package driverootcreatedbyuserserviceprovisioningerror

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDriveRootCreatedByUserServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServiceProvisioningError
}

type ListDriveRootCreatedByUserServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServiceProvisioningError
}

type ListDriveRootCreatedByUserServiceProvisioningErrorsOperationOptions struct {
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

func DefaultListDriveRootCreatedByUserServiceProvisioningErrorsOperationOptions() ListDriveRootCreatedByUserServiceProvisioningErrorsOperationOptions {
	return ListDriveRootCreatedByUserServiceProvisioningErrorsOperationOptions{}
}

func (o ListDriveRootCreatedByUserServiceProvisioningErrorsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveRootCreatedByUserServiceProvisioningErrorsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveRootCreatedByUserServiceProvisioningErrorsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveRootCreatedByUserServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveRootCreatedByUserServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveRootCreatedByUserServiceProvisioningErrors - Get serviceProvisioningErrors property value. Errors published
// by a federated service describing a nontransient, service-specific error regarding the properties or link from a user
// object. Supports $filter (eq, not, for isResolved and serviceInstance).
func (c DriveRootCreatedByUserServiceProvisioningErrorClient) ListDriveRootCreatedByUserServiceProvisioningErrors(ctx context.Context, id stable.MeDriveId, options ListDriveRootCreatedByUserServiceProvisioningErrorsOperationOptions) (result ListDriveRootCreatedByUserServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveRootCreatedByUserServiceProvisioningErrorsCustomPager{},
		Path:          fmt.Sprintf("%s/root/createdByUser/serviceProvisioningErrors", id.ID()),
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

	temp := make([]stable.ServiceProvisioningError, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalServiceProvisioningErrorImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.ServiceProvisioningError (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDriveRootCreatedByUserServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c DriveRootCreatedByUserServiceProvisioningErrorClient) ListDriveRootCreatedByUserServiceProvisioningErrorsComplete(ctx context.Context, id stable.MeDriveId, options ListDriveRootCreatedByUserServiceProvisioningErrorsOperationOptions) (ListDriveRootCreatedByUserServiceProvisioningErrorsCompleteResult, error) {
	return c.ListDriveRootCreatedByUserServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, options, ServiceProvisioningErrorOperationPredicate{})
}

// ListDriveRootCreatedByUserServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveRootCreatedByUserServiceProvisioningErrorClient) ListDriveRootCreatedByUserServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id stable.MeDriveId, options ListDriveRootCreatedByUserServiceProvisioningErrorsOperationOptions, predicate ServiceProvisioningErrorOperationPredicate) (result ListDriveRootCreatedByUserServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]stable.ServiceProvisioningError, 0)

	resp, err := c.ListDriveRootCreatedByUserServiceProvisioningErrors(ctx, id, options)
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

	result = ListDriveRootCreatedByUserServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
