import { Divider, Heading, Spacer, VStack } from "@chakra-ui/react";
import { useLocation, useNavigate } from "react-router-dom";
import SidebarMenu from "./sidebar-menu";
import { useAuth } from "@/contexts";

const Sidebar = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const { user, logout } = useAuth();

  const currentMenu = location.pathname.split("?")[0].includes("barang") ? "barang" : "perusahaan";

  return (
    <VStack w="250px" gap={6} py={4} h="100vh" bg="primary.500" color="white" textAlign="center">
      <Heading>{user.username}</Heading>
      <Divider mt={8} />
      <SidebarMenu currentMenu={currentMenu} menu="Barang" onClick={() => navigate("/barang")} />
      <SidebarMenu currentMenu={currentMenu} menu="Perusahaan" onClick={() => navigate("/perusahaan")} />
      <Spacer />
      <Divider mt={8} />
      <SidebarMenu currentMenu={currentMenu} menu="Logout" onClick={logout} />
    </VStack>
  );
};

export default Sidebar;
