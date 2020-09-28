package client

import (
	"testing"

	tdd "github.com/stretchr/testify/assert"
)

func TestSDK(t *testing.T) {
	assert := tdd.New(t)
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
		})

		t.Run("Resolve", func(t *testing.T) {
			doc, err := cl.DID.Resolve("did:elem:EiCIbjMlHw5aGdGCMJ21OqDyxAvOcE2r2xrunazFE037dw")
			assert.Nil(err, "create")
			t.Logf("%s", doc)
		})

		t.Run("Authenticate", func(t *testing.T) {
			isAuth, vcs, err := cl.DID.Authenticate("did:elem:EiCIbjMlHw5aGdGCMJ21OqDyxAvOcE2r2xrunazFE037dw", "S134RV")
			assert.Nil(err, "authenticate")
			assert.True(isAuth, "authenticate")
			t.Logf("%s", vcs)
		})

		t.Run("ResetMaterial", func(t *testing.T) {
			err := cl.DID.ResetMaterial("did:elem:EiCIbjMlHw5aGdGCMJ21OqDyxAvOcE2r2xrunazFE037dw", "S134RV")
			assert.NotNil(err, "reset material")
		})
	})

	t.Run("VC", func(t *testing.T) {
		t.Run("Issue", func(t *testing.T) {
			sampleDID := "did:elem:EiCIbjMlHw5aGdGCMJ21OqDyxAvOcE2r2xrunazFE037dw"
			iss := &Issuer{
				DID: sampleDID,
				PIN: "S134RV",
			}
			payload := map[string]interface{}{
				"resourceType": "string",
			}
			vc, err := cl.VC.Issue(iss, sampleDID, payload)
			assert.Nil(err, "issue")
			t.Logf("%s", vc)
		})

		t.Run("Verify", func(t *testing.T) {
			cred := map[string]interface{}{
				"exitingVC": "sample-structure",
				"foo":       "bar",
				"baz":       true,
				"numeric":   10,
			}
			res, err := cl.VC.Verify(cred)
			assert.Nil(err, "verify")
			assert.False(res, "invalid credential")
		})

		t.Run("Store", func(t *testing.T) {
			sampleDID := "did:elem:EiCIbjMlHw5aGdGCMJ21OqDyxAvOcE2r2xrunazFE037dw"
			cred := map[string]interface{}{
				"exitingVC": "sample-structure",
				"foo":       "bar",
				"baz":       true,
				"numeric":   10,
			}
			_, err := cl.VC.Store(sampleDID, "S134RV", cred)
			assert.Nil(err, "store failed")
		})
	})
}
