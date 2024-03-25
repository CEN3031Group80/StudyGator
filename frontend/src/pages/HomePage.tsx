// src/pages/HomePage.tsx
import React from 'react';
import { useNavigate } from 'react-router-dom';

const HomePage = () => {
  const navigate = useNavigate();

  return (
    <div style={{ textAlign: 'center', padding: '50px' }}>
      <h1>Welcome to StudyGator!</h1>
      <p>Your portal to academic resources at UF.</p>
      <button
        style={{
          padding: '10px 20px',
          backgroundColor: '#FFA500',
          color: 'white',
          border: 'none',
          borderRadius: '5px',
          cursor: 'pointer',
          fontSize: '16px',
          marginTop: '20px',
        }}
        onClick={() => navigate('/search')} 
      >
        Search for Materials
      </button>
    </div>
  );
};

export default HomePage;