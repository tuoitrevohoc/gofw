import { Add, Delete, Edit } from "@mui/icons-material";
import {
  Button,
  Stack,
  Typography,
  Link,
  Card,
  CardContent,
} from "@mui/material";
import { Link as RouterLink } from "react-router-dom";
import Page from "../../components/layout/Page";
import { graphql, useLazyLoadQuery } from "react-relay";
import { RestaurantListQuery } from "./__generated__/RestaurantListQuery.graphql";
import { DataGrid } from "@mui/x-data-grid";
export default function RestaurantList() {
  const { restaurants } = useLazyLoadQuery<RestaurantListQuery>(
    graphql`
      query RestaurantListQuery {
        restaurants {
          edges {
            node {
              id
              name
              address
            }
          }
          totalCount
          pageInfo {
            hasNextPage
            startCursor
          }
        }
      }
    `,
    {}
  );

  return (
    <Page
      title="Restaurants"
      breadcrumbs={[
        { label: "Dashboard", path: "/" },
        { label: "Restaurants", path: "/restaurants" },
      ]}
    >
      {restaurants.totalCount === 0 ? <CreateRestaurantButton /> : <></>}
      {restaurants.totalCount > 0 && (
        <Card>
          <CardContent>Your restaurants</CardContent>
          <DataGrid
            rows={restaurants.edges!.map((edge) => edge!.node!)}
            columns={[
              { field: "name", headerName: "Name", flex: 1 },
              {
                field: "address",
                headerName: "Address",
                flex: 1,
                sortable: false,
              },
              {
                field: "actions",
                headerName: "Actions",
                sortable: false,
                width: 200,
                filterable: false,
                hideable: false,
                renderCell: () => {
                  return (
                    <Stack
                      direction="row"
                      gap={2}
                      alignItems="center"
                      justifyContent="center"
                      height="100%"
                    >
                      <Button size="small" variant="text" startIcon={<Edit />}>
                        Edit
                      </Button>
                      <Button
                        size="small"
                        variant="text"
                        color="warning"
                        startIcon={<Delete />}
                      >
                        Delete
                      </Button>
                    </Stack>
                  );
                },
              },
            ]}
            sx={{
              border: "none",
              "& .MuiDataGrid-columnHeaders": {
                backgroundColor: "lightgray",
              },
              "& .MuiDataGrid-iconSeparator": {
                color: "transparent",
              },
              "& .MuiDataGrid-withBorderColor": {
                border: "none",
              },
            }}
            pageSizeOptions={[5, 10, 25]}
            paginationModel={{
              pageSize: 5,
              page: 0,
            }}
          />
        </Card>
      )}
    </Page>
  );
}

function CreateRestaurantButton() {
  return (
    <>
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
    </>
  );
}
