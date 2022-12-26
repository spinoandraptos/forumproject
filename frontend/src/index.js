//this is the entry point for the react app
//this file sets out all the possible routes within the document
//first we wrap the entire app within a router to handle routing via Link
//

import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import Homepage from './components/Homepage';
import Loginpage from './components/Loginpage';
import Registerpage from './components/Registerpage'
import Threadspage from './components/Threadspage';
import Threadpage from './components/Threadpage';
import Createthread from './components/Createthreadpage';
import Createcomment from './components/Postcommentpage';
import Logoutpage from './components/Logoutpage';

export default function Forum() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Homepage />} />
        <Route path="/users/login" element={<Loginpage />} />
        <Route path="/users/logout" element={<Logoutpage />} />
        <Route path="/users/signup" element={<Registerpage />} />
        <Route path="/:id" element={<Threadspage />} />
        <Route path="/:id/threads/:id" element={<Threadpage />} />
        <Route path="/:id/threads/create" element={<Createthread />} />
        <Route path="/:id/threads/:id/comments/create" element={<Createcomment />} />
      </Routes>
    </Router>
  );
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Forum />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
