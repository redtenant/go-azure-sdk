package chromeosonboardingsetting

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

type ListChromeOSOnboardingSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ChromeOSOnboardingSettings
}

type ListChromeOSOnboardingSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ChromeOSOnboardingSettings
}

type ListChromeOSOnboardingSettingsOperationOptions struct {
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

func DefaultListChromeOSOnboardingSettingsOperationOptions() ListChromeOSOnboardingSettingsOperationOptions {
	return ListChromeOSOnboardingSettingsOperationOptions{}
}

func (o ListChromeOSOnboardingSettingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListChromeOSOnboardingSettingsOperationOptions) ToOData() *odata.Query {
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

func (o ListChromeOSOnboardingSettingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListChromeOSOnboardingSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListChromeOSOnboardingSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListChromeOSOnboardingSettings - Get chromeOSOnboardingSettings from deviceManagement. Collection of
// ChromeOSOnboardingSettings settings associated with account.
func (c ChromeOSOnboardingSettingClient) ListChromeOSOnboardingSettings(ctx context.Context, options ListChromeOSOnboardingSettingsOperationOptions) (result ListChromeOSOnboardingSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListChromeOSOnboardingSettingsCustomPager{},
		Path:          "/deviceManagement/chromeOSOnboardingSettings",
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
		Values *[]beta.ChromeOSOnboardingSettings `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListChromeOSOnboardingSettingsComplete retrieves all the results into a single object
func (c ChromeOSOnboardingSettingClient) ListChromeOSOnboardingSettingsComplete(ctx context.Context, options ListChromeOSOnboardingSettingsOperationOptions) (ListChromeOSOnboardingSettingsCompleteResult, error) {
	return c.ListChromeOSOnboardingSettingsCompleteMatchingPredicate(ctx, options, ChromeOSOnboardingSettingsOperationPredicate{})
}

// ListChromeOSOnboardingSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChromeOSOnboardingSettingClient) ListChromeOSOnboardingSettingsCompleteMatchingPredicate(ctx context.Context, options ListChromeOSOnboardingSettingsOperationOptions, predicate ChromeOSOnboardingSettingsOperationPredicate) (result ListChromeOSOnboardingSettingsCompleteResult, err error) {
	items := make([]beta.ChromeOSOnboardingSettings, 0)

	resp, err := c.ListChromeOSOnboardingSettings(ctx, options)
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

	result = ListChromeOSOnboardingSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
