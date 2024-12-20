package checkdataconnectorrequirements

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorsCheckRequirementsPostOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DataConnectorRequirementsState
}

// DataConnectorsCheckRequirementsPost ...
func (c CheckDataConnectorRequirementsClient) DataConnectorsCheckRequirementsPost(ctx context.Context, id WorkspaceId, input DataConnectorsCheckRequirements) (result DataConnectorsCheckRequirementsPostOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/dataConnectorsCheckRequirements", id.ID()),
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

	var model DataConnectorRequirementsState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
