package threatintelligence

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndicatorGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        ThreatIntelligenceInformation
}

// IndicatorGet ...
func (c ThreatIntelligenceClient) IndicatorGet(ctx context.Context, id IndicatorId) (result IndicatorGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       id.ID(),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := UnmarshalThreatIntelligenceInformationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
