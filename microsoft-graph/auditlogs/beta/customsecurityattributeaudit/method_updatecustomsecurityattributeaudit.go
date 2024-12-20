package customsecurityattributeaudit

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCustomSecurityAttributeAuditOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCustomSecurityAttributeAuditOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateCustomSecurityAttributeAuditOperationOptions() UpdateCustomSecurityAttributeAuditOperationOptions {
	return UpdateCustomSecurityAttributeAuditOperationOptions{}
}

func (o UpdateCustomSecurityAttributeAuditOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCustomSecurityAttributeAuditOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCustomSecurityAttributeAuditOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCustomSecurityAttributeAudit - Update the navigation property customSecurityAttributeAudits in auditLogs
func (c CustomSecurityAttributeAuditClient) UpdateCustomSecurityAttributeAudit(ctx context.Context, id beta.AuditLogCustomSecurityAttributeAuditId, input beta.CustomSecurityAttributeAudit, options UpdateCustomSecurityAttributeAuditOperationOptions) (result UpdateCustomSecurityAttributeAuditOperationResponse, err error) {
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
