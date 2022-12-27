import React from "react"
import { Link, useNavigate } from "react-router-dom";

export default function Logooutpage() {
  
  const navigate = useNavigate();

  function postdata(input){
    input.preventDefault();
    fetch(`http://localhost:3000/users/logout`, {
      method: "POST",
      credentials: "include",
      headers: { 'Content-Type': 'application/json' },
    })
    .then((response) => {
      if (response.ok) {
          console.log("Response:" + response)
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
          <form onSubmit={postdata} className="form">
                <button className="formsubmitbutton">
                  Submit
                </button>
            </form>
            <Link to = "/authenticated">
                <button className="headerbutton">
                  Homepage
                </button>
              </Link>
        </div>
    )
}