const axios = require('axios');

// API endpoint. Hardcoded to staging server for now.
const endpoint = 'https://caregiver-gateway.staging.affinity-project.org/api';

// Utility function to execute XMLHttpRequest using a
// promise wrapper.
async function req(method, url, data, customHeader = '') {
	try {
		const axiosConfig = {
			method,
			url: endpoint + url,
			data,
			responseType: 'json',
			headers: {
				'Content-Type': 'application/json',
				'Api-Key': customHeader,
			},
		};

		const response = await axios(axiosConfig);

		return response.data;
	} catch (error) {
		throw error;
		/*
		return {
			status: error.response.status,
			error: error.message,
		};
		*/
	}
}

// Affinity SDK client.
// To facilitate sync and async usage all methods return
// standard JS promises.
class Affinity {
	constructor(apiKey) {
		this.apiKey = apiKey;
	}

	// Returns a newly generated PIN value.
	GetMaterial() {
		return req('GET', '/did/material', null, this.apiKey);
	}

	// Send confirmation code to email
	MaterialReset(did) {
		let data = {
			did: did,
		};
		return req('POST', '/material/userMaterialReset', data, this.apiKey);
	}

	// Persists new material
	MaterialResetConfirm(did, material, confirmation) {
		let data = {
			did: did,
			material: material,
			confirmation: confirmation,
		};
		return req('POST', '/material/userMaterialResetConfirm', data, this.apiKey);
	}

	// Create a new DID with given PIN value.
	CreateDID(pin, email) {
		let data = {
			material: pin,
			branchManagerEmail: email,
		};
		return req('POST', '/did/material', data, this.apiKey);
	}

	// Returns the JSON encoded document associated with the
	// provided DID, if available.
	Resolve(did) {
		let data = {
			did: did,
		};
		return req('PUT', '/did/resolve', data, this.apiKey);
	}

	// Authenticate and existing user with DID and PIN. The VCs from
	// the wallet are returned, if any.
	Authenticate(did, pin) {
		let data = {
			did: did,
			material: pin,
		};
		return req('PUT', '/authentications', data, this.apiKey);
	}

	// Issue a new verifiable credential.
	// - Issuer must contain: "did" and "pin" properties.
	// - Subject is the user's "did"
	// - Payload must be the data to be contained in the credential,
	//   with the proper schema.
	// Example:
	//   let subject = "did:elem:EiDmiLEqyyzCyComEQALl3lsGPhpB4hcPVYrQnJZ-wkO9Q";
	//   let payload = {};
	//   let issuer = {
	//     did: "did:elem:EiDDxkcB1XV4w_NvZrw3E2E6-5YmQGNE3_bddxP88QavWg",
	//     pin: "OX3POH"
	//   }
	//   affinity.IssueVC(issuer, subject, payload).then((response) => {
	//     console.log(response);
	//   })
	IssueVC(issuer, subject, payload) {
		let data = {
			did: subject,
			issuerDid: issuer.did,
			issuerPin: issuer.pin,
			payload: payload,
		};

		return req('POST', '/vc/issue', data, this.apiKey);
	}

	// Verify the validity of an existing credential.
	VerifyVC(vc) {
		let data = {
			credentials: [vc],
			options: {},
		};
		return req('PUT', '/vc/verify', data, this.apiKey);
	}

	// Store a new verifiable credential on the user's wallet.
	// - Subject is the user's DID
	// - PIN is the active authentication material
	// - VC is the credential to store
	StoreVC(subject, pin, vc) {
		let data = {
			did: subject,
			material: pin,
			vc: vc,
		};
		return req('POST', '/vc/store', data, this.apiKey);
	}
}

module.exports = Affinity;
