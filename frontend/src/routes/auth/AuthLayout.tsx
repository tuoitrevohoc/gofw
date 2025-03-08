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
        backgroundImage: `url(${Overlay})`,
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
          md: "center",
        }}
        alignItems="center"
      >
        <Box
          sx={{
            minWidth: 420,
            marginX: "auto",
            marginY: {
              xs: 5,
              md: 0,
            },
            backgroundColor: "var(--palette-background-default);",
            borderRadius: 2,
            p: {
              xs: 2,
              md: 5,
            },
          }}
        >
          <Outlet />
        </Box>
      </Stack>
    </Stack>
  );
}
