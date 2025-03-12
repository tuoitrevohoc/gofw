import { Stack, Typography, Breadcrumbs, Link } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";

interface PageProps {
  title: string;
  breadcrumbs: {
    label: string;
    path: string;
  }[];
  children?: React.ReactNode;
}

export default function Page({ title, breadcrumbs, children }: PageProps) {
  return (
    <Stack gap={2}>
      <Stack>
        <Typography variant="h6">{title}</Typography>
      </Stack>
      <Breadcrumbs separator="›">
        {breadcrumbs.map((breadcrumb) => (
          <Link
            key={breadcrumb.path}
            component={RouterLink}
            to={breadcrumb.path}
          >
            {breadcrumb.label}
          </Link>
        ))}
      </Breadcrumbs>
      {children && <Stack gap={2}>{children}</Stack>}
    </Stack>
  );
}
