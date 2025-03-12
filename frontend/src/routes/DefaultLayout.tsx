import {
  List,
  ListItem,
  ListItemButton,
  ListItemText,
  Stack,
  Drawer,
  Box,
} from "@mui/material";
import { Outlet, useLocation } from "react-router-dom";
import NavigationBar from "../components/navigation/NavigationBar";
import { Link } from "react-router";
import TopBar from "./TopBar";
import { useState } from "react";

export interface MenuItem {
  icon: string;
  label: string;
  path: string;
}

interface DefaultLayoutProps {
  menuItems: MenuItem[];
}

const drawerWidth = 300;

function SideMenuItem({
  item,
  isActive,
}: {
  item: MenuItem;
  isActive: boolean;
}) {
  return (
    <ListItem
      key={item.path}
      sx={{
        pl: 2,
        py: 0.5,
        pr: 1.5,
        borderRadius: 0.75,
        typography: "body2",
        fontWeight: "fontWeightMedium",
        color: "var(--layout-nav-item-color)",
        minHeight: "var(--layout-nav-item-height)",
        width: drawerWidth,
      }}
    >
      <ListItemButton
        sx={{
          gap: 1,
          borderRadius: 0.75,
          ...(isActive && {
            fontWeight: "fontWeightSemiBold",
            bgcolor: "var(--layout-nav-item-active-bg)",
            color: "var(--layout-nav-item-active-color)",
            "&:hover": {
              bgcolor: "var(--layout-nav-item-hover-bg)",
            },
          }),
        }}
        component={Link}
        to={item.path}
      >
        <img src={item.icon} alt={item.label} />
        <ListItemText primary={item.label} />
      </ListItemButton>
    </ListItem>
  );
}

function SideMenu({ menuItems }: DefaultLayoutProps) {
  const location = useLocation();
  const currentPath = location.pathname;

  return (
    <List sx={{ gap: 0 }}>
      {menuItems.map((item) => (
        <SideMenuItem
          key={item.path}
          item={item}
          isActive={currentPath === item.path}
        />
      ))}
    </List>
  );
}

export default function DefaultLayout({ menuItems }: DefaultLayoutProps) {
  const [menuOpen, setMenuOpen] = useState(false);

  const handleDrawerToggle = () => {
    setMenuOpen(!menuOpen);
  };

  return (
    <Stack direction="row" height="100vh">
      <Drawer
        id="desktop-drawer"
        variant="permanent"
        open={true}
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          display: { xs: "none", lg: "block" },
        }}
      >
        <NavigationBar />
        <SideMenu menuItems={menuItems} />
      </Drawer>
      <Drawer
        id="mobile-drawer"
        variant="temporary"
        open={menuOpen}
        onClose={handleDrawerToggle}
      >
        <NavigationBar />
        <SideMenu menuItems={menuItems} />
      </Drawer>
      <Stack
        flexGrow={1}
        display="flex"
        flexDirection="column"
        justifyContent="flex-start"
        alignItems="flex-start"
      >
        <TopBar onMenuClick={handleDrawerToggle} />
        <Box
          flexGrow={1}
          width="100%"
          paddingX={4}
          paddingY={2}
          sx={{ maxWidth: 1000, marginX: "auto" }}
        >
          <Outlet />
        </Box>
      </Stack>
    </Stack>
  );
}
