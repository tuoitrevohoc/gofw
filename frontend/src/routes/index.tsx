import { Routes, Route } from "react-router-dom";
import Home from "./Home";
import DefaultLayout, { MenuItem } from "./layouts/DefaultLayout";
import Analytics from "../assets/icons/navbar/ic-analytics.svg";
import User from "../assets/icons/navbar/ic-user.svg";
import Product from "../assets/icons/navbar/ic-cart.svg";
import Blog from "../assets/icons/navbar/ic-blog.svg";

const SidebarMenu: MenuItem[] = [
  { icon: Analytics, label: "Home", path: "/" },
  { icon: User, label: "Users", path: "/users" },
  { icon: Product, label: "Products", path: "/products" },
  { icon: Blog, label: "Blogs", path: "/blogs" },
];

export default function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<DefaultLayout menuItems={SidebarMenu} />}>
        <Route path="/*" element={<Home />} />
      </Route>
    </Routes>
  );
}
