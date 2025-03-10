import {
  Button,
  Link,
  Stack,
  TextField,
  Typography,
  Alert,
  IconButton,
} from "@mui/material";
import { useCallback, useState } from "react";
import { Link as RouterLink, useNavigate } from "react-router-dom";
import getErrorMessage from "../../relay/getErrorMessage";
import { handleLogin } from "../../auth";
import {
  useSignInMutation,
  useSignInBeginMutation,
  useSignInFinishMutation,
} from "../../relay/mutations/signIn";
import { KeyRounded } from "@mui/icons-material";
import { stringToArrayBuffer, base64Encode } from "../../utils/encoding";

function isValidateForm(email: string, password: string) {
  return email.length > 0 && password.length > 0;
}

function encodeCredential(credential: PublicKeyCredential) {
  const response = credential.response as AuthenticatorAssertionResponse;
  return JSON.stringify({
    id: credential.id,
    type: credential.type,
    rawId: base64Encode(credential.rawId),
    response: {
      authenticatorData: base64Encode(response.authenticatorData),
      clientDataJSON: base64Encode(response.clientDataJSON),
      signature: base64Encode(response.signature),
      userHandle: response.userHandle
        ? base64Encode(response.userHandle)
        : null,
    },
  });
}

export default function LoginPage() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [commitSignIn, isSigningIn] = useSignInMutation();
  const [commitSignInBegin, isSigningInBegin] = useSignInBeginMutation();
  const [commitSignInFinish] = useSignInFinishMutation();

  const onError = useCallback(
    function (error: Error) {
      const errorMessage = getErrorMessage(error);
      setError(errorMessage);
    },
    [setError]
  );

  const onSignIn = useCallback(() => {
    commitSignIn({
      variables: {
        input: {
          email,
          password,
        },
      },
      onCompleted: (data) => {
        handleLogin(data.signIn);
        navigate("/");
      },
      onError,
    });
  }, [email, password, commitSignIn, navigate, onError]);

  const onSignInBegin = useCallback(() => {
    async function startLogin(requestJSON: string) {
      try {
        const request = JSON.parse(requestJSON);
        console.log(request);
        request.publicKey.challenge = stringToArrayBuffer(
          request.publicKey.challenge
        );

        if (request.publicKey?.user?.id) {
          request.publicKey.user.id = stringToArrayBuffer(
            request.publicKey.user.id
          );
        }

        const credential = await navigator.credentials.get(request);

        if (!credential) {
          return;
        }

        const encodedCredential = encodeCredential(
          credential as PublicKeyCredential
        );

        commitSignInFinish({
          variables: { response: encodedCredential },
          onCompleted: (data) => {
            handleLogin(data.finishAuthnLogin);
            navigate("/");
          },
          onError,
        });
      } catch (error) {
        console.error(error);
        setError("Failed to create credential");
      }
    }

    commitSignInBegin({
      variables: {},
      onCompleted: (data) => {
        startLogin(data.beginAuthnLogin.credentialRequest);
      },
      onError,
    });
  }, [commitSignInBegin, commitSignInFinish, onError, navigate]);

  return (
    <Stack direction="column" gap={1.5} alignItems="center" width="100%">
      <Typography variant="h5">Login</Typography>
      <Typography variant="body2" color="text.secondary">
        First time here?{" "}
        <Link to="/auth/register" component={RouterLink}>
          Create an account
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
          slotProps={{
            input: {
              endAdornment: (
                <IconButton onClick={onSignInBegin}>
                  <KeyRounded />
                </IconButton>
              ),
            },
          }}
        />

        {email.length > 0 && (
          <>
            <TextField
              fullWidth
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              label="Password"
            />
            <Button
              disabled={!isValidateForm(email, password)}
              fullWidth
              variant="outlined"
              size="large"
              onClick={onSignIn}
              loading={isSigningIn}
            >
              Login
            </Button>
          </>
        )}
        <Button
          fullWidth
          variant="outlined"
          size="large"
          onClick={onSignInBegin}
          loading={isSigningInBegin}
        >
          Login with passkey
        </Button>
      </Stack>
    </Stack>
  );
}
