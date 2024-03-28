import * as React from 'react';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';
import GroupIcon from '@mui/icons-material/Group';
import DashboardIcon from '@mui/icons-material/Dashboard';
import ClassIcon from '@mui/icons-material/School';



export const mainListItems = (
  <React.Fragment>
    <ListItemButton>
      <ListItemIcon>
        <DashboardIcon />
      </ListItemIcon>
      <ListItemText primary="Dashboard" />
    </ListItemButton>
    <ListItemButton>
      <ListItemIcon>
        <ClassIcon />
      </ListItemIcon>
      <ListItemText primary="Classes" />
    </ListItemButton>
    <ListItemButton>
      <ListItemIcon>
        <GroupIcon />
      </ListItemIcon>
      <ListItemText primary="My Groups" />
    </ListItemButton>
  </React.Fragment>
);

export const secondaryListItems = (
  <React.Fragment>
    <ListSubheader component="div" inset>
      Favorite Groups
    </ListSubheader>
    <ListItemButton>
      <ListItemIcon>
        <GroupIcon />
      </ListItemIcon>
      <ListItemText primary="DSA Study Group" />
    </ListItemButton>
    <ListItemButton>
      <ListItemIcon>
        <GroupIcon />
      </ListItemIcon>
      <ListItemText primary="Quest Class Study Group" />
    </ListItemButton>
  </React.Fragment>
);