package appplatform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationLiveViewsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationLiveViewResource
}

type ApplicationLiveViewsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApplicationLiveViewResource
}

type ApplicationLiveViewsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ApplicationLiveViewsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ApplicationLiveViewsList ...
func (c AppPlatformClient) ApplicationLiveViewsList(ctx context.Context, id commonids.SpringCloudServiceId) (result ApplicationLiveViewsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ApplicationLiveViewsListCustomPager{},
		Path:       fmt.Sprintf("%s/applicationLiveViews", id.ID()),
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
		Values *[]ApplicationLiveViewResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ApplicationLiveViewsListComplete retrieves all the results into a single object
func (c AppPlatformClient) ApplicationLiveViewsListComplete(ctx context.Context, id commonids.SpringCloudServiceId) (ApplicationLiveViewsListCompleteResult, error) {
	return c.ApplicationLiveViewsListCompleteMatchingPredicate(ctx, id, ApplicationLiveViewResourceOperationPredicate{})
}

// ApplicationLiveViewsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppPlatformClient) ApplicationLiveViewsListCompleteMatchingPredicate(ctx context.Context, id commonids.SpringCloudServiceId, predicate ApplicationLiveViewResourceOperationPredicate) (result ApplicationLiveViewsListCompleteResult, err error) {
	items := make([]ApplicationLiveViewResource, 0)

	resp, err := c.ApplicationLiveViewsList(ctx, id)
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

	result = ApplicationLiveViewsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
