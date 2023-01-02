import React, { useContext, useEffect } from "react";
import { AuthContext } from "./Authenticate";
import { useNavigate, useParams }  from "react-router-dom";
import { useState } from "react";

export default function Edituser() {
   
    const navigate = useNavigate();
    const {categoryid, threadid} = useParams();
    const [user, setUser] = useState({});
    const {flag, Fetchusername} = useContext(AuthContext);
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleUsername = (input) => {
      setUsername(input.target.value);
    };

    const handlePassword = (input) => {
      setPassword(input.target.value);
    };
    
    useEffect(()=>
      Fetchusername(),
    [])

    useEffect(()=>{
      fetch(`http://localhost:3000/users/${JSON.parse(localStorage.getItem("jwt"))}`, {
      method: "POST",
      credentials: "include",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        Username: JSON.parse(localStorage.getItem("jwt"))
      })
      })
      .then((response)=> {
        if (response.ok) {
          return response.json();
        }
      })
      .then((userdata)=> {
        setUser(userdata)
      })
  },[])

    function Clickhomepage(){
        if (flag == true) {
          navigate("/authenticated")
        } else {
          navigate("/")
        }
      }

    function Deleteuser(){
      fetch(`/users/${user.id}`, {
        method: "DELETE",
        credentials: "include",
        headers: { 'Content-Type': 'application/json' },
      })
      .then((response) => {
        if (response.ok) {
          localStorage.removeItem("jwt");
          alert("User Deleted Successfully!")
          navigate(`/`)
          window.location.reload() 
        } else if (response.status===401) {
          alert("Server Does Not Detect JWT")
        } else {
          alert("Error: User Cannot be Deleted")
        }
      })
    }

    function postdata(input){

        if (flag == true) {
        input.preventDefault();

        if(username === ""){
          fetch(`/users/${user.id}/password`, {
            method: "PUT",
            credentials: "include",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
              Password: password
            })
          })
          .then((response) => {
            if (response.ok) {
              console.log("Response:" + response)
              alert("Password Updated Successfully!")
              navigate(`/authenticated`)
            } else if (response.status===401) {
              alert("Server Does Not Detect JWT")
            } else {
              alert("Error: Password Cannot Update")
            }
          })
      } else if (password === ""){
        fetch(`/users/${user.id}/username`, {
          method: "PUT",
          credentials: "include",
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            Username: username
          })
        })
        .then((response) => {
          if (response.ok) {
            console.log("Response:" + response)
            alert("Username Updated Successfully!")
            localStorage.setItem("jwt", JSON.stringify(username))
            navigate(`/authenticated`)
          } else if (response.status===401) {
            alert("Server Does Not Detect JWT")
          } else {
            alert("Error: Username Cannot Update")
          }
        })
      } else {
        fetch(`/users/${user.id}`, {
          method: "PUT",
          credentials: "include",
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            Username: username,
            Password: password
          })
        })
        .then((response) => {
          if (response.ok) {
            console.log("Response:" + response)
            alert("User Updated Successfully!")
            localStorage.setItem("jwt", JSON.stringify(username))
            navigate(`/authenticated`)
          } else if (response.status===401) {
            alert("Server Does Not Detect JWT")
          } else {
            alert("Error: User Cannot Update")
          }
        })
      }
    } else {
        alert("Please Login First")
        navigate("/users/login")
      }
    }
    
    return (
        <div>
        <div className="herocontent">
          <div className="herotext">
            Bop Fish Nation 🦈 
          </div>
        </div>
        <header id = "homepageheader">
          <div className = "headerlinks">
              <button className="headerbutton" onClick={Clickhomepage}>
                Back to Homepage
              </button>
            <button className="headerbuttonwarning" onClick={Deleteuser}>
                Delete User
            </button>
          </div>
        </header>
        <div className="editcomment">
          <div className="userinfo">
            <div className="editusertitle">
                Edit User Info
            </div>
            <form onSubmit={postdata} className="editcommentform">
                <div className="loginbox commentbox">
                    <input id="username" placeholder="Update Username Here (Leave blank if no changes)" size={50} type="text" value={username} onChange={handleUsername} />
                    <br />
                    <input id="password" placeholder="Update Password Here (Leave blank if no changes)" size={50} type="password" value={password} onChange={handlePassword}/>
                    <br />
                    <button className="formsubmitbutton">
                       Submit
                    </button>
                </div>
            </form>
          </div>
        </div>
     </div>
    )
    
}