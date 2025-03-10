import { useMutation, graphql } from "react-relay";
import { signInMutation } from "./__generated__/signInMutation.graphql";
import { signInBeginAuthnLoginMutation } from "./__generated__/signInBeginAuthnLoginMutation.graphql";
import { signInFinishAuthnLoginMutation } from "./__generated__/signInFinishAuthnLoginMutation.graphql";
export function useSignInMutation() {
  return useMutation<signInMutation>(graphql`
    mutation signInMutation($input: SignInInput!) {
      signIn(input: $input) {
        accessToken
        expiry
      }
    }
  `);
}

export function useSignInBeginMutation() {
  return useMutation<signInBeginAuthnLoginMutation>(graphql`
    mutation signInBeginAuthnLoginMutation {
      beginAuthnLogin {
        credentialRequest
      }
    }
  `);
}

export function useSignInFinishMutation() {
  return useMutation<signInFinishAuthnLoginMutation>(graphql`
    mutation signInFinishAuthnLoginMutation($response: String!) {
      finishAuthnLogin(response: $response) {
        accessToken
        expiry
      }
    }
  `);
}
