package mailfolderchildfoldermessageextension

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateMailFolderChildFolderMessageExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateMailFolderChildFolderMessageExtensionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateMailFolderChildFolderMessageExtensionOperationOptions() UpdateMailFolderChildFolderMessageExtensionOperationOptions {
	return UpdateMailFolderChildFolderMessageExtensionOperationOptions{}
}

func (o UpdateMailFolderChildFolderMessageExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateMailFolderChildFolderMessageExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateMailFolderChildFolderMessageExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateMailFolderChildFolderMessageExtension - Update the navigation property extensions in me
func (c MailFolderChildFolderMessageExtensionClient) UpdateMailFolderChildFolderMessageExtension(ctx context.Context, id stable.MeMailFolderIdChildFolderIdMessageIdExtensionId, input stable.Extension, options UpdateMailFolderChildFolderMessageExtensionOperationOptions) (result UpdateMailFolderChildFolderMessageExtensionOperationResponse, err error) {
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
