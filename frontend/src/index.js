//this is the entry point for the react app
//this file renders app

import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import Homepage from './components/HomepageLayout';
import { Routes, Route, BrowserRouter } from 'react-router-dom';
import Loginpage from "./components/LoginpageLayout";
import Registerpage from "./components/RegisterpageLayout";


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <Homepage>
        <Routes>
          <Route exact path="/login" element={<Loginpage />} />
          <Route exact path="/register" element={<Registerpage />} />
        </Routes>
      </Homepage>
    </BrowserRouter>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
