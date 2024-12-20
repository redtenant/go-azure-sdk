package calendar

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

type GetCalendarSchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ScheduleInformation
}

type GetCalendarSchedulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ScheduleInformation
}

type GetCalendarSchedulesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultGetCalendarSchedulesOperationOptions() GetCalendarSchedulesOperationOptions {
	return GetCalendarSchedulesOperationOptions{}
}

func (o GetCalendarSchedulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCalendarSchedulesOperationOptions) ToOData() *odata.Query {
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

func (o GetCalendarSchedulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetCalendarSchedulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetCalendarSchedulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetCalendarSchedules - Invoke action getSchedule. Get the free/busy availability information for a collection of
// users, distributions lists, or resources (rooms or equipment) for a specified time period.
func (c CalendarClient) GetCalendarSchedules(ctx context.Context, id beta.GroupId, input GetCalendarSchedulesRequest, options GetCalendarSchedulesOperationOptions) (result GetCalendarSchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetCalendarSchedulesCustomPager{},
		Path:          fmt.Sprintf("%s/calendar/getSchedule", id.ID()),
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
		Values *[]beta.ScheduleInformation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetCalendarSchedulesComplete retrieves all the results into a single object
func (c CalendarClient) GetCalendarSchedulesComplete(ctx context.Context, id beta.GroupId, input GetCalendarSchedulesRequest, options GetCalendarSchedulesOperationOptions) (GetCalendarSchedulesCompleteResult, error) {
	return c.GetCalendarSchedulesCompleteMatchingPredicate(ctx, id, input, options, ScheduleInformationOperationPredicate{})
}

// GetCalendarSchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarClient) GetCalendarSchedulesCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, input GetCalendarSchedulesRequest, options GetCalendarSchedulesOperationOptions, predicate ScheduleInformationOperationPredicate) (result GetCalendarSchedulesCompleteResult, err error) {
	items := make([]beta.ScheduleInformation, 0)

	resp, err := c.GetCalendarSchedules(ctx, id, input, options)
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

	result = GetCalendarSchedulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
