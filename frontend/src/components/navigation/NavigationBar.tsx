import { AppBar, Box, Stack } from "@mui/material";

import Icon from "../icons/Icon";

export default function NavigationBar() {
  return (
    <AppBar
      position="sticky"
      sx={{
        backgroundColor: "transparent",
        boxShadow: "none",
        padding: 2,
        borderBottom: {
          xs: `1px solid var(--palette-divider)`,
          sm: "none",
        },
      }}
      elevation={0}
    >
      <Stack direction="row" spacing={2}>
        <Box sx={{ width: 40, height: 40 }}>
          <Icon />
        </Box>
      </Stack>
    </AppBar>
  );
}
