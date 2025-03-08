import {
  Environment,
  Network,
  RecordSource,
  Store,
  FetchFunction,
  GraphQLResponse,
  RequestParameters,
  Variables,
} from "relay-runtime";
import { getAccessToken } from "../auth";

const fetchQuery: FetchFunction = async (
  operation: RequestParameters,
  variables: Variables
): Promise<GraphQLResponse> => {
  // get the token
  const token = await getAccessToken();

  const response = await fetch("/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({
      query: operation.text,
      variables,
    }),
  });

  return response.json();
};

const environment = new Environment({
  network: Network.create(fetchQuery),
  store: new Store(new RecordSource()),
});

export default environment;
