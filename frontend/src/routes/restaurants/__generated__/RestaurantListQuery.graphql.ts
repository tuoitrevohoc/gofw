/**
 * @generated SignedSource<<20d9f2d5c025bedbc88b13958c769df8>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
export type RestaurantListQuery$variables = Record<PropertyKey, never>;
export type RestaurantListQuery$data = {
  readonly restaurants: {
    readonly edges: ReadonlyArray<{
      readonly node: {
        readonly address: string;
        readonly id: string;
        readonly name: string;
      } | null | undefined;
    } | null | undefined> | null | undefined;
    readonly pageInfo: {
      readonly hasNextPage: boolean;
      readonly startCursor: any | null | undefined;
    };
    readonly totalCount: number;
  };
};
export type RestaurantListQuery = {
  response: RestaurantListQuery$data;
  variables: RestaurantListQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "alias": null,
    "args": null,
    "concreteType": "RestaurantConnection",
    "kind": "LinkedField",
    "name": "restaurants",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "concreteType": "RestaurantEdge",
        "kind": "LinkedField",
        "name": "edges",
        "plural": true,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "Restaurant",
            "kind": "LinkedField",
            "name": "node",
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
                "name": "name",
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "address",
                "storageKey": null
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "totalCount",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "concreteType": "PageInfo",
        "kind": "LinkedField",
        "name": "pageInfo",
        "plural": false,
        "selections": [
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "hasNextPage",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "startCursor",
            "storageKey": null
          }
        ],
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
    "name": "RestaurantListQuery",
    "selections": (v0/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "RestaurantListQuery",
    "selections": (v0/*: any*/)
  },
  "params": {
    "cacheID": "83798ab4bb03ad585ef3705483e65a74",
    "id": null,
    "metadata": {},
    "name": "RestaurantListQuery",
    "operationKind": "query",
    "text": "query RestaurantListQuery {\n  restaurants {\n    edges {\n      node {\n        id\n        name\n        address\n      }\n    }\n    totalCount\n    pageInfo {\n      hasNextPage\n      startCursor\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "023d141c02ff01a0e4b9c63fe756dcda";

export default node;
