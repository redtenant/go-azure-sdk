package workspaces

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/client/pollers"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceAadAdminsDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// WorkspaceAadAdminsDelete ...
func (c WorkspacesClient) WorkspaceAadAdminsDelete(ctx context.Context, id WorkspaceId) (result WorkspaceAadAdminsDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod: http.MethodDelete,
		Path:       fmt.Sprintf("%s/administrators/activeDirectory", id.ID()),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// WorkspaceAadAdminsDeleteThenPoll performs WorkspaceAadAdminsDelete then polls until it's completed
func (c WorkspacesClient) WorkspaceAadAdminsDeleteThenPoll(ctx context.Context, id WorkspaceId) error {
	result, err := c.WorkspaceAadAdminsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing WorkspaceAadAdminsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after WorkspaceAadAdminsDelete: %+v", err)
	}

	return nil
}
