import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import HomePage from './pages/HomePage';
import Search from './pages/Search'; // Renamed IndexPage to SearchPage for clarity

const Router: React.FC = () => (
    <BrowserRouter>
        <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/search" element={<Search />} /> 
        </Routes>
    </BrowserRouter>
);

export default Router;