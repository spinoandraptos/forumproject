import React, { useContext, useEffect } from "react";
import { AuthContext } from "./Authenticate";
import { Link, useNavigate, useParams }  from "react-router-dom";
import { useState } from "react";

export default function Createcomment() {

  const navigate = useNavigate();

  const {categoryid, threadid} = useParams()
  console.log(categoryid, threadid)


  function postdata(input){
    input.preventDefault();
    fetch(`/${categoryid}/threads/${threadid}/comments`, {
      method: "POST",
      credentials: "include",
      headers: { 'Content-Type': 'application/json' },
    })
    .then((response) => {
      if (response.ok) {
          console.log("Response:" + response)
          localStorage.removeItem("jwt");
          alert("Logout Successful!")
          navigate("/")
      } else {
        alert("Logout Failed")
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
          <div className="logout">
            <div className="logoutbox">
              <form onSubmit={postdata} className="form">
                <button className="formsubmitbuttonblue">
                  Logout
                </button>
              </form>
            </div>
          </div>
          <div className="logoutfooter">
            <Link to = "/authenticated">
              <button className="footerbutton">
                Return To Homepage
              </button>
            </Link>
          </div>
        </div>
    )

}