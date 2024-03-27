import { Navigate, Outlet } from 'react-router-dom'
import { apiUrl } from '../main';
import { Alert, Container, LinearProgress, Stack, Typography } from '@mui/material';
import Copyright from './Copyright';
import { ReportProblem } from '@mui/icons-material';
import { useEffect, useState } from 'react';
export default function SecuredRoute() {
    let [{ loading, error, data}, setStatus] = useState({ loading: true, error: "", data: false})

    useEffect(() => {
        async function grabStatus() {
            fetch(apiUrl + "/auth/status", { credentials: "include" }).then(async res => {
                let body = await res.json();
                
                if ("status" in body && body.status == "logged in") {
                    setStatus({ loading: false, error: "", data: true});
                } else {
                    setStatus({ loading: false, error: "", data: false});
                }
            }).catch(_ => {
                setStatus({ loading: false, error: "failed to pull API data", data: false});
            });
        }

        grabStatus();
    }, []);

    if (loading || data == undefined) {
        return (
            <Container component="main" maxWidth="xs">
                <Stack
                sx={{
                    marginTop: 50,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
                >
                
                </Stack>
                <LinearProgress color='secondary' />
                <Copyright sx={{ mt: 8, mb: 4 }} />
            </Container>
        )
    }

    if (error) {
        return (
            <Container component="main" maxWidth="xs">
                <Stack
                sx={{
                    marginTop: 50,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
                >
                </Stack>
                <Alert icon={<ReportProblem fontSize="inherit" />} severity="error">
                    Could not communicate with StudyGator API: {error}
                </Alert>
                <Copyright sx={{ mt: 8, mb: 4 }} />
            </Container>
        )
    }

    return (
        data ? <Outlet /> : <Navigate to='/login' />
    )
};
