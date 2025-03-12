import {
  Card,
  CardContent,
  Divider,
  TextField,
  Typography,
  Button,
  Stack,
} from "@mui/material";
import Page from "../../components/layout/Page";
import { useState } from "react";
import { graphql } from "react-relay";
import { useMutation } from "react-relay";
import { RestaurantEditorSaveRestaurantMutation } from "./__generated__/RestaurantEditorSaveRestaurantMutation.graphql";

interface Restaurant {
  id: string | null;
  name: string;
  address: string;
}

interface RestaurantEditorProps {
  restaurant: Restaurant;
}

export default function RestaurantEditor({
  restaurant = { id: null, name: "", address: "" },
}: RestaurantEditorProps) {
  const [data, setData] = useState<Restaurant>(restaurant);

  const [commitSaveRestaurant, isSaving] =
    useMutation<RestaurantEditorSaveRestaurantMutation>(graphql`
      mutation RestaurantEditorSaveRestaurantMutation(
        $input: SaveRestaurantInput!
      ) {
        saveRestaurant(input: $input) {
          id
        }
      }
    `);

  const handleSaveRestaurant = () => {
    commitSaveRestaurant({
      variables: {
        input: { id: data.id, name: data.name, address: data.address },
      },
      onCompleted: (response) => {
        if (response.saveRestaurant) {
          setData({
            ...data,
            id: response.saveRestaurant.id,
          });
        }
      },
    });
  };

  return (
    <Page
      title="Add Restaurant"
      breadcrumbs={[
        { label: "Dashboard", path: "/" },
        { label: "Restaurants", path: "/restaurants" },
        { label: "Add Restaurant", path: "/restaurants/add" },
      ]}
    >
      <Card>
        <CardContent>
          <Typography variant="h6">Restaurant</Typography>
          <Typography variant="body2">
            Add a new restaurant to your account.
          </Typography>
        </CardContent>
        <Divider />
        <CardContent sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
          <TextField
            value={data.name}
            onChange={(e) => setData({ ...data, name: e.currentTarget.value })}
            fullWidth
            label="Name"
          />
          <TextField
            value={data.address}
            onChange={(e) =>
              setData({ ...data, address: e.currentTarget.value })
            }
            fullWidth
            label="Address"
          />
        </CardContent>
        <Divider />
      </Card>
      <Stack direction="row" justifyContent="flex-end">
        <Button
          variant="contained"
          size="large"
          color="primary"
          loading={isSaving}
          onClick={handleSaveRestaurant}
        >
          Save
        </Button>
      </Stack>
    </Page>
  );
}
