package userexperienceanalyticsbatteryhealthdeviceappimpact

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsBatteryHealthDeviceAppImpact
}

type CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions() CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions {
	return CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions{}
}

func (o CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpact - Create new navigation property to
// userExperienceAnalyticsBatteryHealthDeviceAppImpact for deviceManagement
func (c UserExperienceAnalyticsBatteryHealthDeviceAppImpactClient) CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpact(ctx context.Context, input beta.UserExperienceAnalyticsBatteryHealthDeviceAppImpact, options CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions) (result CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationResponse, err error) {
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
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceAppImpact",
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

	var model beta.UserExperienceAnalyticsBatteryHealthDeviceAppImpact
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
