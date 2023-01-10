// this will serve as the component rendering the page for logging out
// on clicking logout button request is sent to Go backend through fetch API
// upon logging out the jwt item with value username stored in local storage is removed and jwt cookie terminated
// user automatically redirected to the unauthenticated mainpage
// link and navigate allow users to navigate to other components

import React from "react"
import { Link, useNavigate } from "react-router-dom";

export default function Logooutpage() {
  
  const navigate = useNavigate();

  function postdata(input){
    input.preventDefault();
    fetch(`https://bopfishforum2.onrender.com/api/users/logout`, {
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
          window.location.reload() 
      } else {
        alert("Logout Failed")
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