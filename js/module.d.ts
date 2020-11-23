declare module affinityClient {
  export class Affinity {
		static Stripe: typeof Affinity;

		constructor(apiKey: string);
  }

  type DidCreation = {
    did: string;
	}

	type FIHR = {
    resourceType: string;
    identifier: string;
    status: string;
    type: {
      valueCodeableConcept: string;
    };
    policyHolder: {
      resourceType: string;
      name: [{ text: string }];
      gender: string;
      birthDate: string;
    };
    beneficiary: {
      resourceType: string;
      name: [{ text: string }];
      gender: string;
      birthDate: string;
    };
    period: {
      start: string;
      end: string;
    };
    payor: [
      {
        identifier: {
          type: string;
          display: string;
        };
      }
    ];
    class: [{ name: string }];
  };

  type Issuer = {
    did: string;
    pin: string;
  };

  type VCIssue = {
    issuer: Issuer;
    subject: string;
    payload: FIHR;
	};

	type VCStore = {
		subject: string;
		pin: string
    vc: any;
  };

  type Authenticate = {
		did: string;
    pin: string;
    apiKey: string;
  };

  type MaterialResetParameters = {
    did: string;
  }

  type MaterialResetConfirmParameters = {
    did: string;
    material: string;
    confirmation: string;
  }

  CreateDID(pin: string, email: string): DidCreation;

	IssueVC(issuer: Issuer, subject: string, payload: FIHR): VCIssue;

  StoreVC(subject: string, pin: string, vc: any): VCStore;

  Authenticate(did: string, pin: string, apiKey: string): Authenticate;
  
  MaterialReset(did: string): MaterialResetParameters;

  MaterialResetConfirm(did: string, material: string, confirmation: string): MaterialResetConfirmParameters;
}
