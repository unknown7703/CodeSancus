import { lazy } from "react";
const Main = lazy(() => import("../pages/Main"));
const Upload = lazy(() => import("../components/Upload"));
const Result = lazy(() => import("../components/Result"));
const routes = {
  main: {
    path: "/",
    element: Main,
  },
  upload: {
    path: "/upload",
    element: Upload,
  },
  result: {
    path: "/result",
    element: Result,
  },
  invalid: {
    path: "/*",
    element: Upload,
  },
};
export { routes };
