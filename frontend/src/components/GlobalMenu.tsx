import { gql, useQuery } from "@apollo/client";
import { ChevronLeft, Menu, Notifications } from "@mui/icons-material";
import { AppBar, Avatar, Badge, Box, Container, Divider, Drawer, Grid, IconButton, List, Paper, Toolbar, Typography } from "@mui/material";
import Copyright from "../components/Copyright";
import { mainListItems, secondaryListItems } from "../components/Menu/listitems";
import { useState } from "react";
import { Outlet } from "react-router-dom";

const USER_QUERY = gql`
query TestQuery {
    me {
        avatarURL
        authInfo {
            name
        }
    }
}
`

const GlobalMenu: React.FC = () => {
    const { loading, error, data } = useQuery(USER_QUERY);
    const [open, setOpen] = useState(false);

    if (loading) {
        return (
            <h1>Loading</h1>
        )
    }

    if (error) {
        return (
            <h1>{error.message}</h1>
        )
    }
    console.log(data)

    return (
        <Box sx={{ display: 'flex' }}>
            <AppBar position="absolute">
                <Toolbar
                    sx={{
                        pr: '24px', // keep right padding when drawer closed
                    }}
                >
                    <IconButton
                        edge="start"
                        color="inherit"
                        aria-label="open drawer"
                        onClick={() => setOpen(!open)}
                        sx={{
                            marginRight: '36px',
                            ...(open && { display: 'none' }),
                        }}
                    >
                        <Menu />
                    </IconButton>
                    <Typography
                        component="h1"
                        variant="h6"
                        color="inherit"
                        noWrap
                        sx={{ flexGrow: 1 }}
                    >
                        StudyGator
                    </Typography>
                    <IconButton color="inherit" sx={{ marginRight: "10px" }}>
                        <Badge badgeContent={4} color="warning" overlap="circular" anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}>
                            <Avatar alt={data.me.authInfo.name} src={data.me.avatarURL} sx={{ height: 48, width: 48}} />
                        </Badge>
                    </IconButton>
                </Toolbar>
            </AppBar>
            <Drawer variant="persistent" 
            open={open}
            >
                <Toolbar
                    sx={{
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'flex-end',
                        px: [1],
                    }}
                >
                    <IconButton 
                    onClick={() => setOpen(!open)}
                    >
                        <ChevronLeft />
                    </IconButton>
                </Toolbar>
                <Divider />
                <List component="nav">
                    {mainListItems}
                    <Divider sx={{ my: 1 }} />
                    {secondaryListItems}
                </List>
            </Drawer>
            <Box
                component="main"
                sx={{
                    backgroundColor: (theme) =>
                        theme.palette.mode === 'light'
                            ? theme.palette.grey[100]
                            : theme.palette.grey[900],
                    flexGrow: 1,
                    height: '100vh',
                    overflow: 'auto',
                }}
            >
                <Toolbar />
                <Container maxWidth="xl" sx={{ mt: 4, mb: 4 }}>
                    <Outlet />
                    <Copyright sx={{ pt: 4 }} />
                </Container>
            </Box>
        </Box>
    )
}

export default GlobalMenu;