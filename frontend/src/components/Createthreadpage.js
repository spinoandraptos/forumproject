// this will serve as the component rendering the page for posting threads
// we will create states to manage data for title input and content input
// upon inputting in the textbox data states for title and content will update
// then after submitting the data stored in the states will be sent to Go backend through fetch API
// link and navigate allow users to navigate to other components
// note: with the use of flag, submission is only possible if an user is logged in

import React, { useContext, useEffect } from "react";
import { AuthContext } from "./Authenticate";
import { Link, useNavigate, useParams }  from "react-router-dom";
import { useState } from "react";


export default function Createthread() {
  
  const navigate = useNavigate();
  const {flag, userid, Fetchusername} = useContext(AuthContext)
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const {categoryid} = useParams();

  const handleTitle = (input) => {
    setTitle(input.target.value);
  };
  const handleContent = (input) => {
    setContent(input.target.value);
  };
  
  useEffect(() => 
    Fetchusername(), 
  [])

  function postdata(input){
    input.preventDefault();

    if (flag === true) {
  
     fetch(`https://bopfishforum2.onrender.com/api/${categoryid}/threads`, {
      method: "POST",
      credentials: "include",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        Title: title,
        Content: content,
        AuthorID: parseInt(userid),
        CategoryID: parseInt(categoryid)
      })
    })
    .then((response) => {
      if (response.ok) {
          console.log("Response:" + response)
          alert("Thread Posted Successfully!")
          navigate(`/${categoryid}`)
      } else if (response.status===401) {
        alert("Server Does Not Detect JWT")
      } else {
        alert("Error: Thread Cannot Posted")
      }
    })
   } else {
    alert("Please Login First")
    navigate("/users/login")
   }
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
          Input Thread Details 😊
        </div>
        <div className="loginformwhole">
          <div className="postform">
          <form onSubmit={postdata} className="form">
            <div className="postlabel">
              <label htmlFor="title">
                Title:
              </label>
              <br />
              <label htmlFor="content">
                Content:
              </label>
            </div>
            <div className="loginbox">
              <textarea placeholder="Write Title Here" id="title" maxLength={500} minLength={1} type="text" rows={1} cols={40} value={title} onChange={handleTitle} />
              <br />
              <textarea placeholder="Write Content Here" id="content" maxLength={1000} minLength={1} type="text" rows={5} cols={40} value={content} onChange={handleContent}/>
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
        <Link to = {`/${categoryid}`}>
          <button className="footerbutton">
            Back to Threads
          </button>
        </Link>
      </div>
    </div>
    )
}