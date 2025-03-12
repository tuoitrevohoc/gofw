/**
 * @generated SignedSource<<7b709b20da3f6f39bd3204e3ac336ecc>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type SaveRestaurantInput = {
  address: string;
  id?: string | null | undefined;
  name: string;
};
export type RestaurantEditorSaveRestaurantMutation$variables = {
  input: SaveRestaurantInput;
};
export type RestaurantEditorSaveRestaurantMutation$data = {
  readonly saveRestaurant: {
    readonly id: string;
  };
};
export type RestaurantEditorSaveRestaurantMutation = {
  response: RestaurantEditorSaveRestaurantMutation$data;
  variables: RestaurantEditorSaveRestaurantMutation$variables;
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
    "concreteType": "Restaurant",
    "kind": "LinkedField",
    "name": "saveRestaurant",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "id",
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
    "name": "RestaurantEditorSaveRestaurantMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "RestaurantEditorSaveRestaurantMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "4f99538fe69e93acfbfad812003d2056",
    "id": null,
    "metadata": {},
    "name": "RestaurantEditorSaveRestaurantMutation",
    "operationKind": "mutation",
    "text": "mutation RestaurantEditorSaveRestaurantMutation(\n  $input: SaveRestaurantInput!\n) {\n  saveRestaurant(input: $input) {\n    id\n  }\n}\n"
  }
};
})();

(node as any).hash = "9562ea3fe36d3271cae1b8d0ac756ebe";

export default node;
