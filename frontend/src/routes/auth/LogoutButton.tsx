import { useNavigate } from "react-router-dom";
import { useMutation, graphql } from "react-relay";
import { useCallback } from "react";
import { Button } from "@mui/material";
import { LogoutButtonMutation } from "./__generated__/LogoutButtonMutation.graphql";
import { handleLogout } from "../../auth";

export default function LogoutButton() {
  const [commitLogout, isLoggingOut] = useMutation<LogoutButtonMutation>(
    graphql`
      mutation LogoutButtonMutation {
        signOut
      }
    `
  );
  const navigate = useNavigate();

  const onLogoutClicked = useCallback(() => {
    handleLogout();

    commitLogout({
      variables: {},
      onCompleted: (data) => {
        if (data.signOut) {
          navigate("/login");
        }
      },
    });
  }, [commitLogout, navigate]);

  return (
    <Button
      fullWidth
      color="error"
      onClick={onLogoutClicked}
      loading={isLoggingOut}
    >
      Logout
    </Button>
  );
}
