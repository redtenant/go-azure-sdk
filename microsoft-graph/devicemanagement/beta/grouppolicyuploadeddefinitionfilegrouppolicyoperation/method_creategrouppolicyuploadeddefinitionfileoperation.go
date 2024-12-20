package grouppolicyuploadeddefinitionfilegrouppolicyoperation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateGroupPolicyUploadedDefinitionFileOperationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.GroupPolicyOperation
}

type CreateGroupPolicyUploadedDefinitionFileOperationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateGroupPolicyUploadedDefinitionFileOperationOperationOptions() CreateGroupPolicyUploadedDefinitionFileOperationOperationOptions {
	return CreateGroupPolicyUploadedDefinitionFileOperationOperationOptions{}
}

func (o CreateGroupPolicyUploadedDefinitionFileOperationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateGroupPolicyUploadedDefinitionFileOperationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateGroupPolicyUploadedDefinitionFileOperationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateGroupPolicyUploadedDefinitionFileOperation - Create new navigation property to groupPolicyOperations for
// deviceManagement
func (c GroupPolicyUploadedDefinitionFileGroupPolicyOperationClient) CreateGroupPolicyUploadedDefinitionFileOperation(ctx context.Context, id beta.DeviceManagementGroupPolicyUploadedDefinitionFileId, input beta.GroupPolicyOperation, options CreateGroupPolicyUploadedDefinitionFileOperationOperationOptions) (result CreateGroupPolicyUploadedDefinitionFileOperationOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/groupPolicyOperations", id.ID()),
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

	var model beta.GroupPolicyOperation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
