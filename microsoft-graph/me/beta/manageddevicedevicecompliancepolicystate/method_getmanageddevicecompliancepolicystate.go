package manageddevicedevicecompliancepolicystate

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetManagedDeviceCompliancePolicyStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceCompliancePolicyState
}

type GetManagedDeviceCompliancePolicyStateOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetManagedDeviceCompliancePolicyStateOperationOptions() GetManagedDeviceCompliancePolicyStateOperationOptions {
	return GetManagedDeviceCompliancePolicyStateOperationOptions{}
}

func (o GetManagedDeviceCompliancePolicyStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetManagedDeviceCompliancePolicyStateOperationOptions) ToOData() *odata.Query {
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

func (o GetManagedDeviceCompliancePolicyStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetManagedDeviceCompliancePolicyState - Get deviceCompliancePolicyStates from me. Device compliance policy states for
// this device.
func (c ManagedDeviceDeviceCompliancePolicyStateClient) GetManagedDeviceCompliancePolicyState(ctx context.Context, id beta.MeManagedDeviceIdDeviceCompliancePolicyStateId, options GetManagedDeviceCompliancePolicyStateOperationOptions) (result GetManagedDeviceCompliancePolicyStateOperationResponse, err error) {
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

	var model beta.DeviceCompliancePolicyState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
