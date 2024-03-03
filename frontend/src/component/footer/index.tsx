import { Box, Link, Typography } from "@mui/material";

const Footer = () => {
  return (
    <footer>
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          padding: "1rem",
          backgroundColor: "#101010b6",
          marginTop: "4rem",
          color: "white",
        }}
      >
        <Typography>
          Copyright © Deviskey Solution. All Rights Reserved. Developed by
          <Link
            href={"https://deviskey.com/"}
            sx={{ fontWeight: "600", marginLeft: "4px" }}
          >
            Deviskey
          </Link>
        </Typography>
      </Box>
    </footer>
  );
};

export default Footer;
