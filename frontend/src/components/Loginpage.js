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
      fetch(`http://localhost:3000/users/login`, {
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
            Bop Fish Nation 
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