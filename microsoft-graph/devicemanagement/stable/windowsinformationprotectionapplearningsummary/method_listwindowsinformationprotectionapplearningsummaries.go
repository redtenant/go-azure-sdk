package windowsinformationprotectionapplearningsummary

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

type ListWindowsInformationProtectionAppLearningSummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.WindowsInformationProtectionAppLearningSummary
}

type ListWindowsInformationProtectionAppLearningSummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.WindowsInformationProtectionAppLearningSummary
}

type ListWindowsInformationProtectionAppLearningSummariesOperationOptions struct {
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

func DefaultListWindowsInformationProtectionAppLearningSummariesOperationOptions() ListWindowsInformationProtectionAppLearningSummariesOperationOptions {
	return ListWindowsInformationProtectionAppLearningSummariesOperationOptions{}
}

func (o ListWindowsInformationProtectionAppLearningSummariesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListWindowsInformationProtectionAppLearningSummariesOperationOptions) ToOData() *odata.Query {
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

func (o ListWindowsInformationProtectionAppLearningSummariesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListWindowsInformationProtectionAppLearningSummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsInformationProtectionAppLearningSummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsInformationProtectionAppLearningSummaries - List windowsInformationProtectionAppLearningSummaries. List
// properties and relationships of the windowsInformationProtectionAppLearningSummary objects.
func (c WindowsInformationProtectionAppLearningSummaryClient) ListWindowsInformationProtectionAppLearningSummaries(ctx context.Context, options ListWindowsInformationProtectionAppLearningSummariesOperationOptions) (result ListWindowsInformationProtectionAppLearningSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListWindowsInformationProtectionAppLearningSummariesCustomPager{},
		Path:          "/deviceManagement/windowsInformationProtectionAppLearningSummaries",
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
		Values *[]stable.WindowsInformationProtectionAppLearningSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsInformationProtectionAppLearningSummariesComplete retrieves all the results into a single object
func (c WindowsInformationProtectionAppLearningSummaryClient) ListWindowsInformationProtectionAppLearningSummariesComplete(ctx context.Context, options ListWindowsInformationProtectionAppLearningSummariesOperationOptions) (ListWindowsInformationProtectionAppLearningSummariesCompleteResult, error) {
	return c.ListWindowsInformationProtectionAppLearningSummariesCompleteMatchingPredicate(ctx, options, WindowsInformationProtectionAppLearningSummaryOperationPredicate{})
}

// ListWindowsInformationProtectionAppLearningSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsInformationProtectionAppLearningSummaryClient) ListWindowsInformationProtectionAppLearningSummariesCompleteMatchingPredicate(ctx context.Context, options ListWindowsInformationProtectionAppLearningSummariesOperationOptions, predicate WindowsInformationProtectionAppLearningSummaryOperationPredicate) (result ListWindowsInformationProtectionAppLearningSummariesCompleteResult, err error) {
	items := make([]stable.WindowsInformationProtectionAppLearningSummary, 0)

	resp, err := c.ListWindowsInformationProtectionAppLearningSummaries(ctx, options)
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

	result = ListWindowsInformationProtectionAppLearningSummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
