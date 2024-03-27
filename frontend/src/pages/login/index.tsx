import Container from '@mui/material/Container';
import GithubLoginButton from '../../components/LogIn/GithubLogInButton';
import Copyright from '../../components/Copyright';
import { Stack } from '@mui/material';


export default function LogIn() {

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
          <GithubLoginButton />
        </Stack>
        <Copyright sx={{ mt: 8, mb: 4 }} />
      </Container>
  );
}