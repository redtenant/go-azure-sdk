package onlinemeetingtranscriptcontent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetOnlineMeetingTranscriptContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetOnlineMeetingTranscriptContentOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetOnlineMeetingTranscriptContentOperationOptions() SetOnlineMeetingTranscriptContentOperationOptions {
	return SetOnlineMeetingTranscriptContentOperationOptions{}
}

func (o SetOnlineMeetingTranscriptContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetOnlineMeetingTranscriptContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetOnlineMeetingTranscriptContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetOnlineMeetingTranscriptContent - Update content for the navigation property transcripts in users. The content of
// the transcript. Read-only.
func (c OnlineMeetingTranscriptContentClient) SetOnlineMeetingTranscriptContent(ctx context.Context, id stable.UserIdOnlineMeetingIdTranscriptId, input []byte, options SetOnlineMeetingTranscriptContentOperationOptions) (result SetOnlineMeetingTranscriptContentOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/content", id.ID()),
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
