import { Routes, Route, Navigate } from "react-router-dom";
import { lazy } from "react";
import Home from "./Home";
import DefaultLayout, { MenuItem } from "./DefaultLayout";
import { routesQuery } from "./__generated__/routesQuery.graphql";

import Analytics from "../assets/icons/navbar/ic-analytics.svg";
import User from "../assets/icons/navbar/ic-user.svg";
import Product from "../assets/icons/navbar/ic-cart.svg";
import Blog from "../assets/icons/navbar/ic-blog.svg";
import AuthLayout from "./auth/AuthLayout";
import RegisterPage from "./auth/RegisterPage";
import { graphql, useLazyLoadQuery } from "react-relay";

const LoginPage = lazy(() => import("./auth/LoginPage"));
const ErrorPage = lazy(() => import("./error/ErrorPage"));
const RestaurantList = lazy(() => import("./restaurants/RestaurantList"));
const NewRestaurant = lazy(() => import("./restaurants/NewRestaurant"));

const SidebarMenu: MenuItem[] = [
  { icon: Analytics, label: "Home", path: "/" },
  { icon: Product, label: "Restaurants", path: "/restaurants" },
  { icon: Blog, label: "Menu", path: "/menu" },
  { icon: User, label: "Customers", path: "/customers" },
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
          <Route path="/" element={<Home />} />
          <Route path="/restaurants" element={<RestaurantList />} />
          <Route path="/restaurants/add" element={<NewRestaurant />} />
        </Route>
        <Route
          path="/*"
          element={
            <ErrorPage message="I could not find the page you are looking for. Forgot to plug in the router?" />
          }
        />
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
