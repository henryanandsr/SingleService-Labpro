import { Barang, Config, Login, Perusahaan } from "@/pages";
import { Outlet, createBrowserRouter } from "react-router-dom";
import { DefaultLayout } from "./layouts";
import { APIProvider } from "./providers";
import { Text } from "@chakra-ui/react";
import AuthProvider from "./providers/AuthProvider";

const Providers = () => {
  return (
    <APIProvider>
      <AuthProvider>
        <Outlet />
        {/* Karena males bikin layer lagi jd disini aja 👍 */}
        <Text position="absolute" right={4} bottom={4}>
          Made with love by Labpro 2020 💕
        </Text>
      </AuthProvider>
    </APIProvider>
  );
};

const Router = createBrowserRouter([
  {
    path: "/",
    element: <Providers />, // Tbh, I don't know how to do it better (cus this is new). deal with it 😎
    children: [
      {
        path: "/",
        element: <DefaultLayout />,
        children: [
          {
            path: "/barang",
            element: <Barang />,
          },
          {
            path: "/perusahaan",
            element: <Perusahaan />,
          },
        ],
      },
      {
        path: "/config",
        element: <Config />,
      },
      {
        path: "/login",
        element: <Login />,
      },
    ],
  },
]);

export default Router;
