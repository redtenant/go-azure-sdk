package amlfilesystems

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRequiredAmlFSSubnetsSizeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RequiredAmlFilesystemSubnetsSize
}

// GetRequiredAmlFSSubnetsSize ...
func (c AmlFilesystemsClient) GetRequiredAmlFSSubnetsSize(ctx context.Context, id commonids.SubscriptionId, input RequiredAmlFilesystemSubnetsSizeInfo) (result GetRequiredAmlFSSubnetsSizeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.StorageCache/getRequiredAmlFSSubnetsSize", id.ID()),
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

	var model RequiredAmlFilesystemSubnetsSize
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
