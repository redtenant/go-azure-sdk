package devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions() UpdateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions {
	return UpdateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions{}
}

func (o UpdateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceCompliancePolicyScheduledActionsForRuleConfiguration - Update the navigation property
// scheduledActionConfigurations in deviceManagement
func (c DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) UpdateDeviceCompliancePolicyScheduledActionsForRuleConfiguration(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId, input beta.DeviceComplianceActionItem, options UpdateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) (result UpdateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationResponse, err error) {
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
