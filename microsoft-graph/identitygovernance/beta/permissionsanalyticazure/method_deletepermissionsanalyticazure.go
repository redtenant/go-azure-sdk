package permissionsanalyticazure

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePermissionsAnalyticAzureOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeletePermissionsAnalyticAzureOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeletePermissionsAnalyticAzureOperationOptions() DeletePermissionsAnalyticAzureOperationOptions {
	return DeletePermissionsAnalyticAzureOperationOptions{}
}

func (o DeletePermissionsAnalyticAzureOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeletePermissionsAnalyticAzureOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeletePermissionsAnalyticAzureOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeletePermissionsAnalyticAzure - Delete navigation property azure for identityGovernance
func (c PermissionsAnalyticAzureClient) DeletePermissionsAnalyticAzure(ctx context.Context, options DeletePermissionsAnalyticAzureOperationOptions) (result DeletePermissionsAnalyticAzureOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/identityGovernance/permissionsAnalytics/azure",
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
