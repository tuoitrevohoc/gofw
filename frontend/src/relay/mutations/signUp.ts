import { graphql, useMutation } from "react-relay";
import { signUpMutation } from "./__generated__/signUpMutation.graphql";

export default function useSignUpMutation() {
  return useMutation<signUpMutation>(graphql`
    mutation signUpMutation($input: SignUpInput!) {
      signUp(input: $input) {
        id
        email
      }
    }
  `);
}
