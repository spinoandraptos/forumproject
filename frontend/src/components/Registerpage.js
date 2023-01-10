// this will serve as the component rendering the page for registering a new account
// we will create states to manage data for username input and password input
// upon inputting in the textbox data states for username and password will update
// then after submitting the data stored in the state will be sent to Go backend through fetch API
// link and navigate allow users to navigate to other components

import React from "react"
import { Link, useNavigate } from "react-router-dom";
import { useState } from "react";

export default function Registerpage() {
  
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
    fetch(`https://bopfishforum2.onrender.com/api/users/signup`, {
      method: "POST",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        username: username,
        password: password
      })
    })
    .then((response) => {
      if (response.ok) {
          console.log("Response:" + response)
          alert("User Creation Successful!")
          navigate("/users/login")
      } else if (!response.ok) {
        alert("User Creation Failed")
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
          Signup Page
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
            Already Have An Account?
            <br/>
            <Link to = "/users/login">
              Login Here!
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