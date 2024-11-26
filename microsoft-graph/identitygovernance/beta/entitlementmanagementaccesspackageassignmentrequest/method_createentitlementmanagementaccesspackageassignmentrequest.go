package entitlementmanagementaccesspackageassignmentrequest

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementAccessPackageAssignmentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageAssignmentRequest
}

type CreateEntitlementManagementAccessPackageAssignmentRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageAssignmentRequestOperationOptions() CreateEntitlementManagementAccessPackageAssignmentRequestOperationOptions {
	return CreateEntitlementManagementAccessPackageAssignmentRequestOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageAssignmentRequest - Create accessPackageAssignmentRequest. In Microsoft Entra
// Entitlement Management, create a new accessPackageAssignmentRequest object. This operation is used to assign a user
// to an access package, or to remove an access package assignment.
func (c EntitlementManagementAccessPackageAssignmentRequestClient) CreateEntitlementManagementAccessPackageAssignmentRequest(ctx context.Context, input beta.AccessPackageAssignmentRequest, options CreateEntitlementManagementAccessPackageAssignmentRequestOperationOptions) (result CreateEntitlementManagementAccessPackageAssignmentRequestOperationResponse, err error) {
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
		Path:          "/identityGovernance/entitlementManagement/accessPackageAssignmentRequests",
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

	var model beta.AccessPackageAssignmentRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
