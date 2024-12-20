package virtualendpointprovisioningpolicy

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

type CreateVirtualEndpointProvisioningPolicyApplyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateVirtualEndpointProvisioningPolicyApplyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateVirtualEndpointProvisioningPolicyApplyOperationOptions() CreateVirtualEndpointProvisioningPolicyApplyOperationOptions {
	return CreateVirtualEndpointProvisioningPolicyApplyOperationOptions{}
}

func (o CreateVirtualEndpointProvisioningPolicyApplyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointProvisioningPolicyApplyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointProvisioningPolicyApplyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointProvisioningPolicyApply - Invoke action apply. Apply the current provisioning policy
// configuration to all Cloud PC devices under a specified policy. Currently, the region is the only policy setting that
// you can apply.
func (c VirtualEndpointProvisioningPolicyClient) CreateVirtualEndpointProvisioningPolicyApply(ctx context.Context, id beta.DeviceManagementVirtualEndpointProvisioningPolicyId, input CreateVirtualEndpointProvisioningPolicyApplyRequest, options CreateVirtualEndpointProvisioningPolicyApplyOperationOptions) (result CreateVirtualEndpointProvisioningPolicyApplyOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/apply", id.ID()),
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
