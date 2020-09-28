package client

import (
	"encoding/json"
	"errors"
)

type didService struct {
	sdk *SDK
}

// GetMaterial returns a newly generated PIN value.
func (m *didService) GetMaterial() (string, error) {
	pl := map[string]interface{}{
		"pin": "",
	}
	err := m.sdk.request("GET", "/did/material", nil, pl)
	if err != nil {
		return "", err
	}
	return pl["pin"].(string), nil
}

// Create a new DID with given material.
func (m *didService) Create(material, email string) (string, error) {
	req := map[string]string{
		"material":           material,
		"branchManagerEmail": email,
	}
	pl := map[string]interface{}{
		"did": "",
	}
	err := m.sdk.request("POST", "/did/material", req, pl)
	if err != nil {
		return "", err
	}
	return pl["did"].(string), nil
}

// Resolve returns the JSON encoded document associated with the
// provided DID, if available.
func (m *didService) Resolve(did string) ([]byte, error) {
	req := map[string]string{
		"did": did,
	}
	pl := map[string]interface{}{
		"did":      "",
		"document": nil,
	}
	err := m.sdk.request("PUT", "/did/resolve", req, pl)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(pl["document"], "", "  ")
}

// Authenticate and existing user with DID and material. The VCs from the wallet
// are returned, if any.
func (m *didService) Authenticate(did, material string) (bool, []byte, error) {
	req := map[string]string{
		"did":      did,
		"material": material,
	}
	pl := map[string]interface{}{
		"isAuthenticated": false,
		"vcs":             []interface{}{},
	}
	err := m.sdk.request("PUT", "/authentications", req, pl)
	if err != nil {
		return false, nil, err
	}
	isAuth, _ := pl["isAuthenticated"].(bool)
	vcs, _ := json.MarshalIndent(pl["vcs"], "", "  ")
	return isAuth, vcs, nil
}

// ResetMaterial allows a user to remove the current PIN associated with
// a given identifier.
func (m *didService) ResetMaterial(did, material string) error {
	return errors.New("not yet implemented")
	// // Get confirmation code
	// r1 := map[string]string{
	// 	"did": did,
	// }
	// p1 := map[string]interface{}{
	// 	"confirmation": "",
	// }
	// err := m.sdk.request("POST", "/material/userMaterialReset", r1, p1)
	// if err != nil {
	// 	return err
	// }
	//
	// // Send confirmation code
	// r2 := map[string]string{
	// 	"did": did,
	// 	"material": material,
	// 	"confirmation": p1["confirmation"].(string),
	// }
	// p2 := make(map[string]interface{})
	// return m.sdk.request("POST", "/material/userMaterialResetConfirm", r2, p2)
}
