import { Add } from "@mui/icons-material";
import { Button, Stack, Typography } from "@mui/material";

export default function RestaurantList() {
  return (
    <Stack gap={2}>
      <Typography variant="h6">Restaurants</Typography>

      <Stack gap={2}>
        <Typography variant="body1">
          You currently don't have any restaurants. Please create one.
        </Typography>
        <Stack direction="row">
          <Button variant="contained" startIcon={<Add />}>
            Add
          </Button>
        </Stack>
      </Stack>
    </Stack>
  );
}
