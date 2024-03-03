import { useEffect, useState } from "react";
import Box from "@mui/material/Box";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Paper from "@mui/material/Paper";
import IconButton from "@mui/material/IconButton";
import Tooltip from "@mui/material/Tooltip";
import { Button } from "@mui/material";

function createData(
  name: string,
  calories: number,
  fat: number,
  carbs: number,
  protein: number
) {
  return { name, calories, fat, carbs, protein };
}

const rows = [
  createData("Frozen yoghurt", 159, 6.0, 24, 4.0),
  createData("Ice cream sandwich", 237, 9.0, 37, 4.3),
  createData("Eclair", 262, 16.0, 24, 6.0),
  createData("Cupcake", 305, 3.7, 67, 4.3),
  createData("Gingerbread", 356, 16.0, 49, 3.9),
];

const tableHeadStyle = {
  textAlign: "center",
  fontWeight: "600",
  fontSize: "1rem",
};

const tableCell = {
  textAlign: "center",
  fontSize: "1rem",
};

const styleDeleteAction = {
  marginLeft: "5px",
  fontWeight: "600",
};

let fetchStatus: boolean = true;

const VehicleTable = () => {
  useEffect(() => {
    if (fetchStatus) {
      const fetchVehicleRecord = async () => {
        const raw = await fetch("http://localhost:8080/vehicle-part/list");
        const response = await raw.json();
        console.log(response);
      };
      fetchVehicleRecord();
      fetchStatus = false;
    }
  }, []);

  const handleDownload = async () => {
    try {
      const response = await fetch("http://localhost:8080/generate-pdf"); // Adjust the endpoint path
      response.blob().then((blob) => {
        let url = window.URL.createObjectURL(blob);
        let a = document.createElement("a");
        a.href = url;
        a.download = "employees.pdf";
        a.click();
      });

      if (!response.ok) {
        throw new Error(`Failed to download PDF: ${response.statusText}`);
      }
    } catch (error: any) {
      console.log(error.message);
    }
  };

  return (
    <Box
      sx={{
        boxShadow: "0 -1px 2px 1px rgba(0,0,0,0.1)",
        borderRadius: "10px",
        marginLeft: "auto",
        marginRight: "auto",
        paddingTop: "1rem",
        width: "80vw",
      }}
    >
      <Toolbar
        sx={{
          display: "flex",
          flexDirection: "row",
          justifyContent: "space-between",
          alignItems: "center",
          flexFlow: "wrap",
          marginLeft: "1rem",
          marginRight: "1rem",
        }}
      >
        <Typography variant="h4" fontWeight={700} id="tableTitle">
          Vehicle Parts
        </Typography>
        <Box>
          <Tooltip title="Add Part">
            <IconButton>
              <Button variant="contained" onClick={handleDownload}>
                Add Part
              </Button>
            </IconButton>
          </Tooltip>
          <Tooltip title="PDF Download">
            <IconButton>
              <Button variant="contained" onClick={handleDownload}>
                PDF Download
              </Button>
            </IconButton>
          </Tooltip>
        </Box>
      </Toolbar>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell sx={tableHeadStyle}>Dessert</TableCell>
              <TableCell sx={tableHeadStyle}>Calories</TableCell>
              <TableCell sx={tableHeadStyle}>Fat&nbsp;(g)</TableCell>
              <TableCell sx={tableHeadStyle}>Carbs&nbsp;(g)</TableCell>
              <TableCell sx={tableHeadStyle}>Protein&nbsp;(g)</TableCell>
              <TableCell sx={tableHeadStyle}>Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <TableRow
                key={row.name}
                sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
              >
                <TableCell sx={tableCell}>{row.name}</TableCell>
                <TableCell sx={tableCell}>{row.calories}</TableCell>
                <TableCell sx={tableCell}>{row.fat}</TableCell>
                <TableCell sx={tableCell}>{row.carbs}</TableCell>
                <TableCell sx={tableCell}>{row.protein}</TableCell>
                <TableCell sx={tableHeadStyle}>
                  <Button
                    sx={{ fontWeight: "600" }}
                    variant="contained"
                    color="warning"
                  >
                    Edit
                  </Button>
                  <Button
                    sx={styleDeleteAction}
                    variant="contained"
                    color="error"
                  >
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
};

export default VehicleTable;
