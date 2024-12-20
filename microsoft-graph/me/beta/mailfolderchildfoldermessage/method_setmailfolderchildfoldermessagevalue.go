package mailfolderchildfoldermessage

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

type SetMailFolderChildFolderMessageValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetMailFolderChildFolderMessageValueOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetMailFolderChildFolderMessageValueOperationOptions() SetMailFolderChildFolderMessageValueOperationOptions {
	return SetMailFolderChildFolderMessageValueOperationOptions{}
}

func (o SetMailFolderChildFolderMessageValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetMailFolderChildFolderMessageValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetMailFolderChildFolderMessageValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetMailFolderChildFolderMessageValue - Update media content for the navigation property messages in me. The unique
// identifier for an entity. Read-only.
func (c MailFolderChildFolderMessageClient) SetMailFolderChildFolderMessageValue(ctx context.Context, id beta.MeMailFolderIdChildFolderIdMessageId, input []byte, options SetMailFolderChildFolderMessageValueOperationOptions) (result SetMailFolderChildFolderMessageValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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
