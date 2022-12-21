import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Homepage from "./HomepageLayout";
import Userpage from "./userpage";
import Loginpage from "./LoginpageLayout";
import Registerpage from "./RegisterpageLayout";

//this file sets up the routing within the React application where links between pages are used
export default (
  <Router>
    <Routes>
      <Route path="/" element={<Homepage />} />
      <Route path="/user" element={<Userpage />} />
      <Route path="/login" element={<Loginpage />} />
      <Route path="/register" element={<Registerpage />} />
    </Routes>
  </Router>
);