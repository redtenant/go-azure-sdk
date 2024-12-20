package followedsite

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

type RemoveFollowedSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Site
}

type RemoveFollowedSitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Site
}

type RemoveFollowedSitesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultRemoveFollowedSitesOperationOptions() RemoveFollowedSitesOperationOptions {
	return RemoveFollowedSitesOperationOptions{}
}

func (o RemoveFollowedSitesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveFollowedSitesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o RemoveFollowedSitesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type RemoveFollowedSitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *RemoveFollowedSitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RemoveFollowedSites - Invoke action remove. Unfollow a user's site or multiple sites.
func (c FollowedSiteClient) RemoveFollowedSites(ctx context.Context, id stable.UserId, input RemoveFollowedSitesRequest, options RemoveFollowedSitesOperationOptions) (result RemoveFollowedSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &RemoveFollowedSitesCustomPager{},
		Path:          fmt.Sprintf("%s/followedSites/remove", id.ID()),
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
		Values *[]stable.Site `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RemoveFollowedSitesComplete retrieves all the results into a single object
func (c FollowedSiteClient) RemoveFollowedSitesComplete(ctx context.Context, id stable.UserId, input RemoveFollowedSitesRequest, options RemoveFollowedSitesOperationOptions) (RemoveFollowedSitesCompleteResult, error) {
	return c.RemoveFollowedSitesCompleteMatchingPredicate(ctx, id, input, options, SiteOperationPredicate{})
}

// RemoveFollowedSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FollowedSiteClient) RemoveFollowedSitesCompleteMatchingPredicate(ctx context.Context, id stable.UserId, input RemoveFollowedSitesRequest, options RemoveFollowedSitesOperationOptions, predicate SiteOperationPredicate) (result RemoveFollowedSitesCompleteResult, err error) {
	items := make([]stable.Site, 0)

	resp, err := c.RemoveFollowedSites(ctx, id, input, options)
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

	result = RemoveFollowedSitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
