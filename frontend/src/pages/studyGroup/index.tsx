import React, { useState } from "react";
import { useQuery, gql, useMutation } from "@apollo/client";
import {
  List,
  ListItem,
  ListItemText,
  Typography,
  Button,
  Paper,
  Modal,
  TextField,
  MenuItem,
  FormControl,
  InputLabel,
  Select,
  SelectChangeEvent,
  useTheme,
} from "@mui/material";
import AddIcon from "@mui/icons-material/Add";

const CLASSES_QUERY = gql`
  query ClassesQuery {
    classes {
      id
      classInfo {
        name
      }
    }
  }
`;

const STUDY_GROUP_MUTATION = gql`
mutation StudyGroupMutation($classID: ID!, $name: String!, $description: String!) {
    createStudyGroup(classID: $classID, name: $name, description: $description) {
        id
    }
}
`

const STUDY_GROUPS_QUERY = gql`
  query StudyGroupsQuery {
    studyGroups(onlyFavorites: false) {
      id
      owner {
        id
        authInfo {
          name
        }
      }
      name
      description
      favorite
    }
  }
`;

const StudyGroups: React.FC = () => {
  const theme = useTheme();
  const { loading, error, data, refetch } = useQuery(STUDY_GROUPS_QUERY);
  const [createStudyGroup] = useMutation(STUDY_GROUP_MUTATION)
  const classes = useQuery(CLASSES_QUERY);
  const [open, setOpen] = useState(false);
  const [classValue, setClassValue] = useState("");
  const [description, setDescription] = useState("");
  const [name, setName] = useState("");

  if (loading || classes.loading) return <p>Loading...</p>;
  if (error || classes.error) return <p>Error :{JSON.stringify(error)}</p>;

  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleClassChange = (event: SelectChangeEvent<string>) => {
    setClassValue(event.target.value);
  };

  const handleDescriptionChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setDescription(event.target.value);
  };

  const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setName(event.target.value);
  };

  const handleSubmit = () => {
    // Handle form submission logic here
    createStudyGroup({
        variables: {
            classID: classValue,
            name: name,
            description: description
        }
    }).then(() => {
        handleClose();
        refetch();
    })
  };

  return (
    <Paper elevation={3}>
        <Typography variant="h3">Study Groups</Typography>
      <Button
        variant="contained"
        color="primary"
        style={{ margin: 15 }}
        startIcon={<AddIcon />}
        onClick={handleOpen}
      >
        New Study Group
      </Button>
      <List>
        {data.studyGroups.map(
          (studyGroup: {
            id: string;
            owner: { id: string; authInfo: { name: string } };
            name: string;
            description: string;
            favorite: boolean;
          }) => (
            <ListItem key={studyGroup.id}>
              <ListItemText
                primary={studyGroup.name}
                secondary={
                  <>
                    <Typography variant="body2">
                      {`Owner: ${studyGroup.owner.authInfo.name}`}
                    </Typography>
                    <Typography variant="body2">
                      {studyGroup.description}
                    </Typography>
                  </>
                }
              />
              <Button
                variant="outlined"
                color={studyGroup.favorite ? "primary" : "secondary"}
              >
                {studyGroup.favorite ? "Favorited" : "Favorite"}
              </Button>
            </ListItem>
          )
        )}
      </List>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
      >
        <Paper style={{ margin: 100, padding: 20 }}>
          <Typography variant="h6" id="modal-modal-title" gutterBottom>
            Create a New Study Group
          </Typography>
          <FormControl
            variant="outlined"
            style={{ width: "100%", marginBottom: theme.spacing(2) }}
          >
            <InputLabel id="class-select-label">Select Class</InputLabel>
            <Select
              labelId="class-select-label"
              id="class-select"
              value={classValue}
              onChange={handleClassChange}
              label="Select Class"
            >
              {classes.data.classes.map(
                (dat: { id: string; classInfo: { name: string } }) => (
                  <MenuItem value={dat.id}>{dat.classInfo.name}</MenuItem>
                )
              )}
              {/* Add more options as needed */}
            </Select>
          </FormControl>
          <TextField
            id="name"
            label="Name"
            variant="outlined"
            fullWidth
            value={name}
            onChange={handleNameChange}
          />
          <TextField
            id="description"
            label="Description"
            variant="outlined"
            fullWidth
            value={description}
            onChange={handleDescriptionChange}
            style={{ marginTop: 15 }}
          />
          <Button
            variant="contained"
            color="primary"
            onClick={handleSubmit}
            style={{ margin: 5 }}
          >
            Submit
          </Button>
        </Paper>
      </Modal>
    </Paper>
  );
};

export default StudyGroups;
