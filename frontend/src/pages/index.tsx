import React, { useState } from 'react';
import ReactDOM from 'react-dom/client';
import SearchBar from '../components/SearchBar';
import { ApolloProvider, ApolloClient, InMemoryCache } from '@apollo/client';

const client = new ApolloClient({
  uri: process.env.NODE_ENV === 'development' ? 'http://localhost:8080/query' : 'https://studygator-api.chasemacdonnell.net/query',
  cache: new InMemoryCache(),
});

const Index = () => {
  const [searchResults, setSearchResults] = useState([]);
  const [searchTerm, setSearchTerm] = useState(''); // Manage searchTerm state here.

  const handleSearchResults = (results) => {
    setSearchResults(results ?? []);
  };

  const renderSearchResults = () => {
    if (searchResults.length === 0 && searchTerm.trim() !== '') {
      return <div>No results found.</div>;
    }

    return searchResults.map((result) => (
      <div key={result.id} style={{ marginBottom: '20px', textAlign: 'center' }}>
        <img
          src={`http://localhost:8080/${result.imagePath}`}
          alt={result.classTitle}
          style={{ maxWidth: '800px', borderRadius: '10px' }}
        />
        <div>{result.classTitle}</div>
      </div>
    ));
  };

  return (
    <ApolloProvider client={client}>
      <div style={{ backgroundColor: '#FF4A00', minHeight: '100vh', padding: '20px' }}>
        <SearchBar onResults={handleSearchResults} setSearchTerm={setSearchTerm} searchTerm={searchTerm} />
        <div>
          <h2>Search Results:</h2>
          {renderSearchResults()}
        </div>
      </div>
    </ApolloProvider>
  );
};

const rootElement = document.getElementById('root');
if (rootElement) {
  ReactDOM.createRoot(rootElement).render(
    <React.StrictMode>
      <Index />
    </React.StrictMode>,
  );
}

export default Index;