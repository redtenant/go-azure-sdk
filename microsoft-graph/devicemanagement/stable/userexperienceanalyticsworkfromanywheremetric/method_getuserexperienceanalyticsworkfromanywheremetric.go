package userexperienceanalyticsworkfromanywheremetric

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsWorkFromAnywhereMetricOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsWorkFromAnywhereMetric
}

type GetUserExperienceAnalyticsWorkFromAnywhereMetricOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetUserExperienceAnalyticsWorkFromAnywhereMetricOperationOptions() GetUserExperienceAnalyticsWorkFromAnywhereMetricOperationOptions {
	return GetUserExperienceAnalyticsWorkFromAnywhereMetricOperationOptions{}
}

func (o GetUserExperienceAnalyticsWorkFromAnywhereMetricOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsWorkFromAnywhereMetricOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetUserExperienceAnalyticsWorkFromAnywhereMetricOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsWorkFromAnywhereMetric - Get userExperienceAnalyticsWorkFromAnywhereMetrics from
// deviceManagement. User experience analytics work from anywhere metrics.
func (c UserExperienceAnalyticsWorkFromAnywhereMetricClient) GetUserExperienceAnalyticsWorkFromAnywhereMetric(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId, options GetUserExperienceAnalyticsWorkFromAnywhereMetricOperationOptions) (result GetUserExperienceAnalyticsWorkFromAnywhereMetricOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model stable.UserExperienceAnalyticsWorkFromAnywhereMetric
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
