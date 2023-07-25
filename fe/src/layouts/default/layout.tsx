import Sidebar from "@/components/sidebar/component";
import { Flex } from "@chakra-ui/react";
import { Outlet } from "react-router-dom";

const DefaultLayout = () => {
  return (
    <Flex w="100vw" h="100vh" flexDir="row">
      <Sidebar />
      <Outlet />
    </Flex>
  );
};

export default DefaultLayout;
