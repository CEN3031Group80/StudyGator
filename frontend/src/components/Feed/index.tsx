import React from 'react';
import { Avatar, Card, CardContent, CardHeader, Divider, Grid, List, ListItem, ListItemIcon, ListItemText, Typography } from '@mui/material';
import Markdown from 'markdown-to-jsx';

// Sample data for the social media feed
const feedData = [
  {
    id: 1,
    poster: {
      name: 'John Doe',
      avatarUrl: 'https://via.placeholder.com/150',
    },
    content: `
      ## Hello World
      This is a **markdown** example.

      ![Sample Image](https://via.placeholder.com/200)
    `,
    attachments: [
      { name: 'Document.pdf', type: 'pdf' },
      { name: 'Spreadsheet.xlsx', type: 'excel' },
    ],
  },
  // Add more feed items here...
];

const Feed: React.FC = () => {

  return (
    <List>
      {feedData.map((item) => (
        <Grid item xs={12} key={item.id}>
          <Card>
            <CardHeader
              avatar={
                <Avatar aria-label="avatar" src={item.poster.avatarUrl} />
              }
              title={item.poster.name}
              subheader="Posted on: January 1, 2023"
            />
            <CardContent>
              <Markdown>{item.content}</Markdown>
              <Divider />
              <Typography variant="subtitle2" color="textSecondary" gutterBottom>
                Attachments:
              </Typography>
              <List>
                {item.attachments.map((attachment, index) => (
                  <ListItem key={index}>
                    <ListItemIcon>
                      {/* Display attachment icon based on file type */}
                      {attachment.type === 'pdf' ? (
                        <img src="/pdf-icon.png" alt="pdf icon" style={{ width: 24, height: 24 }} />
                      ) : attachment.type === 'excel' ? (
                        <img src="/excel-icon.png" alt="excel icon" style={{ width: 24, height: 24 }} />
                      ) : (
                        <img src="/default-icon.png" alt="default icon" style={{ width: 24, height: 24 }} />
                      )}
                    </ListItemIcon>
                    <ListItemText primary={attachment.name} />
                  </ListItem>
                ))}
              </List>
            </CardContent>
          </Card>
        </Grid>
      ))}
    </List>
  );
};

export default Feed;