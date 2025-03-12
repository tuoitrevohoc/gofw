import { AppBar, Box, IconButton, Stack } from "@mui/material";
import StackBar from "../components/icons/StackBar";
import AccountMenu from "./auth/AccountMenu";

interface TopBarProps {
  onMenuClick: () => void;
}

export default function TopBar({ onMenuClick }: TopBarProps) {
  return (
    <AppBar
      position="sticky"
      sx={{
        width: "100%",
        backdropFilter: "blur(10px)",
        backgroundColor: {
          xs: "var(--palette-background-paper)",
          sm: "transparent",
        },
        boxShadow: "none",
        padding: 2,
        borderBottom: {
          xs: `1px solid var(--palette-divider)`,
          sm: "none",
        },
      }}
      elevation={0}
    >
      <Stack direction="row" alignItems="center" spacing={2}>
        <IconButton
          onClick={onMenuClick}
          sx={{ display: { xs: "block", lg: "none" } }}
        >
          <StackBar />
        </IconButton>
        <Box flexGrow={1} />
        <AccountMenu />
      </Stack>
    </AppBar>
  );
}
