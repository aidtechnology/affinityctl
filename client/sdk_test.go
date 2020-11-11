package client

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	tdd "github.com/stretchr/testify/assert"
)

var sampleDID = "did:elem:EiDotMNs0iqUqmrWZ7zq0sSyhl1WLRCkr-BTa6RQ59887Q"

// var sampleAPIKey = "95273b4f-dd15-4eea-8e0e-0de5eb2d80e0"
// var sampleAPIKeyUser = "aidtechnology"
var sampleAPIKeyHash = "e838f9d866ebc239500920054acf5d34688e3af956726ade9a32ef5343020bc9"

func TestSDK(t *testing.T) {
	assert := tdd.New(t)
	opts := DefaultOptions()
	opts.Key = sampleAPIKeyHash
	cl, err := New(nil)
	assert.Nil(err, "new client")

	t.Run("DID", func(t *testing.T) {
		t.Run("GetMaterial", func(t *testing.T) {
			pin, err := cl.DID.GetMaterial()
			assert.Nil(err, "get material")
			t.Log(pin)
		})

		t.Run("Create", func(t *testing.T) {
			did, err := cl.DID.Create("S134RV", "sample@aid.technolgy")
			assert.Nil(err, "create")
			t.Log(did)
			sampleDID = did
		})

		t.Run("Resolve", func(t *testing.T) {
			doc, err := cl.DID.Resolve("did:elem:EiCIbjMlHw5aGdGCMJ21OqDyxAvOcE2r2xrunazFE037dw")
			assert.Nil(err, "create")
			t.Logf("%s", doc)
		})

		t.Run("Authenticate", func(t *testing.T) {
			isAuth, vcs, err := cl.DID.Authenticate(sampleDID, "S134RV")
			assert.Nil(err, "authenticate")
			assert.True(isAuth, "authenticate")
			t.Logf("%s", vcs)
		})

		t.Run("ResetMaterial", func(t *testing.T) {
			err := cl.DID.ResetMaterial(sampleDID)
			assert.Nil(err, "reset material")
		})
	})

	t.Run("VC", func(t *testing.T) {
		// Sample issuer
		issuer := &Issuer{
			DID: sampleDID, // "did:elem:EiDDxkcB1XV4w_NvZrw3E2E6-5YmQGNE3_bddxP88QavWg",
			PIN: "S134RV",  // "OX3POH",
		}

		t.Run("Issue", func(t *testing.T) {
			// Get sample credential payload
			credentialPayload, _ := ioutil.ReadFile("testdata/payload.json")
			payload := make(map[string]interface{})
			_ = json.Unmarshal(credentialPayload, &payload)

			vc, err := cl.VC.Issue(issuer, sampleDID, payload)
			assert.Nil(err, "issue")
			t.Logf("%s", vc)
		})

		t.Run("Verify", func(t *testing.T) {
			// Load sample VC
			vc := make(map[string]interface{})
			data, _ := ioutil.ReadFile("testdata/vc.json")
			_ = json.Unmarshal(data, &vc)

			res, err := cl.VC.Verify(vc)
			assert.Nil(err, "verify")
			assert.True(res, "invalid credential")
		})

		t.Run("Store", func(t *testing.T) {
			// Load sample VC
			vc := make(map[string]interface{})
			data, _ := ioutil.ReadFile("testdata/vc.json")
			_ = json.Unmarshal(data, &vc)

			err := cl.VC.Store(sampleDID, "S134RV", vc)
			assert.Nil(err, "store failed")
		})
	})
}
