package informationprotectionpolicylabel

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

type ListInformationProtectionPolicyLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.InformationProtectionLabel
}

type ListInformationProtectionPolicyLabelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.InformationProtectionLabel
}

type ListInformationProtectionPolicyLabelsOperationOptions struct {
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

func DefaultListInformationProtectionPolicyLabelsOperationOptions() ListInformationProtectionPolicyLabelsOperationOptions {
	return ListInformationProtectionPolicyLabelsOperationOptions{}
}

func (o ListInformationProtectionPolicyLabelsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListInformationProtectionPolicyLabelsOperationOptions) ToOData() *odata.Query {
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

func (o ListInformationProtectionPolicyLabelsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListInformationProtectionPolicyLabelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInformationProtectionPolicyLabelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInformationProtectionPolicyLabels - Get labels from users
func (c InformationProtectionPolicyLabelClient) ListInformationProtectionPolicyLabels(ctx context.Context, id beta.UserId, options ListInformationProtectionPolicyLabelsOperationOptions) (result ListInformationProtectionPolicyLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListInformationProtectionPolicyLabelsCustomPager{},
		Path:          fmt.Sprintf("%s/informationProtection/policy/labels", id.ID()),
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
		Values *[]beta.InformationProtectionLabel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInformationProtectionPolicyLabelsComplete retrieves all the results into a single object
func (c InformationProtectionPolicyLabelClient) ListInformationProtectionPolicyLabelsComplete(ctx context.Context, id beta.UserId, options ListInformationProtectionPolicyLabelsOperationOptions) (ListInformationProtectionPolicyLabelsCompleteResult, error) {
	return c.ListInformationProtectionPolicyLabelsCompleteMatchingPredicate(ctx, id, options, InformationProtectionLabelOperationPredicate{})
}

// ListInformationProtectionPolicyLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InformationProtectionPolicyLabelClient) ListInformationProtectionPolicyLabelsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListInformationProtectionPolicyLabelsOperationOptions, predicate InformationProtectionLabelOperationPredicate) (result ListInformationProtectionPolicyLabelsCompleteResult, err error) {
	items := make([]beta.InformationProtectionLabel, 0)

	resp, err := c.ListInformationProtectionPolicyLabels(ctx, id, options)
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

	result = ListInformationProtectionPolicyLabelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
