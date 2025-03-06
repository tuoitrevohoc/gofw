/**
 * @generated SignedSource<<5d547fd2656b96ebcea4dee322c422b4>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type signUpBeginRegistrationMutation$variables = {
  email: string;
};
export type signUpBeginRegistrationMutation$data = {
  readonly beginAuthnRegistration: {
    readonly credentialCreation: string;
  };
};
export type signUpBeginRegistrationMutation = {
  response: signUpBeginRegistrationMutation$data;
  variables: signUpBeginRegistrationMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "email"
  }
],
v1 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "email",
        "variableName": "email"
      }
    ],
    "concreteType": "AuthnRegistrationResponse",
    "kind": "LinkedField",
    "name": "beginAuthnRegistration",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "credentialCreation",
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "signUpBeginRegistrationMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "signUpBeginRegistrationMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "064563107a190f5cb3e7dacd2c205f17",
    "id": null,
    "metadata": {},
    "name": "signUpBeginRegistrationMutation",
    "operationKind": "mutation",
    "text": "mutation signUpBeginRegistrationMutation(\n  $email: String!\n) {\n  beginAuthnRegistration(email: $email) {\n    credentialCreation\n  }\n}\n"
  }
};
})();

(node as any).hash = "09b371bc7b3fb5fbd5990941b54a227f";

export default node;
