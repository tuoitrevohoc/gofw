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
import useSignInMutation from "../../relay/mutations/signIn";
import { KeyRounded } from "@mui/icons-material";

function isValidateForm(email: string, password: string) {
  return email.length > 0 && password.length > 0;
}

export default function LoginPage() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);

  const [commitSignIn, isSigningIn] = useSignInMutation();

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
                <IconButton>
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
        <Button fullWidth variant="outlined" size="large" onClick={onSignIn}>
          Login with passkey
        </Button>
      </Stack>
    </Stack>
  );
}
