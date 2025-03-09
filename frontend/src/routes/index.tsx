import { Routes, Route, Navigate } from "react-router-dom";
import Home from "./Home";
import DefaultLayout, { MenuItem } from "./DefaultLayout";
import Analytics from "../assets/icons/navbar/ic-analytics.svg";
import User from "../assets/icons/navbar/ic-user.svg";
import Product from "../assets/icons/navbar/ic-cart.svg";
import Blog from "../assets/icons/navbar/ic-blog.svg";
import AuthLayout from "./auth/AuthLayout";
import RegisterPage from "./auth/RegisterPage";
import { graphql, useLazyLoadQuery } from "react-relay";
import { routesQuery } from "./__generated__/routesQuery.graphql";
import LoginPage from "./auth/LoginPage";

const SidebarMenu: MenuItem[] = [
  { icon: Analytics, label: "Home", path: "/" },
  { icon: User, label: "Users", path: "/users" },
  { icon: Product, label: "Products", path: "/products" },
  { icon: Blog, label: "Blogs", path: "/blogs" },
];

export default function AppRoutes() {
  const { viewer } = useLazyLoadQuery<routesQuery>(
    graphql`
      query routesQuery {
        viewer {
          userId
          isAuthenticated
        }
      }
    `,
    {}
  );

  if (viewer?.isAuthenticated) {
    return (
      <Routes>
        <Route path="/" element={<DefaultLayout menuItems={SidebarMenu} />}>
          <Route path="/*" element={<Home />} />
        </Route>
      </Routes>
    );
  }

  return (
    <Routes>
      <Route path="/auth" element={<AuthLayout />}>
        <Route path="/auth/login" element={<LoginPage />} />
        <Route path="/auth/register" element={<RegisterPage />} />
        <Route path="*" element={<Navigate to="/auth/register" replace />} />
      </Route>
      <Route path="*" element={<Navigate to="/auth/register" replace />} />
    </Routes>
  );
}
