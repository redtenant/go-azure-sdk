package threatintelligence

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndicatorCreateIndicatorOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        ThreatIntelligenceInformation
}

// IndicatorCreateIndicator ...
func (c ThreatIntelligenceClient) IndicatorCreateIndicator(ctx context.Context, id WorkspaceId, input ThreatIntelligenceIndicatorModel) (result IndicatorCreateIndicatorOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/threatIntelligence/main/createIndicator", id.ID()),
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
