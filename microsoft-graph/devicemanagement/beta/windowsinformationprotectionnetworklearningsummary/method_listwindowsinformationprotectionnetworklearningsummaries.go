package windowsinformationprotectionnetworklearningsummary

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

type ListWindowsInformationProtectionNetworkLearningSummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsInformationProtectionNetworkLearningSummary
}

type ListWindowsInformationProtectionNetworkLearningSummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsInformationProtectionNetworkLearningSummary
}

type ListWindowsInformationProtectionNetworkLearningSummariesOperationOptions struct {
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

func DefaultListWindowsInformationProtectionNetworkLearningSummariesOperationOptions() ListWindowsInformationProtectionNetworkLearningSummariesOperationOptions {
	return ListWindowsInformationProtectionNetworkLearningSummariesOperationOptions{}
}

func (o ListWindowsInformationProtectionNetworkLearningSummariesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListWindowsInformationProtectionNetworkLearningSummariesOperationOptions) ToOData() *odata.Query {
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

func (o ListWindowsInformationProtectionNetworkLearningSummariesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListWindowsInformationProtectionNetworkLearningSummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsInformationProtectionNetworkLearningSummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsInformationProtectionNetworkLearningSummaries - Get windowsInformationProtectionNetworkLearningSummaries
// from deviceManagement. The windows information protection network learning summaries.
func (c WindowsInformationProtectionNetworkLearningSummaryClient) ListWindowsInformationProtectionNetworkLearningSummaries(ctx context.Context, options ListWindowsInformationProtectionNetworkLearningSummariesOperationOptions) (result ListWindowsInformationProtectionNetworkLearningSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListWindowsInformationProtectionNetworkLearningSummariesCustomPager{},
		Path:          "/deviceManagement/windowsInformationProtectionNetworkLearningSummaries",
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
		Values *[]beta.WindowsInformationProtectionNetworkLearningSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsInformationProtectionNetworkLearningSummariesComplete retrieves all the results into a single object
func (c WindowsInformationProtectionNetworkLearningSummaryClient) ListWindowsInformationProtectionNetworkLearningSummariesComplete(ctx context.Context, options ListWindowsInformationProtectionNetworkLearningSummariesOperationOptions) (ListWindowsInformationProtectionNetworkLearningSummariesCompleteResult, error) {
	return c.ListWindowsInformationProtectionNetworkLearningSummariesCompleteMatchingPredicate(ctx, options, WindowsInformationProtectionNetworkLearningSummaryOperationPredicate{})
}

// ListWindowsInformationProtectionNetworkLearningSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsInformationProtectionNetworkLearningSummaryClient) ListWindowsInformationProtectionNetworkLearningSummariesCompleteMatchingPredicate(ctx context.Context, options ListWindowsInformationProtectionNetworkLearningSummariesOperationOptions, predicate WindowsInformationProtectionNetworkLearningSummaryOperationPredicate) (result ListWindowsInformationProtectionNetworkLearningSummariesCompleteResult, err error) {
	items := make([]beta.WindowsInformationProtectionNetworkLearningSummary, 0)

	resp, err := c.ListWindowsInformationProtectionNetworkLearningSummaries(ctx, options)
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

	result = ListWindowsInformationProtectionNetworkLearningSummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
