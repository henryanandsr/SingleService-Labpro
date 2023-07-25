import { Button, Center, Heading, VStack, Input } from "@chakra-ui/react";
import { useAuth } from "@/contexts";
import { useState } from "react";

const LoginPage = () => {
  const { login } = useAuth();
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });

  const handleChangeFormData = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  return (
    <Center w="100%" h="100vh">
      <VStack px={12} py={6} gap={8} borderRadius="xl" bg="primary.100">
        <Heading>Login</Heading>
        <Input
          placeholder="Username"
          value={formData.username}
          name="username"
          onChange={handleChangeFormData}
          bg="white"
        />
        <Input
          type="password"
          placeholder="Password"
          value={formData.password}
          name="password"
          onChange={handleChangeFormData}
          bg="white"
        />
        <Button onClick={() => login(formData.username, formData.password)}>Login</Button>
      </VStack>
    </Center>
  );
};

export default LoginPage;
