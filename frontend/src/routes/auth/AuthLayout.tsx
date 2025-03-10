import { Box, Stack } from "@mui/material";
import { Outlet } from "react-router-dom";
import NavigationBar from "../../components/navigation/NavigationBar";
import Overlay from "../../assets/background/overlay.jpg";

export default function AuthLayout() {
  return (
    <Stack
      height="100vh"
      display="flex"
      flexDirection="column"
      sx={{
        backgroundImage: {
          sm: `url(${Overlay})`,
          xs: "none",
        },
        backgroundSize: "cover",
        backgroundPosition: "center",
      }}
    >
      <NavigationBar />
      <Stack
        flexGrow={1}
        display="flex"
        justifyContent={{
          xs: "flex-start",
          sm: "center",
        }}
        alignItems="center"
      >
        <Box
          sx={{
            minWidth: {
              xs: "100%",
              sm: "420px",
            },
            height: {
              xs: "100%",
              sm: "auto",
            },
            marginX: "auto",
            marginY: {
              xs: 5,
              sm: 0,
            },
            backgroundColor: "var(--palette-background-default);",
            borderRadius: {
              xs: 0,
              sm: 2,
            },
            p: {
              xs: 2,
              sm: 5,
            },
          }}
        >
          <Outlet />
        </Box>
      </Stack>
    </Stack>
  );
}
