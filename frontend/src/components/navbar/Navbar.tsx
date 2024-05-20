import { AppBar, Box, Container, Toolbar, useTheme } from "@mui/material";
import { NAVBAR_HEIGHT_PX } from "../../utils/constants";
import Logo from "../Logo";

export default function Navbar() {
  const theme = useTheme();
  const grey = theme.palette.grey[700]

  return (
    <AppBar position="static" sx={{ bgcolor: grey, minHeight: NAVBAR_HEIGHT_PX }}>
      <Container maxWidth="xl">
        <Toolbar disableGutters sx={{ minHeight: NAVBAR_HEIGHT_PX }}>
          <Logo />
          <Box sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  )
}
