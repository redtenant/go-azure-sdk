package userexperienceanalyticsapphealthoverviewmetricvalue

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsAppHealthOverviewMetricValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsMetric
}

type GetUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions() GetUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions {
	return GetUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions{}
}

func (o GetUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions) ToOData() *odata.Query {
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

func (o GetUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsAppHealthOverviewMetricValue - Get metricValues from deviceManagement. The metric values
// for the user experience analytics category. Read-only.
func (c UserExperienceAnalyticsAppHealthOverviewMetricValueClient) GetUserExperienceAnalyticsAppHealthOverviewMetricValue(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId, options GetUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions) (result GetUserExperienceAnalyticsAppHealthOverviewMetricValueOperationResponse, err error) {
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

	var model beta.UserExperienceAnalyticsMetric
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
