import { Outlet } from "react-router-dom";
import SuspenceLoad from "../components/common/SuspenceLoad";
import Navbar from "../components/Navbar";
import { Suspense } from "react";

const Main = () => {
  return (
    <>
      <Navbar />
      <Suspense fallback={<SuspenceLoad/>}>
        <Outlet />
      </Suspense>
    </>
  );
};
export default Main;
