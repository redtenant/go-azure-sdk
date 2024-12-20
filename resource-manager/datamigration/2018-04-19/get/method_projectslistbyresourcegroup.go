package get

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectsListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Project
}

type ProjectsListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Project
}

type ProjectsListByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ProjectsListByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ProjectsListByResourceGroup ...
func (c GETClient) ProjectsListByResourceGroup(ctx context.Context, id ServiceId) (result ProjectsListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ProjectsListByResourceGroupCustomPager{},
		Path:       fmt.Sprintf("%s/projects", id.ID()),
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
		Values *[]Project `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ProjectsListByResourceGroupComplete retrieves all the results into a single object
func (c GETClient) ProjectsListByResourceGroupComplete(ctx context.Context, id ServiceId) (ProjectsListByResourceGroupCompleteResult, error) {
	return c.ProjectsListByResourceGroupCompleteMatchingPredicate(ctx, id, ProjectOperationPredicate{})
}

// ProjectsListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GETClient) ProjectsListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id ServiceId, predicate ProjectOperationPredicate) (result ProjectsListByResourceGroupCompleteResult, err error) {
	items := make([]Project, 0)

	resp, err := c.ProjectsListByResourceGroup(ctx, id)
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

	result = ProjectsListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
