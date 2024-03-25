import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
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

  //  'useEffect' triggers the search when the search term changes.
  useEffect(() => {
    if (searchTerm.trim() !== '') {
      handleSearch();
    }
  }, [searchTerm]); // dependency array ensures this runs only if searchTerm changes

  return (
    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '20px' }}>
      <div>
        <input
          className="input"
          type="text"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          placeholder="Search for a class"
        />
        <button
          className="button"
          onClick={handleSearch}
        >
          Search
        </button>
      </div>
      <Link to="/">
        <img src={`http://localhost:8080/images/studygator1.png`} alt="StudyGator Logo" style={{ maxWidth: '100px', height: 'auto' }} />
      </Link>
    </div>
  );
};

export default SearchBar;