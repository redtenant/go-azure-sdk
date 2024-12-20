package oauth2permissiongrant

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

type DeleteOAuth2PermissionGrantOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteOAuth2PermissionGrantOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteOAuth2PermissionGrantOperationOptions() DeleteOAuth2PermissionGrantOperationOptions {
	return DeleteOAuth2PermissionGrantOperationOptions{}
}

func (o DeleteOAuth2PermissionGrantOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteOAuth2PermissionGrantOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteOAuth2PermissionGrantOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteOAuth2PermissionGrant - Delete oAuth2PermissionGrant (a delegated permission grant). Delete an
// oAuth2PermissionGrant, representing a delegated permission grant. When a delegated permission grant is deleted, the
// access it granted is revoked. Existing access tokens will continue to be valid for their lifetime, but new access
// tokens will not be granted for the delegated permissions identified in the deleted oAuth2PermissionGrant.
func (c OAuth2PermissionGrantClient) DeleteOAuth2PermissionGrant(ctx context.Context, id beta.OAuth2PermissionGrantId, options DeleteOAuth2PermissionGrantOperationOptions) (result DeleteOAuth2PermissionGrantOperationResponse, err error) {
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
