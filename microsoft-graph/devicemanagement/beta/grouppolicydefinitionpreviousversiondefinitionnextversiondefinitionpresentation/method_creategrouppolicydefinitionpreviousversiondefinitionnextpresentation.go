package grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.GroupPolicyPresentation
}

type CreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions() CreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions {
	return CreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions{}
}

func (o CreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentation - Create new navigation property to
// presentations for deviceManagement
func (c GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient) CreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentation(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, input beta.GroupPolicyPresentation, options CreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions) (result CreateGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/previousVersionDefinition/nextVersionDefinition/presentations", id.ID()),
		RetryFunc:     options.RetryFunc,
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
	model, err := beta.UnmarshalGroupPolicyPresentationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
