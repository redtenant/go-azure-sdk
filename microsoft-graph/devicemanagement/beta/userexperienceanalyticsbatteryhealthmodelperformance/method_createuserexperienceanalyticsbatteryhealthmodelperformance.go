package userexperienceanalyticsbatteryhealthmodelperformance

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsBatteryHealthModelPerformanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsBatteryHealthModelPerformance
}

type CreateUserExperienceAnalyticsBatteryHealthModelPerformanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsBatteryHealthModelPerformanceOperationOptions() CreateUserExperienceAnalyticsBatteryHealthModelPerformanceOperationOptions {
	return CreateUserExperienceAnalyticsBatteryHealthModelPerformanceOperationOptions{}
}

func (o CreateUserExperienceAnalyticsBatteryHealthModelPerformanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsBatteryHealthModelPerformanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsBatteryHealthModelPerformanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsBatteryHealthModelPerformance - Create new navigation property to
// userExperienceAnalyticsBatteryHealthModelPerformance for deviceManagement
func (c UserExperienceAnalyticsBatteryHealthModelPerformanceClient) CreateUserExperienceAnalyticsBatteryHealthModelPerformance(ctx context.Context, input beta.UserExperienceAnalyticsBatteryHealthModelPerformance, options CreateUserExperienceAnalyticsBatteryHealthModelPerformanceOperationOptions) (result CreateUserExperienceAnalyticsBatteryHealthModelPerformanceOperationResponse, err error) {
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
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthModelPerformance",
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

	var model beta.UserExperienceAnalyticsBatteryHealthModelPerformance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
