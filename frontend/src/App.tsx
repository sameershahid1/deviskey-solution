import Box from "@mui/material/Box";
import Header from "./component/header";
import Footer from "./component/footer";
import VehicleTable from "./component/vehicle-table";

const layout = {
  display: "flex",
  flexDirection: "column",
  justifyContent: "center",
  alignItems: "space-between",
  gap: "6.1rem",
};

function App() {
  return (
    <Box sx={layout}>
      <Header />
      <VehicleTable />
      <Footer />
    </Box>
  );
}

export default App;
