package assetendpointprofiles

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

type CreateOrReplaceOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *AssetEndpointProfile
}

// CreateOrReplace ...
func (c AssetEndpointProfilesClient) CreateOrReplace(ctx context.Context, id AssetEndpointProfileId, input AssetEndpointProfile) (result CreateOrReplaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// CreateOrReplaceThenPoll performs CreateOrReplace then polls until it's completed
func (c AssetEndpointProfilesClient) CreateOrReplaceThenPoll(ctx context.Context, id AssetEndpointProfileId, input AssetEndpointProfile) error {
	result, err := c.CreateOrReplace(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrReplace: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after CreateOrReplace: %+v", err)
	}

	return nil
}
