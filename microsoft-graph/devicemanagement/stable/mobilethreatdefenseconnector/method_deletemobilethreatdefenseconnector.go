package mobilethreatdefenseconnector

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

type DeleteMobileThreatDefenseConnectorOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteMobileThreatDefenseConnectorOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteMobileThreatDefenseConnectorOperationOptions() DeleteMobileThreatDefenseConnectorOperationOptions {
	return DeleteMobileThreatDefenseConnectorOperationOptions{}
}

func (o DeleteMobileThreatDefenseConnectorOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteMobileThreatDefenseConnectorOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteMobileThreatDefenseConnectorOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteMobileThreatDefenseConnector - Delete mobileThreatDefenseConnector. Deletes a mobileThreatDefenseConnector.
func (c MobileThreatDefenseConnectorClient) DeleteMobileThreatDefenseConnector(ctx context.Context, id stable.DeviceManagementMobileThreatDefenseConnectorId, options DeleteMobileThreatDefenseConnectorOperationOptions) (result DeleteMobileThreatDefenseConnectorOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
