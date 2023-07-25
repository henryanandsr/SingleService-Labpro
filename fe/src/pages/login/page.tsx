import { Button, Center, Heading, VStack, Input } from "@chakra-ui/react";
import { useAuth } from "@/contexts";
import { useState } from "react";
import { LoginResponseData } from "@/types";

// Define an interface for the login response data
interface LoginResponseData {
  user: {
    username: string;
    name: string;
  };
  token: string;
}

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

  const handleLogin = async () => {
    try {
      // Explicitly type the response variable
      const response: LoginResponseData = await login(formData.username, formData.password);
      console.log("Login response :", response);
      const token = response?.token;
      if (token) {
        console.log("Token stored:", window.localStorage.getItem("token"));
      } else {
        console.log("FKDJKFLDKL");
      }
    } catch (error) {
      console.error("An error occurred during login:", error);
    }
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
        <Button onClick={handleLogin}>Login</Button>
      </VStack>
    </Center>
  );
};

export default LoginPage;
