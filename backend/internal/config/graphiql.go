package config

import "net/http"

const GraphiQL = `
	<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <!-- Include GraphiQL CSS -->
    <link href="https://unpkg.com/graphiql/graphiql.min.css" rel="stylesheet" />
    <title>GraphiQL Example</title>
  </head>
  <body style="margin: 0; height: 100vh;">
    <!-- Container for GraphiQL -->
    <div id="graphiql" style="height: 100vh;"></div>

    <!-- Include React and ReactDOM -->
    <script
		src="https://cdn.jsdelivr.net/npm/react@18.2.0/umd/react.production.min.js"
		crossorigin="anonymous"
	></script>
	<script
		src="https://cdn.jsdelivr.net/npm/react-dom@18.2.0/umd/react-dom.production.min.js"
		crossorigin="anonymous"
	></script>
    <!-- Include GraphiQL -->
    <script crossorigin src="https://cdn.jsdelivr.net/npm/graphiql@3.7.0/graphiql.min.js"></script>

    <script>
	function saveToken(viewer) {
		sessionStorage.setItem("accessToken", viewer.accessToken);
		sessionStorage.setItem("expiresAt", viewer.expiry.toString());
	}
	async function refreshToken() {
		const response = await fetch("/api/refresh-token", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			credentials: "include",
		});

		if (!response.ok) {
			return null;
		}

		const data = await response.json();
		saveToken(data);
		return data.accessToken;
	}

	async function getAccessToken() {
		const accessToken = sessionStorage.getItem("accessToken");
		const expiresAt = sessionStorage.getItem("expiresAt");

		console.log(accessToken, expiresAt);

		if (!accessToken || !expiresAt) {
			return null;
		}

		if (Date.now() > parseInt(expiresAt) * 1000) {
			return refreshToken();
		}

		return accessToken;
	}
      // Define a fetcher function for GraphQL queries.
	async function graphQLFetcher(graphQLParams) {
		const accessToken = await getAccessToken();	
		const headers = {
			"Content-Type": "application/json",
			"Authorization": "Bearer " + accessToken,
		}

        return fetch('/graphql', { // Update this URL to your GraphQL endpoint.
			method: 'post',
			headers: headers,
			body: JSON.stringify(graphQLParams),
        }).then(response => response.json());
	}

      // Render the GraphiQL React component.
	ReactDOM.render(
		React.createElement(GraphiQL, { fetcher: graphQLFetcher }),
		document.getElementById('graphiql')
	);
    </script>
  </body>
</html>

`

func GraphiQLHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(GraphiQL))
	})
}
