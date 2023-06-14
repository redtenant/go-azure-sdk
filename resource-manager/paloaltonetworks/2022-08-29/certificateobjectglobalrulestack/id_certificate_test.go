package certificateobjectglobalrulestack

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CertificateId{}

func TestNewCertificateID(t *testing.T) {
	id := NewCertificateID("globalRuleStackValue", "certificateValue")

	if id.GlobalRuleStackName != "globalRuleStackValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GlobalRuleStackName'", id.GlobalRuleStackName, "globalRuleStackValue")
	}

	if id.CertificateName != "certificateValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CertificateName'", id.CertificateName, "certificateValue")
	}
}

func TestFormatCertificateID(t *testing.T) {
	actual := NewCertificateID("globalRuleStackValue", "certificateValue").ID()
	expected := "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/certificates/certificateValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCertificateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CertificateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/certificates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/certificates/certificateValue",
			Expected: &CertificateId{
				GlobalRuleStackName: "globalRuleStackValue",
				CertificateName:     "certificateValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/certificates/certificateValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCertificateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GlobalRuleStackName != v.Expected.GlobalRuleStackName {
			t.Fatalf("Expected %q but got %q for GlobalRuleStackName", v.Expected.GlobalRuleStackName, actual.GlobalRuleStackName)
		}

		if actual.CertificateName != v.Expected.CertificateName {
			t.Fatalf("Expected %q but got %q for CertificateName", v.Expected.CertificateName, actual.CertificateName)
		}

	}
}

func TestParseCertificateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CertificateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/certificates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/cErTiFiCaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/certificates/certificateValue",
			Expected: &CertificateId{
				GlobalRuleStackName: "globalRuleStackValue",
				CertificateName:     "certificateValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/certificates/certificateValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/cErTiFiCaTeS/cErTiFiCaTeVaLuE",
			Expected: &CertificateId{
				GlobalRuleStackName: "gLoBaLrUlEsTaCkVaLuE",
				CertificateName:     "cErTiFiCaTeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/cErTiFiCaTeS/cErTiFiCaTeVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCertificateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GlobalRuleStackName != v.Expected.GlobalRuleStackName {
			t.Fatalf("Expected %q but got %q for GlobalRuleStackName", v.Expected.GlobalRuleStackName, actual.GlobalRuleStackName)
		}

		if actual.CertificateName != v.Expected.CertificateName {
			t.Fatalf("Expected %q but got %q for CertificateName", v.Expected.CertificateName, actual.CertificateName)
		}

	}
}

func TestSegmentsForCertificateId(t *testing.T) {
	segments := CertificateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CertificateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
