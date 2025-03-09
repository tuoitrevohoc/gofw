import { useMutation, graphql } from "react-relay";
import { signInMutation } from "./__generated__/signInMutation.graphql";

function useSignInMutation() {
  return useMutation<signInMutation>(graphql`
    mutation signInMutation($input: SignInInput!) {
      signIn(input: $input) {
        accessToken
        expiry
      }
    }
  `);
}

export default useSignInMutation;
