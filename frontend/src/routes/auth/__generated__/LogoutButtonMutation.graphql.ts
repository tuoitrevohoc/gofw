/**
 * @generated SignedSource<<7c7f1c79fa1f6bd4fa02cb9b368d1c2d>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type LogoutButtonMutation$variables = Record<PropertyKey, never>;
export type LogoutButtonMutation$data = {
  readonly signOut: boolean;
};
export type LogoutButtonMutation = {
  response: LogoutButtonMutation$data;
  variables: LogoutButtonMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "alias": null,
    "args": null,
    "kind": "ScalarField",
    "name": "signOut",
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "LogoutButtonMutation",
    "selections": (v0/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "LogoutButtonMutation",
    "selections": (v0/*: any*/)
  },
  "params": {
    "cacheID": "7a6bddb7c902c561ca119f2ed0d9f710",
    "id": null,
    "metadata": {},
    "name": "LogoutButtonMutation",
    "operationKind": "mutation",
    "text": "mutation LogoutButtonMutation {\n  signOut\n}\n"
  }
};
})();

(node as any).hash = "9e8d71011324e056372711b42b2f5683";

export default node;
