package userexperienceanalyticsanomalydevice

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsAnomalyDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsAnomalyDevice
}

type CreateUserExperienceAnalyticsAnomalyDeviceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsAnomalyDeviceOperationOptions() CreateUserExperienceAnalyticsAnomalyDeviceOperationOptions {
	return CreateUserExperienceAnalyticsAnomalyDeviceOperationOptions{}
}

func (o CreateUserExperienceAnalyticsAnomalyDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsAnomalyDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsAnomalyDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsAnomalyDevice - Create new navigation property to userExperienceAnalyticsAnomalyDevice
// for deviceManagement
func (c UserExperienceAnalyticsAnomalyDeviceClient) CreateUserExperienceAnalyticsAnomalyDevice(ctx context.Context, input beta.UserExperienceAnalyticsAnomalyDevice, options CreateUserExperienceAnalyticsAnomalyDeviceOperationOptions) (result CreateUserExperienceAnalyticsAnomalyDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsAnomalyDevice",
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

	var model beta.UserExperienceAnalyticsAnomalyDevice
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
