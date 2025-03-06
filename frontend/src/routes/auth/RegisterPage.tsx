import {
  Button,
  Divider,
  Link,
  Stack,
  TextField,
  Typography,
  Alert,
} from "@mui/material";
import { useState } from "react";
import { Link as RouterLink, useNavigate } from "react-router-dom";
import useSignUpMutation from "../../relay/mutations/signUp";
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

export default function RegisterPage() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [commitSignUp, loading] = useSignUpMutation();

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
      onError: (error) => {
        const errorMessage = getErrorMessage(error);
        setError(errorMessage);
      },
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
