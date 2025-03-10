import { Box, IconButton, Popover } from "@mui/material";
import { AccountCircle } from "@mui/icons-material";
import { useRef, useState } from "react";
import LogoutButton from "./LogoutButton";

export default function AccountMenu() {
  const [menuOpened, setMenuOpened] = useState(false);
  const anchorEl = useRef<HTMLButtonElement>(null);

  return (
    <>
      <IconButton onClick={() => setMenuOpened(!menuOpened)} ref={anchorEl}>
        <AccountCircle />
      </IconButton>

      <Popover
        open={menuOpened}
        onClose={() => setMenuOpened(false)}
        anchorEl={anchorEl.current}
        anchorOrigin={{
          vertical: "bottom",
          horizontal: "right",
        }}
        transformOrigin={{
          vertical: "top",
          horizontal: "right",
        }}
        slotProps={{
          paper: {
            sx: { width: 200 },
          },
        }}
      >
        <Box sx={{ p: 1 }}>
          <LogoutButton />
        </Box>
      </Popover>
    </>
  );
}
