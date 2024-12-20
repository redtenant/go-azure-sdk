package lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresultsubjectmailboxsetting

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

type GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.MailboxSettings
}

type GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSettingOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSettingOperationOptions() GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSettingOperationOptions {
	return GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSettingOperationOptions{}
}

func (o GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSetting - Get mailboxSettings
// property value. Settings for the primary mailbox of the signed-in user. You can get or update settings for sending
// automatic replies to incoming messages, locale, and time zone. Returned only on $select.
func (c LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectMailboxSettingClient) GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSetting(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskIdTaskProcessingResultId, options GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSettingOperationOptions) (result GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultSubjectMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/subject/mailboxSettings", id.ID()),
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

	var model stable.MailboxSettings
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
