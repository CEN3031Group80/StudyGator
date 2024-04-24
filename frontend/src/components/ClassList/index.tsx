import React from 'react';
import {List, ListItem, ListItemText, Typography } from '@mui/material';

interface Class {
  name: string;
  university_name: string;
  description: string;
}

interface ClassListProps {
  classes: Class[];
}

const ClassList: React.FC<ClassListProps> = ({ classes }) => {
  
    return (
      <List>
        {classes.map((classItem, index) => (
          <ListItem key={index}>
            <ListItemText
              primary={classItem.name}
              secondary={
                <>
                  <Typography variant="body2" style={{fontSize: 16}}>
                    {classItem.university_name}
                  </Typography>
                  <Typography variant="body2">
                    {classItem.description}
                  </Typography>
                </>
              }
            />
          </ListItem>
        ))}
      </List>
    );
  };
  
  export default ClassList;