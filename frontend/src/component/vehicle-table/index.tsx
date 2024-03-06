import { ChangeEvent, useEffect, useState } from "react";
import Box from "@mui/material/Box";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TablePagination from "@mui/material/TablePagination";
import TableRow from "@mui/material/TableRow";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Paper from "@mui/material/Paper";
import IconButton from "@mui/material/IconButton";
import Tooltip from "@mui/material/Tooltip";
import FormControlLabel from "@mui/material/FormControlLabel";
import Switch from "@mui/material/Switch";
import { visuallyHidden } from "@mui/utils";
import AddFormModal from "../add-form-modal";
import { Button } from "@mui/material";

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

type VehiclePartTableData = {
  ID: number;
  Name: string;
  Description: string;
  Price: number;
  CreatedAt: string;
  UpdatedAt: string;
};

let fetchStatus: boolean = true;
const VehicleTable = () => {
  const [tableData, setTableData] = useState<VehiclePartTableData[]>([]);
  const [isModal, setIsModal] = useState<boolean>(false);
  const [edit, setEdit] = useState({ id: 0, status: false });
  const [isFetch, setIsFetch] = useState(false);
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [totalSize, setTotalSize] = useState(0);

  useEffect(() => {
    if (fetchStatus) {
      const fetchVehicleRecord = async () => {
        try {
          const raw = await fetch("http://localhost:8080/vehicle-part/list", {
            method: "POST",
            headers: {
              Accept: "application.json",
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              pageNo: page + 1,
              perPage: 10,
            }),
          });
          const response = await raw.json();
          setTableData(response.recordList);
          setTotalSize(response.totalCount);
        } catch (error) {
          console.log(error);
        }
      };
      fetchVehicleRecord();
      fetchStatus = false;
    }
  }, [isFetch]);

  const handleEdit = async (values: any) => {
    try {
      const raw = await fetch(`http://localhost:8080/vehicle-part/${edit.id}`, {
        method: "PATCH",
        headers: {
          Accept: "application.json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: values.Name,
          description: values.Description,
          price: parseInt(values.Price as unknown as string),
        }),
      });
      const response = await raw.json();

      if (response.status) {
        fetchStatus = true;
        setIsFetch(!isFetch);
        handleToggle();
      }
    } catch (error: any) {
      console.log(error);
    }
  };

  const isFetchToggle = () => {
    fetchStatus = true;
  };

  const handleDelete = async (id: number) => {
    try {
      const raw = await fetch(`http://localhost:8080/vehicle-part/${id}`, {
        method: "DELETE",
      });
      const response = await raw.json();
      if (response.status) {
        fetchStatus = true;
        setIsFetch(!isFetch);
      }
    } catch (error: any) {
      console.log(error);
    }
  };

  const handleToggle = () => {
    setIsModal(!isModal);
    if (edit.status) {
      setEdit({
        id: 0,
        status: false,
      });
    }
  };

  const handleDownload = async () => {
    try {
      const response = await fetch("http://localhost:8080/generate-pdf", {
        method: "POST",
        headers: {
          Accept: "application.json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify(tableData),
      });
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

  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage);
    fetchStatus = true;
    setIsFetch(!isFetch);
  };

  const handleChangeRowsPerPage = (evt: ChangeEvent<HTMLInputElement>) => {
    setRowsPerPage(parseInt(evt.target.value));
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
              <Button variant="contained" onClick={handleToggle}>
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
              <TableCell sx={tableHeadStyle}>Name</TableCell>
              <TableCell sx={tableHeadStyle}>Description</TableCell>
              <TableCell sx={tableHeadStyle}>Price</TableCell>
              <TableCell sx={tableHeadStyle}>Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {tableData?.map((row: VehiclePartTableData) => (
              <TableRow
                key={row.ID}
                sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
              >
                <TableCell sx={tableCell}>{row.Name}</TableCell>
                <TableCell sx={tableCell}>{row.Description}</TableCell>
                <TableCell sx={tableCell}>{row.Price}</TableCell>
                <TableCell sx={tableHeadStyle}>
                  <Button
                    sx={{ fontWeight: "600" }}
                    variant="contained"
                    color="warning"
                    onClick={() => {
                      setEdit({
                        id: row.ID,
                        status: true,
                      });
                      handleToggle();
                    }}
                  >
                    Edit
                  </Button>
                  <Button
                    sx={styleDeleteAction}
                    variant="contained"
                    color="error"
                    onClick={() => handleDelete(row.ID)}
                  >
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <TablePagination
        rowsPerPageOptions={[10]}
        component="div"
        count={totalSize}
        rowsPerPage={rowsPerPage}
        page={page}
        onPageChange={handleChangePage}
        onRowsPerPageChange={handleChangeRowsPerPage}
      />
      <AddFormModal
        handleToggle={handleToggle}
        handleEdit={handleEdit}
        isModal={isModal}
        isEdit={edit.status}
        setIsFetch={setIsFetch}
        isFetchToggle={isFetchToggle}
      />
    </Box>
  );
};

export default VehicleTable;
