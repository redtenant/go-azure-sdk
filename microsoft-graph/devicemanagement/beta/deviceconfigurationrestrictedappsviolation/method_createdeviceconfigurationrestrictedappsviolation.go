package deviceconfigurationrestrictedappsviolation

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDeviceConfigurationRestrictedAppsViolationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.RestrictedAppsViolation
}

type CreateDeviceConfigurationRestrictedAppsViolationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDeviceConfigurationRestrictedAppsViolationOperationOptions() CreateDeviceConfigurationRestrictedAppsViolationOperationOptions {
	return CreateDeviceConfigurationRestrictedAppsViolationOperationOptions{}
}

func (o CreateDeviceConfigurationRestrictedAppsViolationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceConfigurationRestrictedAppsViolationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceConfigurationRestrictedAppsViolationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceConfigurationRestrictedAppsViolation - Create new navigation property to
// deviceConfigurationRestrictedAppsViolations for deviceManagement
func (c DeviceConfigurationRestrictedAppsViolationClient) CreateDeviceConfigurationRestrictedAppsViolation(ctx context.Context, input beta.RestrictedAppsViolation, options CreateDeviceConfigurationRestrictedAppsViolationOperationOptions) (result CreateDeviceConfigurationRestrictedAppsViolationOperationResponse, err error) {
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
		Path:          "/deviceManagement/deviceConfigurationRestrictedAppsViolations",
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

	var model beta.RestrictedAppsViolation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
