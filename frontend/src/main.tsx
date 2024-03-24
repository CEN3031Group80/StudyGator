import React from 'react';
import ReactDOM from 'react-dom/client';
import Router from './Router';
import { ApolloClient, ApolloProvider, InMemoryCache } from '@apollo/client';

const uri = import.meta.env.DEV ? 'http://localhost:8080/query' : 'https://studygator-api.chasemacdonnell.net/query';

const client = new ApolloClient({
  uri: uri,
  cache: new InMemoryCache(),
  credentials: 'include', // Add this line if you need to send cookies or auth headers
});

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <Router />
    </ApolloProvider>
  </React.StrictMode>
);