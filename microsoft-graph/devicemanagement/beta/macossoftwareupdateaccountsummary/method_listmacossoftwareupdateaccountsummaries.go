package macossoftwareupdateaccountsummary

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

type ListMacOSSoftwareUpdateAccountSummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MacOSSoftwareUpdateAccountSummary
}

type ListMacOSSoftwareUpdateAccountSummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MacOSSoftwareUpdateAccountSummary
}

type ListMacOSSoftwareUpdateAccountSummariesOperationOptions struct {
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

func DefaultListMacOSSoftwareUpdateAccountSummariesOperationOptions() ListMacOSSoftwareUpdateAccountSummariesOperationOptions {
	return ListMacOSSoftwareUpdateAccountSummariesOperationOptions{}
}

func (o ListMacOSSoftwareUpdateAccountSummariesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMacOSSoftwareUpdateAccountSummariesOperationOptions) ToOData() *odata.Query {
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

func (o ListMacOSSoftwareUpdateAccountSummariesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMacOSSoftwareUpdateAccountSummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMacOSSoftwareUpdateAccountSummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMacOSSoftwareUpdateAccountSummaries - Get macOSSoftwareUpdateAccountSummaries from deviceManagement. The MacOS
// software update account summaries for this account.
func (c MacOSSoftwareUpdateAccountSummaryClient) ListMacOSSoftwareUpdateAccountSummaries(ctx context.Context, options ListMacOSSoftwareUpdateAccountSummariesOperationOptions) (result ListMacOSSoftwareUpdateAccountSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMacOSSoftwareUpdateAccountSummariesCustomPager{},
		Path:          "/deviceManagement/macOSSoftwareUpdateAccountSummaries",
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
		Values *[]beta.MacOSSoftwareUpdateAccountSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMacOSSoftwareUpdateAccountSummariesComplete retrieves all the results into a single object
func (c MacOSSoftwareUpdateAccountSummaryClient) ListMacOSSoftwareUpdateAccountSummariesComplete(ctx context.Context, options ListMacOSSoftwareUpdateAccountSummariesOperationOptions) (ListMacOSSoftwareUpdateAccountSummariesCompleteResult, error) {
	return c.ListMacOSSoftwareUpdateAccountSummariesCompleteMatchingPredicate(ctx, options, MacOSSoftwareUpdateAccountSummaryOperationPredicate{})
}

// ListMacOSSoftwareUpdateAccountSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MacOSSoftwareUpdateAccountSummaryClient) ListMacOSSoftwareUpdateAccountSummariesCompleteMatchingPredicate(ctx context.Context, options ListMacOSSoftwareUpdateAccountSummariesOperationOptions, predicate MacOSSoftwareUpdateAccountSummaryOperationPredicate) (result ListMacOSSoftwareUpdateAccountSummariesCompleteResult, err error) {
	items := make([]beta.MacOSSoftwareUpdateAccountSummary, 0)

	resp, err := c.ListMacOSSoftwareUpdateAccountSummaries(ctx, options)
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

	result = ListMacOSSoftwareUpdateAccountSummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
