package client

import (
	"encoding/json"
)

// Issuer of verifiable credentials.
type Issuer struct {
	DID string `json:"issuerDid"`
	PIN string `json:"issuerPin"`
}

type vcService struct {
	sdk *SDK
}

// Issue a new verifiable credential.
func (m *vcService) Issue(iss *Issuer, did string, payload interface{}) ([]byte, error) {
	req := map[string]interface{}{
		"did":       did,
		"issuerDid": iss.DID,
		"issuerPin": iss.PIN,
		"payload":   payload,
	}
	pl := map[string]interface{}{
		"credential": nil,
	}
	err := m.sdk.request("POST", "/vc/issue", req, pl)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(pl["credential"], "", "  ")
}

// Verify the validity of an existing credential.
func (m *vcService) Verify(credential interface{}) (bool, error) {
	req := map[string]interface{}{
		"credentials": []interface{}{credential},
		"options":     map[string]string{},
	}
	pl := map[string]interface{}{
		"isOk": false,
	}
	err := m.sdk.request("PUT", "/vc/verify", req, pl)
	if err != nil {
		return false, err
	}
	return pl["isOk"].(bool), nil
}

// Store a new verifiable credential on the user's wallet.
func (m *vcService) Store(did string, material string, vc interface{}) error {
	req := map[string]interface{}{
		"did":      did,
		"material": material,
		"vc":       vc,
	}
	pl := map[string]interface{}{
		"storedVcReference": false,
	}
	return m.sdk.request("POST", "/vc/store", req, pl)
}
