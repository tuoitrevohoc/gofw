import { StrictMode, Suspense } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import { ThemeProvider } from "./theme/theme-provider.tsx";
import AppRoutes from "./routes/index.tsx";
import GlobalStyle from "@mui/material/GlobalStyles";
import { baseVars } from "./theme/base-vars.ts";
import { RelayEnvironmentProvider } from "react-relay";
import environment from "./relay/environment.ts";

import "./global.css";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <RelayEnvironmentProvider environment={environment}>
      <ThemeProvider>
        <GlobalStyle styles={(theme) => ({ body: baseVars(theme) })} />
        <BrowserRouter>
          <Suspense fallback={<div>Loading...</div>}>
            <AppRoutes />
          </Suspense>
        </BrowserRouter>
      </ThemeProvider>
    </RelayEnvironmentProvider>
  </StrictMode>
);
