import { 
  Button,
  TextField,
  Box,
  Stack,
 } from "@mui/material";
export function Welcome() {
  return (
    <main className="flex items-center justify-center pt-16 pb-4">
      <Stack component="section" sx={{ p: 2, border: '1px solid grey' }}>
        <TextField id="standard-basic" label="Username/Email" variant="standard" />
        <TextField id="standard-basic" label="Password" variant="standard" sx={{mb: "30px"}}/>
        <Button variant="outlined">Login</Button>
      </Stack>
    </main>
  );
}
