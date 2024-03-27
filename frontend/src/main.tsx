import React from 'react'
import ReactDOM from 'react-dom/client'
import Router from './Router.tsx'
import { ApolloClient, ApolloProvider, InMemoryCache } from '@apollo/client';

let client = new ApolloClient({
  uri: 'https://studygator-api.chasemacdonnell.net/query',
  cache: new InMemoryCache(),
  credentials: "include"
});

if (import.meta.env.DEV) {
  client = new ApolloClient({
    uri: '/api/query',
    cache: new InMemoryCache(),
    credentials: "include"
  });
}

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <Router />
    </ApolloProvider>
  </React.StrictMode>,
)
