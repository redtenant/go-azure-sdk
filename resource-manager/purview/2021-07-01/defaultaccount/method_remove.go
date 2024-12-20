package defaultaccount

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveOperationOptions struct {
	Scope         *string
	ScopeTenantId *string
	ScopeType     *ScopeType
}

func DefaultRemoveOperationOptions() RemoveOperationOptions {
	return RemoveOperationOptions{}
}

func (o RemoveOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RemoveOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Scope != nil {
		out.Append("scope", fmt.Sprintf("%v", *o.Scope))
	}
	if o.ScopeTenantId != nil {
		out.Append("scopeTenantId", fmt.Sprintf("%v", *o.ScopeTenantId))
	}
	if o.ScopeType != nil {
		out.Append("scopeType", fmt.Sprintf("%v", *o.ScopeType))
	}
	return &out
}

// Remove ...
func (c DefaultAccountClient) Remove(ctx context.Context, options RemoveOperationOptions) (result RemoveOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/providers/Microsoft.Purview/removeDefaultAccount",
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
