package joinedteamprimarychannelfilesfolder

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

type GetJoinedTeamPrimaryChannelFilesFolderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DriveItem
}

type GetJoinedTeamPrimaryChannelFilesFolderOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetJoinedTeamPrimaryChannelFilesFolderOperationOptions() GetJoinedTeamPrimaryChannelFilesFolderOperationOptions {
	return GetJoinedTeamPrimaryChannelFilesFolderOperationOptions{}
}

func (o GetJoinedTeamPrimaryChannelFilesFolderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetJoinedTeamPrimaryChannelFilesFolderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetJoinedTeamPrimaryChannelFilesFolderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetJoinedTeamPrimaryChannelFilesFolder - Get filesFolder from me. Metadata for the location where the channel's files
// are stored.
func (c JoinedTeamPrimaryChannelFilesFolderClient) GetJoinedTeamPrimaryChannelFilesFolder(ctx context.Context, id stable.MeJoinedTeamId, options GetJoinedTeamPrimaryChannelFilesFolderOperationOptions) (result GetJoinedTeamPrimaryChannelFilesFolderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/primaryChannel/filesFolder", id.ID()),
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

	var model stable.DriveItem
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
