package entitlementmanagementaccesspackage

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"

type AccessPackageOperationPredicate struct {
}

func (p AccessPackageOperationPredicate) Matches(input stable.AccessPackage) bool {

	return true
}

type AccessPackageAssignmentRequestRequirementsOperationPredicate struct {
}

func (p AccessPackageAssignmentRequestRequirementsOperationPredicate) Matches(input stable.AccessPackageAssignmentRequestRequirements) bool {

	return true
}
