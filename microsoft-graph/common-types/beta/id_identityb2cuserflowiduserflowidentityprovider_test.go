package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2cUserFlowIdUserFlowIdentityProviderId{}

func TestNewIdentityB2cUserFlowIdUserFlowIdentityProviderID(t *testing.T) {
	id := NewIdentityB2cUserFlowIdUserFlowIdentityProviderID("b2cIdentityUserFlowId", "identityProviderBaseId")

	if id.B2cIdentityUserFlowId != "b2cIdentityUserFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'B2cIdentityUserFlowId'", id.B2cIdentityUserFlowId, "b2cIdentityUserFlowId")
	}

	if id.IdentityProviderBaseId != "identityProviderBaseId" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityProviderBaseId'", id.IdentityProviderBaseId, "identityProviderBaseId")
	}
}

func TestFormatIdentityB2cUserFlowIdUserFlowIdentityProviderID(t *testing.T) {
	actual := NewIdentityB2cUserFlowIdUserFlowIdentityProviderID("b2cIdentityUserFlowId", "identityProviderBaseId").ID()
	expected := "/identity/b2cUserFlows/b2cIdentityUserFlowId/userFlowIdentityProviders/identityProviderBaseId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2cUserFlowIdUserFlowIdentityProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2cUserFlowIdUserFlowIdentityProviderId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/userFlowIdentityProviders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/userFlowIdentityProviders/identityProviderBaseId",
			Expected: &IdentityB2cUserFlowIdUserFlowIdentityProviderId{
				B2cIdentityUserFlowId:  "b2cIdentityUserFlowId",
				IdentityProviderBaseId: "identityProviderBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/userFlowIdentityProviders/identityProviderBaseId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2cUserFlowIdUserFlowIdentityProviderID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.B2cIdentityUserFlowId != v.Expected.B2cIdentityUserFlowId {
			t.Fatalf("Expected %q but got %q for B2cIdentityUserFlowId", v.Expected.B2cIdentityUserFlowId, actual.B2cIdentityUserFlowId)
		}

		if actual.IdentityProviderBaseId != v.Expected.IdentityProviderBaseId {
			t.Fatalf("Expected %q but got %q for IdentityProviderBaseId", v.Expected.IdentityProviderBaseId, actual.IdentityProviderBaseId)
		}

	}
}

func TestParseIdentityB2cUserFlowIdUserFlowIdentityProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2cUserFlowIdUserFlowIdentityProviderId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/userFlowIdentityProviders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/uSeRfLoWiDeNtItYpRoViDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/userFlowIdentityProviders/identityProviderBaseId",
			Expected: &IdentityB2cUserFlowIdUserFlowIdentityProviderId{
				B2cIdentityUserFlowId:  "b2cIdentityUserFlowId",
				IdentityProviderBaseId: "identityProviderBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/userFlowIdentityProviders/identityProviderBaseId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/uSeRfLoWiDeNtItYpRoViDeRs/iDeNtItYpRoViDeRbAsEiD",
			Expected: &IdentityB2cUserFlowIdUserFlowIdentityProviderId{
				B2cIdentityUserFlowId:  "b2cIdEnTiTyUsErFlOwId",
				IdentityProviderBaseId: "iDeNtItYpRoViDeRbAsEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/uSeRfLoWiDeNtItYpRoViDeRs/iDeNtItYpRoViDeRbAsEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2cUserFlowIdUserFlowIdentityProviderIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.B2cIdentityUserFlowId != v.Expected.B2cIdentityUserFlowId {
			t.Fatalf("Expected %q but got %q for B2cIdentityUserFlowId", v.Expected.B2cIdentityUserFlowId, actual.B2cIdentityUserFlowId)
		}

		if actual.IdentityProviderBaseId != v.Expected.IdentityProviderBaseId {
			t.Fatalf("Expected %q but got %q for IdentityProviderBaseId", v.Expected.IdentityProviderBaseId, actual.IdentityProviderBaseId)
		}

	}
}

func TestSegmentsForIdentityB2cUserFlowIdUserFlowIdentityProviderId(t *testing.T) {
	segments := IdentityB2cUserFlowIdUserFlowIdentityProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2cUserFlowIdUserFlowIdentityProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
