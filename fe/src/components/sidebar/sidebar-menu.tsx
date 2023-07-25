import { Center, Heading } from "@chakra-ui/react";

interface SidebarMenuProps {
  currentMenu: string;
  menu: string;
  onClick: () => void;
}

const SidebarMenu = ({ currentMenu, menu, onClick }: SidebarMenuProps) => {
  return (
    <Center
      w="100%"
      h={24}
      bg={currentMenu === menu.toLowerCase() ? "primary.400" : "primary.500"}
      _hover={{ bg: "primary.400", cursor: "pointer" }}
      onClick={onClick}
    >
      <Heading size="lg">{menu}</Heading>
    </Center>
  );
};

export default SidebarMenu;
