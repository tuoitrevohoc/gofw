import { graphql, useMutation } from "react-relay";
import { signUpMutation } from "./__generated__/signUpMutation.graphql";
import { signUpBeginRegistrationMutation } from "./__generated__/signUpBeginRegistrationMutation.graphql";
import { signUpFinishRegistrationMutation } from "./__generated__/signUpFinishRegistrationMutation.graphql";
export function useSignUpMutation() {
  return useMutation<signUpMutation>(graphql`
    mutation signUpMutation($input: SignUpInput!) {
      signUp(input: $input) {
        id
        email
      }
    }
  `);
}

export function useBeginAuthnRegistrationMutation() {
  return useMutation<signUpBeginRegistrationMutation>(graphql`
    mutation signUpBeginRegistrationMutation($email: String!) {
      beginAuthnRegistration(email: $email) {
        credentialCreation
      }
    }
  `);
}

export function useFinishAuthnRegistrationMutation() {
  return useMutation<signUpFinishRegistrationMutation>(graphql`
    mutation signUpFinishRegistrationMutation(
      $email: String!
      $response: String!
    ) {
      finishAuthnRegistration(email: $email, response: $response) {
        id
        email
      }
    }
  `);
}
