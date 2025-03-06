import {
  Button,
  Divider,
  Link,
  Stack,
  TextField,
  Typography,
  Alert,
} from "@mui/material";
import { useCallback, useState } from "react";
import { Link as RouterLink, useNavigate } from "react-router-dom";
import {
  useSignUpMutation,
  useBeginAuthnRegistrationMutation,
  useFinishAuthnRegistrationMutation,
} from "../../relay/mutations/signUp";
import getErrorMessage from "../../relay/getErrorMessage";

function isValidateForm(
  email: string,
  password: string,
  confirmPassword: string,
  usePassword: boolean
) {
  if (usePassword) {
    return (
      password.length > 0 &&
      confirmPassword.length > 0 &&
      password === confirmPassword
    );
  }
  return email.length > 0;
}

function stringToArrayBuffer(binaryString: string) {
  return Uint8Array.from(binaryString, (c) => c.charCodeAt(0)).buffer;
}

function base64Encode(binary: ArrayBuffer) {
  const binaryString = String.fromCharCode(...new Uint8Array(binary));
  return btoa(binaryString)
    .replace(/\+/g, "-")
    .replace(/\//g, "_")
    .replace(/=/g, "");
}

function encodeCredential(credential: PublicKeyCredential) {
  const response = credential.response as AuthenticatorAttestationResponse;
  return JSON.stringify({
    id: credential.id,
    type: credential.type,
    rawId: base64Encode(credential.rawId),
    response: {
      clientDataJSON: base64Encode(response.clientDataJSON),
      attestationObject: base64Encode(response.attestationObject),
    },
  });
}

export default function RegisterPage() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [commitSignUp, loading] = useSignUpMutation();
  const [commitBeginAuthnRegistration, beginningAuthnRegistration] =
    useBeginAuthnRegistrationMutation();

  const [commitFinishAuthnRegistration] = useFinishAuthnRegistrationMutation();

  async function startRegistration(creatorOptionsJSON: string) {
    try {
      const creatorOptions = JSON.parse(creatorOptionsJSON);
      creatorOptions.publicKey.challenge = stringToArrayBuffer(
        creatorOptions.publicKey.challenge
      );
      creatorOptions.publicKey.user.id = stringToArrayBuffer(
        creatorOptions.publicKey.user.id
      );

      const credential = await navigator.credentials.create(creatorOptions);

      console.log(credential);
      console.log(JSON.stringify(credential));

      commitFinishAuthnRegistration({
        variables: {
          email,
          response: encodeCredential(
            credential as unknown as PublicKeyCredential
          ),
        },
        onCompleted: () => {
          navigate("/");
        },
        onError,
      });
    } catch (error) {
      console.error(error);
      setError("Failed to create credential");
    }
  }

  const onError = useCallback(
    function (error: Error) {
      const errorMessage = getErrorMessage(error);
      setError(errorMessage);
    },
    [setError]
  );

  function handleSignUp() {
    setError(null);
    commitSignUp({
      variables: {
        input: {
          email,
          password,
        },
      },
      onCompleted: () => {
        navigate("/");
      },
      onError,
    });
  }

  function handleBeginAuthnRegistration() {
    commitBeginAuthnRegistration({
      variables: {
        email,
      },
      onCompleted: (data) => {
        startRegistration(data.beginAuthnRegistration.credentialCreation);
      },
      onError,
    });
  }

  return (
    <Stack direction="column" gap={1.5} alignItems="center" width="100%">
      <Typography variant="h5">Register</Typography>
      <Typography variant="body2" color="text.secondary">
        Have an account?{" "}
        <Link to="/auth/login" component={RouterLink}>
          Sign in
        </Link>
      </Typography>
      {error && (
        <Alert severity="error" sx={{ width: "100%" }}>
          {error}
        </Alert>
      )}
      <Stack direction="column" gap={1.5} width="100%">
        <TextField
          fullWidth
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          label="Email"
        />
        <Button
          disabled={!isValidateForm(email, password, confirmPassword, false)}
          fullWidth
          variant="contained"
          size="large"
          onClick={handleBeginAuthnRegistration}
          loading={beginningAuthnRegistration}
        >
          Continue with no password
        </Button>
        <Divider>
          <Typography variant="body2" color="text.secondary">
            Or use a password
          </Typography>
        </Divider>
        <>
          <TextField
            fullWidth
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            label="Password"
          />
          <TextField
            fullWidth
            type="password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            error={password !== confirmPassword}
            label="Confirm Password"
          />
        </>
        <Button
          disabled={!isValidateForm(email, password, confirmPassword, true)}
          fullWidth
          variant="outlined"
          size="large"
          onClick={handleSignUp}
          loading={loading}
        >
          Register
        </Button>
      </Stack>
    </Stack>
  );
}
