package userexperienceanalyticsapphealthdeviceperformancedetail

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions() UpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions {
	return UpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetail - Update the navigation property
// userExperienceAnalyticsAppHealthDevicePerformanceDetails in deviceManagement
func (c UserExperienceAnalyticsAppHealthDevicePerformanceDetailClient) UpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetail(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId, input stable.UserExperienceAnalyticsAppHealthDevicePerformanceDetails, options UpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions) (result UpdateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationResponse, err error) {
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
