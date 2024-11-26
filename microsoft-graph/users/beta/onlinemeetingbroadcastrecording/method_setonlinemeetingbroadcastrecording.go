package onlinemeetingbroadcastrecording

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

type SetOnlineMeetingBroadcastRecordingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetOnlineMeetingBroadcastRecordingOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetOnlineMeetingBroadcastRecordingOperationOptions() SetOnlineMeetingBroadcastRecordingOperationOptions {
	return SetOnlineMeetingBroadcastRecordingOperationOptions{}
}

func (o SetOnlineMeetingBroadcastRecordingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetOnlineMeetingBroadcastRecordingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetOnlineMeetingBroadcastRecordingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetOnlineMeetingBroadcastRecording - Update broadcastRecording for the navigation property onlineMeetings in users
func (c OnlineMeetingBroadcastRecordingClient) SetOnlineMeetingBroadcastRecording(ctx context.Context, id beta.UserIdOnlineMeetingId, input []byte, options SetOnlineMeetingBroadcastRecordingOperationOptions) (result SetOnlineMeetingBroadcastRecordingOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/broadcastRecording", id.ID()),
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
