import React from "react"
import { Link, useNavigate } from "react-router-dom";

export default function Logooutpage() {
  
  const navigate = useNavigate();

  function postdata(input){
    input.preventDefault();
    fetch(`http://localhost:3000/api/users/logout`, {
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