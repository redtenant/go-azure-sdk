package devicecompliancepolicyuserstatusoverview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceCompliancePolicyUserStatusOverviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceCompliancePolicyUserStatusOverviewOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDeviceCompliancePolicyUserStatusOverviewOperationOptions() UpdateDeviceCompliancePolicyUserStatusOverviewOperationOptions {
	return UpdateDeviceCompliancePolicyUserStatusOverviewOperationOptions{}
}

func (o UpdateDeviceCompliancePolicyUserStatusOverviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceCompliancePolicyUserStatusOverviewOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceCompliancePolicyUserStatusOverviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceCompliancePolicyUserStatusOverview - Update deviceComplianceUserOverview. Update the properties of a
// deviceComplianceUserOverview object.
func (c DeviceCompliancePolicyUserStatusOverviewClient) UpdateDeviceCompliancePolicyUserStatusOverview(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyId, input stable.DeviceComplianceUserOverview, options UpdateDeviceCompliancePolicyUserStatusOverviewOperationOptions) (result UpdateDeviceCompliancePolicyUserStatusOverviewOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/userStatusOverview", id.ID()),
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
