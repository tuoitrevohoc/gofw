/**
 * @generated SignedSource<<7ac8014adde37ddd03390701be0690af>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type signUpFinishRegistrationMutation$variables = {
  email: string;
  response: string;
};
export type signUpFinishRegistrationMutation$data = {
  readonly finishAuthnRegistration: {
    readonly email: string;
    readonly id: string;
  };
};
export type signUpFinishRegistrationMutation = {
  response: signUpFinishRegistrationMutation$data;
  variables: signUpFinishRegistrationMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "email"
  },
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "response"
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
      },
      {
        "kind": "Variable",
        "name": "response",
        "variableName": "response"
      }
    ],
    "concreteType": "User",
    "kind": "LinkedField",
    "name": "finishAuthnRegistration",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "id",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "email",
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
    "name": "signUpFinishRegistrationMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "signUpFinishRegistrationMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "960133cae203e2adeb79f7e20f2ec57a",
    "id": null,
    "metadata": {},
    "name": "signUpFinishRegistrationMutation",
    "operationKind": "mutation",
    "text": "mutation signUpFinishRegistrationMutation(\n  $email: String!\n  $response: String!\n) {\n  finishAuthnRegistration(email: $email, response: $response) {\n    id\n    email\n  }\n}\n"
  }
};
})();

(node as any).hash = "187c442a5984907ba84d69fb9bbfe544";

export default node;
