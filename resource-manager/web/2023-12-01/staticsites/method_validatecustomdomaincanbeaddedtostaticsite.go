package staticsites

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

type ValidateCustomDomainCanBeAddedToStaticSiteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ValidateCustomDomainCanBeAddedToStaticSite ...
func (c StaticSitesClient) ValidateCustomDomainCanBeAddedToStaticSite(ctx context.Context, id CustomDomainId, input StaticSiteCustomDomainRequestPropertiesARMResource) (result ValidateCustomDomainCanBeAddedToStaticSiteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/validate", id.ID()),
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

// ValidateCustomDomainCanBeAddedToStaticSiteThenPoll performs ValidateCustomDomainCanBeAddedToStaticSite then polls until it's completed
func (c StaticSitesClient) ValidateCustomDomainCanBeAddedToStaticSiteThenPoll(ctx context.Context, id CustomDomainId, input StaticSiteCustomDomainRequestPropertiesARMResource) error {
	result, err := c.ValidateCustomDomainCanBeAddedToStaticSite(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ValidateCustomDomainCanBeAddedToStaticSite: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ValidateCustomDomainCanBeAddedToStaticSite: %+v", err)
	}

	return nil
}
