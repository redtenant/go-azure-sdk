package sharedemaildomaininvitation

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSharedEmailDomainInvitationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSharedEmailDomainInvitationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSharedEmailDomainInvitationOperationOptions() UpdateSharedEmailDomainInvitationOperationOptions {
	return UpdateSharedEmailDomainInvitationOperationOptions{}
}

func (o UpdateSharedEmailDomainInvitationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSharedEmailDomainInvitationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSharedEmailDomainInvitationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSharedEmailDomainInvitation - Update the navigation property sharedEmailDomainInvitations in domains
func (c SharedEmailDomainInvitationClient) UpdateSharedEmailDomainInvitation(ctx context.Context, id beta.DomainIdSharedEmailDomainInvitationId, input beta.SharedEmailDomainInvitation, options UpdateSharedEmailDomainInvitationOperationOptions) (result UpdateSharedEmailDomainInvitationOperationResponse, err error) {
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
