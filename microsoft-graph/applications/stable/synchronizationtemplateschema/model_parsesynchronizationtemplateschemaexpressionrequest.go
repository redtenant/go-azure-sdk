package synchronizationtemplateschema

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParseSynchronizationTemplateSchemaExpressionRequest struct {
	Expression                nullable.Type[string]         `json:"expression,omitempty"`
	TargetAttributeDefinition *stable.AttributeDefinition   `json:"targetAttributeDefinition,omitempty"`
	TestInputObject           *stable.ExpressionInputObject `json:"testInputObject,omitempty"`
}
