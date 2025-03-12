import { Add } from "@mui/icons-material";
import { Button, Stack, Typography, Link } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";
import Page from "../../components/layout/Page";

export default function RestaurantList() {
  return (
    <Page
      title="Restaurants"
      breadcrumbs={[
        { label: "Dashboard", path: "/" },
        { label: "Restaurants", path: "/restaurants" },
      ]}
    >
      <Typography variant="body1">
        You currently don't have any restaurants. Please create one.
      </Typography>
      <Stack direction="row">
        <Link component={RouterLink} to="/restaurants/add">
          <Button variant="contained" startIcon={<Add />}>
            Add
          </Button>
        </Link>
      </Stack>
    </Page>
  );
}
