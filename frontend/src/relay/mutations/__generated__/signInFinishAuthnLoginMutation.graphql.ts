/**
 * @generated SignedSource<<77686b3327ce7ae519323366fb83e467>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type signInFinishAuthnLoginMutation$variables = {
  response: string;
};
export type signInFinishAuthnLoginMutation$data = {
  readonly finishAuthnLogin: {
    readonly accessToken: string;
    readonly expiry: number;
  };
};
export type signInFinishAuthnLoginMutation = {
  response: signInFinishAuthnLoginMutation$data;
  variables: signInFinishAuthnLoginMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
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
        "name": "response",
        "variableName": "response"
      }
    ],
    "concreteType": "AccessToken",
    "kind": "LinkedField",
    "name": "finishAuthnLogin",
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
    "name": "signInFinishAuthnLoginMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "signInFinishAuthnLoginMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "0d12fc344660cdab1bfc4895f1b9f111",
    "id": null,
    "metadata": {},
    "name": "signInFinishAuthnLoginMutation",
    "operationKind": "mutation",
    "text": "mutation signInFinishAuthnLoginMutation(\n  $response: String!\n) {\n  finishAuthnLogin(response: $response) {\n    accessToken\n    expiry\n  }\n}\n"
  }
};
})();

(node as any).hash = "0594313c7def2481a6ed682e376185c9";

export default node;
