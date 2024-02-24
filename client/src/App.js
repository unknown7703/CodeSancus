import "./App.css";
import { Suspense, lazy } from "react";
import {
  Navigate,
  Route,
  createBrowserRouter,
  createRoutesFromElements,
  RouterProvider,
} from "react-router-dom";
import { routes } from "./routes/routes";
import SuspenceLoad from "./components/common/SuspenceLoad";
const ErrorComponent = lazy(() => import("./components/common/ErrorComponent"));
const router = createBrowserRouter(
  createRoutesFromElements(
    <Route>
      <Route
        path={routes.main.path}
        element={<Navigate to={routes.upload.path} />}
      ></Route>
      <Route path={routes.main.path} element={<routes.main.element />}>
        <Route
          path={routes.upload.path}
          element={<routes.upload.element />}
          errorElement={<ErrorComponent />}
        ></Route>
        <Route
          path={routes.result.path}
          element={<routes.result.element />}
          errorElement={<ErrorComponent />}
        ></Route>
      </Route>
      <Route path={routes.invalid.path}></Route>
    </Route>
  )
);
function App() {
  return (
    <Suspense fallback={<SuspenceLoad />}>
      <RouterProvider router={router} />
    </Suspense>
  );
}

export default App;
