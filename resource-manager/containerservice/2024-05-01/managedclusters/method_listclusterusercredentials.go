package managedclusters

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

type ListClusterUserCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CredentialResults
}

type ListClusterUserCredentialsOperationOptions struct {
	Format     *Format
	ServerFqdn *string
}

func DefaultListClusterUserCredentialsOperationOptions() ListClusterUserCredentialsOperationOptions {
	return ListClusterUserCredentialsOperationOptions{}
}

func (o ListClusterUserCredentialsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListClusterUserCredentialsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListClusterUserCredentialsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Format != nil {
		out.Append("format", fmt.Sprintf("%v", *o.Format))
	}
	if o.ServerFqdn != nil {
		out.Append("server-fqdn", fmt.Sprintf("%v", *o.ServerFqdn))
	}
	return &out
}

// ListClusterUserCredentials ...
func (c ManagedClustersClient) ListClusterUserCredentials(ctx context.Context, id commonids.KubernetesClusterId, options ListClusterUserCredentialsOperationOptions) (result ListClusterUserCredentialsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/listClusterUserCredential", id.ID()),
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

	var model CredentialResults
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
