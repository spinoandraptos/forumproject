// this will serve as the component rendering the page for logging in
// we will create states to manage data for username input and password input
// upon inputting in the textbox data states for username and password will update
// then after submitting the data stored in the states will be sent to Go backend through fetch API
// if login is successful jwt is sent from backend and an item with title jwt and value username is stored in local storage
// user automatically redirected to the authenticated mainpage
// link and navigate allow users to navigate to other components

import React from "react"
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export default function Loginpage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handleUsername = (input) => {
      setUsername(input.target.value);
    };
  
    const handlePassword = (input) => {
      setPassword(input.target.value);
    };

    function postdata(input){
      input.preventDefault();
      fetch(`https://bopfishforum2.onrender.com/api/login`, {
        method: "POST",
        headers: { 'Content-Type': 'application/json' },
        credentials: "include",
        body: JSON.stringify({
          username: username,
          password: password
        })
      })
      .then((response) => {
        if (response.status===200) {
            console.log(response)
            localStorage.setItem("jwt", JSON.stringify(username))
            alert("Login Successful!")
            navigate("/authenticated")
            window.location.reload() 
        } else if (response.status === 500) {
          alert("Username not Found")
        } else if (response.status === 404) {
          alert("Incorrect Password")
        } else if (response.status === 401) {
          alert("Already Logged In")
        } else {
          alert("Error")
        }
      })
    }
    

    return (
      <div>
        <div className="herocontent">
          <div className="herotext">
            Bop Fish Nation 🦈 
          </div>
        </div>
        <div className="loginpage">
          <div className="logintitle">
            Login Page 🔒
          </div>
          <div className="loginformwhole">
            <div  className="loginform">
            <form onSubmit={postdata} className="form">
              <div className="loginlabel">
                <label htmlFor="username">
                  Username:
                </label>
                <br />
                <label htmlFor="password">
                  Password:
                </label>
              </div>
              <div className="loginbox">
                <input id="username" type="text" value={username} onChange={handleUsername} />
                <br />
                <input id="password" type="password" value={password} onChange={handlePassword}/>
                <br />
                <button className="formsubmitbutton">
                  Submit
                </button>
              </div>
            </form>
            </div>
          </div>
        </div>
        <div className="bottomlink">
          <div className="signup">
            Don't have an account?
            <br/>
            <Link to = "/users/signup">
              Sign Up Here!
            </Link>
          </div>
          <br />
          <Link to = "/">
            <button className="footerbutton">
              Back to Homepage
            </button>
          </Link>
        </div>
      </div>
    )
}