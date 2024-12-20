package reusablepolicysettingreferencingconfigurationpolicysetting

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateReusablePolicySettingReferencingConfigurationPolicySettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateReusablePolicySettingReferencingConfigurationPolicySettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateReusablePolicySettingReferencingConfigurationPolicySettingOperationOptions() UpdateReusablePolicySettingReferencingConfigurationPolicySettingOperationOptions {
	return UpdateReusablePolicySettingReferencingConfigurationPolicySettingOperationOptions{}
}

func (o UpdateReusablePolicySettingReferencingConfigurationPolicySettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateReusablePolicySettingReferencingConfigurationPolicySettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateReusablePolicySettingReferencingConfigurationPolicySettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateReusablePolicySettingReferencingConfigurationPolicySetting - Update the navigation property settings in
// deviceManagement
func (c ReusablePolicySettingReferencingConfigurationPolicySettingClient) UpdateReusablePolicySettingReferencingConfigurationPolicySetting(ctx context.Context, id beta.DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId, input beta.DeviceManagementConfigurationSetting, options UpdateReusablePolicySettingReferencingConfigurationPolicySettingOperationOptions) (result UpdateReusablePolicySettingReferencingConfigurationPolicySettingOperationResponse, err error) {
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
