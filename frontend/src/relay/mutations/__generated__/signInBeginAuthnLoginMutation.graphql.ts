/**
 * @generated SignedSource<<6ee97b677fc45f9b66afa12c1a548884>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type signInBeginAuthnLoginMutation$variables = Record<PropertyKey, never>;
export type signInBeginAuthnLoginMutation$data = {
  readonly beginAuthnLogin: {
    readonly credentialRequest: string;
  };
};
export type signInBeginAuthnLoginMutation = {
  response: signInBeginAuthnLoginMutation$data;
  variables: signInBeginAuthnLoginMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "alias": null,
    "args": null,
    "concreteType": "AuthnLoginResponse",
    "kind": "LinkedField",
    "name": "beginAuthnLogin",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "credentialRequest",
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "signInBeginAuthnLoginMutation",
    "selections": (v0/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "signInBeginAuthnLoginMutation",
    "selections": (v0/*: any*/)
  },
  "params": {
    "cacheID": "57810a19f8ef3e3ae9097c0caf29be24",
    "id": null,
    "metadata": {},
    "name": "signInBeginAuthnLoginMutation",
    "operationKind": "mutation",
    "text": "mutation signInBeginAuthnLoginMutation {\n  beginAuthnLogin {\n    credentialRequest\n  }\n}\n"
  }
};
})();

(node as any).hash = "72dd30adb9fb6da9519986a1f9b1b623";

export default node;
