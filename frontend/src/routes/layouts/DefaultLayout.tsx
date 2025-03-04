import {
  List,
  ListItem,
  ListItemButton,
  ListItemText,
  Stack,
  Drawer,
} from "@mui/material";
import { Outlet, useLocation } from "react-router-dom";
import NavigationBar from "../../components/navigation/NavigationBar";
import { Link } from "react-router";

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
        py: 1,
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
          gap: 2,
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
    <List>
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
  return (
    <Stack>
      <Drawer
        id="desktop-drawer"
        variant="permanent"
        open={true}
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          display: { xs: "none", md: "block" },
        }}
      >
        <NavigationBar />
        <SideMenu menuItems={menuItems} />
      </Drawer>
      <Outlet />
    </Stack>
  );
}
