package serviceprincipalcreationpolicyinclude

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

type DeleteServicePrincipalCreationPolicyIncludeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteServicePrincipalCreationPolicyIncludeOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteServicePrincipalCreationPolicyIncludeOperationOptions() DeleteServicePrincipalCreationPolicyIncludeOperationOptions {
	return DeleteServicePrincipalCreationPolicyIncludeOperationOptions{}
}

func (o DeleteServicePrincipalCreationPolicyIncludeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteServicePrincipalCreationPolicyIncludeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteServicePrincipalCreationPolicyIncludeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteServicePrincipalCreationPolicyInclude - Delete navigation property includes for policies
func (c ServicePrincipalCreationPolicyIncludeClient) DeleteServicePrincipalCreationPolicyInclude(ctx context.Context, id beta.PolicyServicePrincipalCreationPolicyIdIncludeId, options DeleteServicePrincipalCreationPolicyIncludeOperationOptions) (result DeleteServicePrincipalCreationPolicyIncludeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
