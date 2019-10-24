package types

import (
	"encoding/base64"
	"reflect"
	"testing"

	"github.com/jcjones/ct-mapreduce/storage"
)

const (
	crlEmptyBase64  = `MIH2AgEBMA0GCSqGSIb3DQEBCwUAMIGRMQswCQYDVQQGEwJVUzELMAkGA1UECBMCQ0ExDzANBgNVBAcTBklydmluZTElMCMGA1UEChMcV2VzdGVybiBEaWdpdGFsIFRlY2hub2xvZ2llczE9MDsGA1UEAxM0V2VzdGVybiBEaWdpdGFsIFRlY2hub2xvZ2llcyBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eRcNMTkwOTAzMDMyMDAxWhcNMTkwOTA3MDMyMDAxWqAwMC4wHwYDVR0jBBgwFoAUWRAanffYNzT9rdULrGiuAvegvYMwCwYDVR0UBAQCAgOG`
	crlFilledBase64 = `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`
)

func Test_DecodeCRL(t *testing.T) {
	data, err := base64.StdEncoding.DecodeString(crlEmptyBase64)
	if err != nil {
		t.Error(err)
	}

	emptyList, err := DecodeRawTBSCertList(data)
	if err != nil {
		t.Error(err)
	}

	if len(emptyList.RevokedCertificates) > 0 {
		t.Error("Expected an empty list.")
	}

	data, err = base64.StdEncoding.DecodeString(crlFilledBase64)
	if err != nil {
		t.Error(err)
	}

	filledList, err := DecodeRawTBSCertList(data)
	if err != nil {
		t.Error(err)
	}

	if len(filledList.RevokedCertificates) != 220 {
		t.Errorf("Expected 220 entries, got %d.", len(filledList.RevokedCertificates))
	}

	expectedSerial := storage.NewSerialFromHex("0101ea518c68c0f00789e9cd92736c75")
	actualSerial := storage.NewSerialFromBytes(filledList.RevokedCertificates[0].SerialNumber.Bytes)
	if !reflect.DeepEqual(expectedSerial, actualSerial) {
		t.Errorf("Expected %s, but got %s", expectedSerial, actualSerial)
	}
}

func Test_SerialSet(t *testing.T) {
	testSerials := []storage.Serial{
		storage.NewSerialFromHex("BB"),
		storage.NewSerialFromHex("AA"),
	}

	set := NewSerialSet()
	isNew := set.Add(testSerials[0])
	if isNew == false {
		t.Error("Should have been new")
	}
	isNew = set.Add(testSerials[0])
	if isNew == true {
		t.Error("Should not have been new")
	}
	isNew = set.Add(testSerials[1])
	if isNew == false {
		t.Error("Should have been new")
	}

	actualSerials := set.List()

	if len(actualSerials) != len(testSerials) {
		t.Error("Length mismatch")
	}

	for _, i := range actualSerials {
		var seen bool
		for _, j := range testSerials {
			if j.ID() == i.ID() {
				seen = true
				break
			}
		}
		if !seen {
			t.Errorf("Didn't find %v", i)
		}
	}
}
