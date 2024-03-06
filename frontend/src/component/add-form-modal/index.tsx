import { Box, Modal, Grid, Button, FormLabel, TextField } from "@mui/material";
import { useState } from "react";
import { modalFields } from "../../data/modal_field";

const style = {
  position: "absolute" as "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 800,
  bgcolor: "background.paper",
  border: "2px solid #000",
  boxShadow: 24,
  p: 4,
};

const handleModalFieldOnChange = (
  data: any,
  newValue: string,
  fieldAlias: string,
  setValues: any
) => {
  setValues({
    ...data,
    [fieldAlias]: newValue,
  });
};

type ValueType = {
  Name: string;
  Description: string;
  Price: number;
};

type Props = {
  handleToggle: any;
  isModal: any;
  handleEdit: any;
  isEdit: boolean;
  setIsFetch: any;
  isFetchToggle: any;
};

const AddFormModal = ({
  handleToggle,
  isModal,
  handleEdit,
  isEdit,
  setIsFetch,
  isFetchToggle,
}: Props) => {
  const [values, setValues] = useState<ValueType>({
    Name: "",
    Description: "",
    Price: 0,
  });

  const handleSubmit = async (evt: any) => {
    evt.preventDefault();
    try {
      const raw = await fetch("http://localhost:8080/vehicle-part", {
        method: "POST",
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
        isFetchToggle();
        setIsFetch((prev: boolean) => !prev);
        handleToggle();
      }
    } catch (error: any) {
      console.log(error);
    }
  };

  return (
    <Modal
      open={isModal}
      onClose={handleToggle}
      aria-labelledby="modal-modal-title"
      aria-describedby="modal-modal-description"
    >
      <Box sx={style}>
        <h1 id="parent-modal-title">Add Vehicle Part</h1>
        <form
          id="parent-modal-description"
          onSubmit={
            isEdit
              ? (evt: any) => {
                  evt.preventDefault();
                  handleEdit(values);
                }
              : handleSubmit
          }
        >
          <Grid spacing={3} container>
            {modalFields?.map((field: any) => {
              return (
                <Grid key={field.id} item xs={12} lg={field.column}>
                  <FormLabel>{field.title}</FormLabel>
                  <TextField
                    id="name"
                    size="small"
                    variant="outlined"
                    fullWidth
                    value={values[field.alias as never]}
                    placeholder={field.placeholder ?? ""}
                    rows={field.rows ?? 1}
                    multiline={field.multiline ?? false}
                    type={field.type}
                    onChange={(evt: any) => {
                      handleModalFieldOnChange(
                        values,
                        evt.target.value,
                        field.alias,
                        setValues
                      );
                    }}
                  />
                </Grid>
              );
            })}
            <Grid item xs={12} lg={12}>
              <Button
                variant="contained"
                color="primary"
                sx={{ mr: 1 }}
                type="submit"
              >
                Submit
              </Button>
              <Button variant="contained" color="error" onClick={handleToggle}>
                Cancel
              </Button>
            </Grid>
          </Grid>
        </form>
      </Box>
    </Modal>
  );
};

export default AddFormModal;
