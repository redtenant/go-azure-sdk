package teamphoto

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

type RemoveTeamPhotoValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveTeamPhotoValueOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveTeamPhotoValueOperationOptions() RemoveTeamPhotoValueOperationOptions {
	return RemoveTeamPhotoValueOperationOptions{}
}

func (o RemoveTeamPhotoValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveTeamPhotoValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveTeamPhotoValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveTeamPhotoValue - Delete media content for the navigation property photo in groups. The profile photo for the
// team.
func (c TeamPhotoClient) RemoveTeamPhotoValue(ctx context.Context, id stable.GroupId, options RemoveTeamPhotoValueOperationOptions) (result RemoveTeamPhotoValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/photo/$value", id.ID()),
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
