package lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubjectmailboxsetting

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

type UpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingOperationOptions() UpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingOperationOptions {
	return UpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingOperationOptions{}
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSetting - Update property
// mailboxSettings value.
func (c LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingClient) UpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSetting(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportIdTaskProcessingResultId, input beta.MailboxSettings, options UpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingOperationOptions) (result UpdateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/subject/mailboxSettings", id.ID()),
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

	return
}
