import React from "react";
import { Link, useNavigate, useParams }  from "react-router-dom";
import { useState } from "react";

export default function Createthread() {
  
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const navigate = useNavigate();
  const {id} = useParams();

  const handleTitle = (input) => {
    setTitle(input.target.value);
  };

  const handleContent = (input) => {
    setContent(input.target.value);
  };

  function postdata(input){
    input.preventDefault();
    fetch(`http://localhost:3000/${id}/threads`, {
      method: "POST",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        title: title,
        content: content
      })
    })
    .then((response) => {
      if (response.ok) {
          console.log("Response:" + response)
          alert("Thread Posted Successfully!")
          navigate("/:id")
      } else if (response.status===401) {
        alert("Please Login First")
      } else {
        alert("Error: Thread Cannot Posted")
      }
    })
  }

    return(
      <div>
      <div className="herocontent">
        <div className="herotext">
          Bop Fish Nation 
        </div>
      </div>
      <div className="loginpage">
        <div className="logintitle">
          Input Thread Details
        </div>
        <div className="loginformwhole">
          <div  className="loginform">
          <form onSubmit={postdata} className="form">
            <div className="loginlabel">
              <label htmlFor="title">
                Title:
              </label>
              <br />
              <label htmlFor="content">
                Content:
              </label>
            </div>
            <div className="loginbox">
              <input id="title" type="text" value={title} onChange={handleTitle} />
              <br />
              <input id="content" type="text" value={content} onChange={handleContent}/>
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
        <Link to = {`/${id}`}>
          <button className="footerbutton">
            Back to Threads
          </button>
        </Link>
      </div>
    </div>
    )
}