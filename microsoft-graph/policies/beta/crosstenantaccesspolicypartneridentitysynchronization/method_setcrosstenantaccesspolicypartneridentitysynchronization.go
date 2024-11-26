package crosstenantaccesspolicypartneridentitysynchronization

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

type SetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions() SetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions {
	return SetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions{}
}

func (o SetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetCrossTenantAccessPolicyPartnerIdentitySynchronization - Update crossTenantIdentitySyncPolicyPartner. Update the
// user synchronization policy of a partner-specific configuration.
func (c CrossTenantAccessPolicyPartnerIdentitySynchronizationClient) SetCrossTenantAccessPolicyPartnerIdentitySynchronization(ctx context.Context, id beta.PolicyCrossTenantAccessPolicyPartnerId, input beta.CrossTenantIdentitySyncPolicyPartner, options SetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions) (result SetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/identitySynchronization", id.ID()),
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
