//this is the entry point for the react app
//this file sets out all the possible routes within the document
//first we wrap the entire app within a router to handle routing via Link
//

import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import { AuthContextProvider } from './components/Authenticate';
import NHomepage from './components/UnauthenticatedHomepage';
import AHomepage from './components/AuthenticatedHomepage';
import Loginpage from './components/Loginpage';
import Registerpage from './components/Registerpage'
import Threadspage from './components/Threadspage';
import Threadpage from './components/Threadpage';
import Createthread from './components/Createthreadpage';
import Createcomment from './components/Postcommentpage';
import Editcomment from './components/EditCommentpage';
import Logoutpage from './components/Logoutpage';



export default function Forum() {
  return (
  <AuthContextProvider>
    <Router>
      <Routes>
        <Route path="/" element={<NHomepage />} />
        <Route path="/authenticated" element={<AHomepage />} />
        <Route path="/users/login" element={<Loginpage />} />
        <Route path="/users/logout" element={<Logoutpage />} />
        <Route path="/users/signup" element={<Registerpage />} />
        <Route path="/:categoryid" element={<Threadspage />} />
        <Route path="/:categoryid/threads/:threadid" element={<Threadpage />} />
        <Route path="/:categoryid/threads/create" element={<Createthread />} />
        <Route path="/:categoryid/threads/:threadid/comments/create" element={<Createcomment />} />
        <Route path="/:categoryid/threads/:threadid/comments/:commentid" element={<Editcomment />} />
      </Routes>
    </Router>
  </AuthContextProvider>
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
