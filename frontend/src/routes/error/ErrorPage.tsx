import { useNavigate } from "react-router-dom";
import { Typography, Button, Stack, Box } from "@mui/material";
import NavigationBar from "../../components/navigation/NavigationBar";

import SadPanda from "./SadPanda@2x.png";

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
          <Box>
            <img src={SadPanda} alt="Sad Panda" style={{ width: 256 }} />
          </Box>
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
