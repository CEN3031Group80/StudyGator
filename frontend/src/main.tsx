import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';

import React from 'react'
import ReactDOM from 'react-dom/client'
import Router from './Router.tsx'
import { ApolloClient, ApolloProvider, InMemoryCache } from '@apollo/client';
import { CssBaseline, ThemeProvider, createTheme } from '@mui/material';

export let apiUrl = "https://studygator-api.chasemacdonnell.net";
export let rootUrl = "https://studygator.chasemacdonnell.net"

if (import.meta.env.DEV) {
  apiUrl = "/api";
  rootUrl = "http://localhost:5174"
}

const client = new ApolloClient({
  uri: apiUrl + "/query",
  cache: new InMemoryCache(),
  credentials: "include"
});

const theme = createTheme();

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <ApolloProvider client={client}>
        <CssBaseline />
        <Router />
      </ApolloProvider>
    </ThemeProvider>
  </React.StrictMode>,
)
