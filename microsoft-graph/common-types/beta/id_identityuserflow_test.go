package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityUserFlowId{}

func TestNewIdentityUserFlowID(t *testing.T) {
	id := NewIdentityUserFlowID("identityUserFlowId")

	if id.IdentityUserFlowId != "identityUserFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityUserFlowId'", id.IdentityUserFlowId, "identityUserFlowId")
	}
}

func TestFormatIdentityUserFlowID(t *testing.T) {
	actual := NewIdentityUserFlowID("identityUserFlowId").ID()
	expected := "/identity/userFlows/identityUserFlowId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityUserFlowID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityUserFlowId
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
			Input: "/identity/userFlows",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/userFlows/identityUserFlowId",
			Expected: &IdentityUserFlowId{
				IdentityUserFlowId: "identityUserFlowId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/userFlows/identityUserFlowId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityUserFlowID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IdentityUserFlowId != v.Expected.IdentityUserFlowId {
			t.Fatalf("Expected %q but got %q for IdentityUserFlowId", v.Expected.IdentityUserFlowId, actual.IdentityUserFlowId)
		}

	}
}

func TestParseIdentityUserFlowIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityUserFlowId
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
			Input: "/identity/userFlows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/uSeRfLoWs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/userFlows/identityUserFlowId",
			Expected: &IdentityUserFlowId{
				IdentityUserFlowId: "identityUserFlowId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/userFlows/identityUserFlowId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/uSeRfLoWs/iDeNtItYuSeRfLoWiD",
			Expected: &IdentityUserFlowId{
				IdentityUserFlowId: "iDeNtItYuSeRfLoWiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/uSeRfLoWs/iDeNtItYuSeRfLoWiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityUserFlowIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IdentityUserFlowId != v.Expected.IdentityUserFlowId {
			t.Fatalf("Expected %q but got %q for IdentityUserFlowId", v.Expected.IdentityUserFlowId, actual.IdentityUserFlowId)
		}

	}
}

func TestSegmentsForIdentityUserFlowId(t *testing.T) {
	segments := IdentityUserFlowId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityUserFlowId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
