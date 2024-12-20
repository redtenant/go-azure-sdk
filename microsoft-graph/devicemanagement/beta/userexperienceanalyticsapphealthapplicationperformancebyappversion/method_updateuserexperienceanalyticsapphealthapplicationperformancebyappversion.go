package userexperienceanalyticsapphealthapplicationperformancebyappversion

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionOperationOptions() UpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionOperationOptions {
	return UpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion - Update the navigation property
// userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion in deviceManagement
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient) UpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId, input beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion, options UpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionOperationOptions) (result UpdateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	return
}
