/**
 * @generated SignedSource<<cc95a4771237abc036975a10fe5f35d4>>
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
    readonly accessToken: string;
    readonly expiry: number;
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
    "concreteType": "AccessToken",
    "kind": "LinkedField",
    "name": "finishAuthnRegistration",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "accessToken",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "expiry",
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
    "cacheID": "f9de7e3f503eee5723bb7d1e0cfe505e",
    "id": null,
    "metadata": {},
    "name": "signUpFinishRegistrationMutation",
    "operationKind": "mutation",
    "text": "mutation signUpFinishRegistrationMutation(\n  $email: String!\n  $response: String!\n) {\n  finishAuthnRegistration(email: $email, response: $response) {\n    accessToken\n    expiry\n  }\n}\n"
  }
};
})();

(node as any).hash = "76a61810a0799e2aeba3d4ab1fd01879";

export default node;
