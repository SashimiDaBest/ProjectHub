import { useState } from "react";
import { 
  Button,
  TextField,
  Box,
  Stack,
 } from "@mui/material";
import { useNavigate } from "react-router";


export function Welcome() {
  const navigate = useNavigate();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  const handleLogin = async () => {
    try {
      const response = await fetch("http://localhost:8000/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      });

      const data = await response.json();

      if (response.ok) {
        // Successfully logged in, data.token contains JWT
        setMessage("Login successful!");
        localStorage.setItem("token", data.token);
        navigate("/dashboard");
      } else {
        setMessage(data || "Login failed");
      }
    } catch (err) {
      console.error(err);
      setMessage("Error connecting to server");
    }
  };

  return (
    <main className="flex items-center justify-center pt-16 pb-4">
      <Stack component="section" sx={{ p: 2, border: '1px solid grey' }}>
        <TextField 
          id="standard-basic" 
          label="Username/Email" 
          variant="standard" 
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <TextField 
          id="standard-basic" 
          label="Password" 
          variant="standard" 
          sx={{mb: "30px"}}
          value={password}
          onChange={(e) => setPassword(e.target.value)}/>
        <Button 
          variant="outlined"
          onClick={handleLogin}
        >
          Login
        </Button>
        {message && <Box sx={{ mt: 2 }}>{message}</Box>}
      </Stack>
    </main>
  );
}
