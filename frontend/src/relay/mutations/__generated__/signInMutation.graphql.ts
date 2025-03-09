/**
 * @generated SignedSource<<fef176098344c58c6a98d1e8fe14bcdb>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type SignInInput = {
  email: string;
  password: string;
};
export type signInMutation$variables = {
  input: SignInInput;
};
export type signInMutation$data = {
  readonly signIn: {
    readonly accessToken: string;
    readonly expiry: number;
  };
};
export type signInMutation = {
  response: signInMutation$data;
  variables: signInMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "input"
  }
],
v1 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "input",
        "variableName": "input"
      }
    ],
    "concreteType": "AccessToken",
    "kind": "LinkedField",
    "name": "signIn",
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
    "name": "signInMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "signInMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "a2f3fcc20a42befef9c0e3c9c8dcb80f",
    "id": null,
    "metadata": {},
    "name": "signInMutation",
    "operationKind": "mutation",
    "text": "mutation signInMutation(\n  $input: SignInInput!\n) {\n  signIn(input: $input) {\n    accessToken\n    expiry\n  }\n}\n"
  }
};
})();

(node as any).hash = "04d1063978272c42a017f3f37bffe56f";

export default node;
