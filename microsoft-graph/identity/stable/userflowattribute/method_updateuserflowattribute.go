package userflowattribute

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserFlowAttributeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserFlowAttributeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateUserFlowAttributeOperationOptions() UpdateUserFlowAttributeOperationOptions {
	return UpdateUserFlowAttributeOperationOptions{}
}

func (o UpdateUserFlowAttributeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserFlowAttributeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserFlowAttributeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserFlowAttribute - Update identityUserFlowAttribute. Update the properties of a custom
// identityUserFlowAttribute object.
func (c UserFlowAttributeClient) UpdateUserFlowAttribute(ctx context.Context, id stable.IdentityUserFlowAttributeId, input stable.IdentityUserFlowAttribute, options UpdateUserFlowAttributeOperationOptions) (result UpdateUserFlowAttributeOperationResponse, err error) {
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
