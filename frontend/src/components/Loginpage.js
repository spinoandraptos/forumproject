import React from "react"
import { useState, useEffect } from "react";
import { Link } from "react-router-dom";

export default function Loginpage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleUsername = (input) => {
      setUsername(input.target.value);
    };
  
    const handlePassword = (input) => {
      setPassword(input.target.value);
    };

    function showdata(){
      alert(username + password);
    }

    function postdata(input){
      input.preventDefault();
      fetch(`http://localhost:3000/users/login/authenticate`, {
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
            alert("Login Successful!")
        } else if (response.status === 400) {
          alert("Username not Found")
        } else if (response.status === 401) {
          alert("Incorrect Password")
        }
      })
    }
    

    return (
      <div className="loginpage">
        <div className="herocontent">
          <div className="herotext">
            Bop Fish Nation 
          </div>
        </div>
        <h1>Login Page</h1>
        <form onSubmit={postdata}>
          <div className="loginbox">
            <label htmlFor="username">
              Username
            </label>
            <br />
            <input id="username" type="text" value={username} onChange={handleUsername} />
          </div>
          <div className="loginbox">
            <label htmlFor="password">
              Password
            </label>
            <br />
            <input id="password" type="password" value={password} onChange={handlePassword}/>
          </div>
          <br />
          <button className="formsubmitbutton">
            Submit
          </button>
        </form>
        <div className="signup">
          Don't have an account?
          <Link to = "/users/signup">
            Sign Up!
          </Link>
        </div>
        <Link to = "/">
          <button className="footerbutton">
            Back to Homepage
          </button>
        </Link>
      </div>
    )
}