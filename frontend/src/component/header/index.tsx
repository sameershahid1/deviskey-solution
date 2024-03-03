import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";

const Header = () => {
  return (
    <Box sx={{ margin: "0" }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div">
            Deviskey Solution
          </Typography>
        </Toolbar>
      </AppBar>
    </Box>
  );
};

export default Header;
