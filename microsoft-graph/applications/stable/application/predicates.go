package application

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"

type ApplicationOperationPredicate struct {
}

func (p ApplicationOperationPredicate) Matches(input stable.Application) bool {

	return true
}

type DirectoryObjectOperationPredicate struct {
}

func (p DirectoryObjectOperationPredicate) Matches(input stable.DirectoryObject) bool {

	return true
}

type ExtensionPropertyOperationPredicate struct {
}

func (p ExtensionPropertyOperationPredicate) Matches(input stable.ExtensionProperty) bool {

	return true
}
