// Import 'useEffect' to handle side effects.
import React, { useState, useEffect } from 'react';
import { useApolloClient } from '@apollo/client';
import gql from 'graphql-tag';

const SEARCH_QUERY = gql`
  query SearchByClassTitle($title: String!) {
    SearchByClassTitle(term: $title) {
      id
      classTitle
      imagePath
    }
  }
`;

const SearchBar = ({ onResults, setSearchTerm, searchTerm }) => {
  const client = useApolloClient();

  const handleSearch = async () => {
    try {
      const { data } = await client.query({
        query: SEARCH_QUERY,
        variables: { title: searchTerm },
      });
      onResults(data.SearchByClassTitle);
    } catch (error) {
      console.error('Error executing search:', error);
      onResults([]);
    }
  };

  // Use 'useEffect' to trigger the search when the search term changes.
  useEffect(() => {
    if (searchTerm.trim() !== '') {
      handleSearch();
    }
  }, [searchTerm]); // Dependency array ensures this runs only if searchTerm changes.

  return (
    <div style={{ marginBottom: '20px' }}>
      <input
        type="text"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        placeholder="Search for a class"
        style={{ padding: '10px', marginRight: '5px', borderRadius: '5px' }}
      />
      <button
        onClick={handleSearch} // This button now explicitly triggers the search.
        style={{
          padding: '10px',
          backgroundColor: '#0021A5',
          color: 'white',
          border: 'none',
          borderRadius: '5px',
          cursor: 'pointer'
        }}
      >
        Search
      </button>
    </div>
  );
};

export default SearchBar;