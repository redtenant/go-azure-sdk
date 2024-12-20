package devicecompliancepolicydevicestatusoverview

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

type GetDeviceCompliancePolicyDeviceStatusOverviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceComplianceDeviceOverview
}

type GetDeviceCompliancePolicyDeviceStatusOverviewOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetDeviceCompliancePolicyDeviceStatusOverviewOperationOptions() GetDeviceCompliancePolicyDeviceStatusOverviewOperationOptions {
	return GetDeviceCompliancePolicyDeviceStatusOverviewOperationOptions{}
}

func (o GetDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) ToOData() *odata.Query {
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

func (o GetDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceCompliancePolicyDeviceStatusOverview - Get deviceStatusOverview from deviceManagement. Device compliance
// devices status overview
func (c DeviceCompliancePolicyDeviceStatusOverviewClient) GetDeviceCompliancePolicyDeviceStatusOverview(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyId, options GetDeviceCompliancePolicyDeviceStatusOverviewOperationOptions) (result GetDeviceCompliancePolicyDeviceStatusOverviewOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceStatusOverview", id.ID()),
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

	var model beta.DeviceComplianceDeviceOverview
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
