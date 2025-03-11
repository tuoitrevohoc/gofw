import React from "react";
import { useNavigate } from "react-router-dom";
import { Container, Typography, Button, Box, Stack } from "@mui/material";
import NavigationBar from "../../components/navigation/NavigationBar";

function ErrorPage() {
  const navigate = useNavigate();

  const handleRedirect = () => {
    navigate("/");
  };

  return (
    <>
      <Stack display="flex" height="100vh" gap={2} direction="column">
        <NavigationBar />
        <Stack
          flexGrow={1}
          display="flex"
          justifyContent="center"
          alignItems="center"
          gap={3}
        >
          <Typography variant="h2">Error</Typography>
          <Typography variant="body1">
            Something went wrong. Please try again later.
          </Typography>
          <Typography variant="body2">
            Please check the server if it's running.
          </Typography>
          <Button
            size="large"
            variant="contained"
            color="primary"
            onClick={handleRedirect}
          >
            Go to Home
          </Button>
        </Stack>
      </Stack>
    </>
  );
}

export default ErrorPage;
