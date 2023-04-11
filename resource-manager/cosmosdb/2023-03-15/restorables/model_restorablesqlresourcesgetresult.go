package restorables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableSqlResourcesGetResult struct {
	CollectionNames *[]string `json:"collectionNames,omitempty"`
	DatabaseName    *string   `json:"databaseName,omitempty"`
	Id              *string   `json:"id,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Type            *string   `json:"type,omitempty"`
}
