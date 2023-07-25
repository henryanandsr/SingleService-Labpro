import { useAPI } from "@/contexts";
import { Button, Center, Heading, Input, VStack } from "@chakra-ui/react";
import { useRef } from "react";
import { useNavigate } from "react-router-dom";

const ConfigPage = () => {
  const { APIUrl, setAPIUrl } = useAPI();
  const navigate = useNavigate();
  const inputAPIRef = useRef<HTMLInputElement>(null);
  const handleSave = () => {
    setAPIUrl(inputAPIRef.current?.value || "");
    navigate("/");
  };
  return (
    <Center w="100vw" h="100vh">
      <VStack p={8} gap={6} bg="primary.100" borderRadius="xl">
        <Heading>Config</Heading>
        <Input placeholder="API Url" defaultValue={APIUrl || ""} ref={inputAPIRef} bg="white" />
        <Button onClick={handleSave}>Save</Button>
      </VStack>
    </Center>
  );
};

export default ConfigPage;
