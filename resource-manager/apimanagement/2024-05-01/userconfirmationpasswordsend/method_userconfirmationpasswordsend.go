package userconfirmationpasswordsend

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserConfirmationPasswordSendOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UserConfirmationPasswordSendOperationOptions struct {
	AppType *AppType
}

func DefaultUserConfirmationPasswordSendOperationOptions() UserConfirmationPasswordSendOperationOptions {
	return UserConfirmationPasswordSendOperationOptions{}
}

func (o UserConfirmationPasswordSendOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UserConfirmationPasswordSendOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o UserConfirmationPasswordSendOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.AppType != nil {
		out.Append("appType", fmt.Sprintf("%v", *o.AppType))
	}
	return &out
}

// UserConfirmationPasswordSend ...
func (c UserConfirmationPasswordSendClient) UserConfirmationPasswordSend(ctx context.Context, id UserId, options UserConfirmationPasswordSendOperationOptions) (result UserConfirmationPasswordSendOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/confirmations/password/send", id.ID()),
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
