package virtualendpointcrosscloudgovernmentorganizationmapping

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions() DeleteVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions {
	return DeleteVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions{}
}

func (o DeleteVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteVirtualEndpointCrossCloudGovernmentOrganizationMapping - Delete navigation property
// crossCloudGovernmentOrganizationMapping for deviceManagement
func (c VirtualEndpointCrossCloudGovernmentOrganizationMappingClient) DeleteVirtualEndpointCrossCloudGovernmentOrganizationMapping(ctx context.Context, options DeleteVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationOptions) (result DeleteVirtualEndpointCrossCloudGovernmentOrganizationMappingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/crossCloudGovernmentOrganizationMapping",
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
