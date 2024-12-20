package provisionedclusterinstances

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/client/pollers"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteKubernetesVersionsOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// DeleteKubernetesVersions ...
func (c ProvisionedClusterInstancesClient) DeleteKubernetesVersions(ctx context.Context, id commonids.ScopeId) (result DeleteKubernetesVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
		},
		HttpMethod: http.MethodDelete,
		Path:       fmt.Sprintf("%s/providers/Microsoft.HybridContainerService/kubernetesVersions/default", id.ID()),
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

// DeleteKubernetesVersionsThenPoll performs DeleteKubernetesVersions then polls until it's completed
func (c ProvisionedClusterInstancesClient) DeleteKubernetesVersionsThenPoll(ctx context.Context, id commonids.ScopeId) error {
	result, err := c.DeleteKubernetesVersions(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteKubernetesVersions: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DeleteKubernetesVersions: %+v", err)
	}

	return nil
}
