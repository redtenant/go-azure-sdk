package entitlementmanagementaccesspackageassignmentaccesspackageassignmentpolicycustomextensionstagesettingcustomextension

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.CustomCalloutExtension
}

type GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions() GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions {
	return GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtension - Get customExtension
// from identityGovernance. Indicates the custom workflow extension that is executed at this stage. Nullable. Supports
// $expand.
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient) GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtension(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId, options GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions) (result GetEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/customExtension", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalCustomCalloutExtensionImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
