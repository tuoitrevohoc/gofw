/**
 * @generated SignedSource<<8ab0b9820fa37d3190f40630b2f961a4>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
export type routesQuery$variables = Record<PropertyKey, never>;
export type routesQuery$data = {
  readonly viewer: {
    readonly isAuthenticated: boolean;
    readonly userId: string | null | undefined;
  };
};
export type routesQuery = {
  response: routesQuery$data;
  variables: routesQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "alias": null,
    "args": null,
    "concreteType": "Viewer",
    "kind": "LinkedField",
    "name": "viewer",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "userId",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "isAuthenticated",
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
    "name": "routesQuery",
    "selections": (v0/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "routesQuery",
    "selections": (v0/*: any*/)
  },
  "params": {
    "cacheID": "7fd9f4c9052ab902ac7c619c6f872213",
    "id": null,
    "metadata": {},
    "name": "routesQuery",
    "operationKind": "query",
    "text": "query routesQuery {\n  viewer {\n    userId\n    isAuthenticated\n  }\n}\n"
  }
};
})();

(node as any).hash = "9ffd2ec12abaae09bc69bfe109862a24";

export default node;
