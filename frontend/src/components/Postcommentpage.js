import React, { useContext, useEffect } from "react";
import { AuthContext } from "./Authenticate";
import { Link, useNavigate, useParams }  from "react-router-dom";
import { useState } from "react";

export default function Createcomment() {

  const navigate = useNavigate();
  const {categoryid, threadid} = useParams();
  const {flag, userid, Fetchusername} = useContext(AuthContext);
  const [thread, setThread] = useState([]);
  const [content, setContent] = useState("");

  const handleContent = (input) => {
    setContent(input.target.value);
  };

  useEffect(() => 
    Fetchusername(), 
  [])

  useEffect(() => {
    fetch(`http://localhost:3000/api/${categoryid}/threads/${threadid}`, {
            method: "GET",
            credentials: "include",
        })
        .then((threadresponse) => {
          if (threadresponse.ok) {
              return threadresponse.json();
          }
        })
          .then((threaddata) => {
              setThread(threaddata);
          });
  },[])

  function postdata(input){

    if (flag === true) {
    input.preventDefault();
    fetch(`/api/${categoryid}/threads/${threadid}/comments`, {
      method: "POST",
      credentials: "include",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        Content: content,
        AuthorID: parseInt(userid),
        ThreadID: parseInt(threadid)
      })
    })
    .then((response) => {
      if (response.ok) {
          console.log("Response:" + response)
          alert("Comment Posted Successfully!")
          navigate(`/${categoryid}/threads/${threadid}/comments`)
      } else if (response.status===401) {
        alert("Server Does Not Detect JWT")
      } else {
        alert("Error: Comment Cannot be Posted")
      }
    })
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
    <div className="logintitle">
        Input Comment 😊
      </div>
    <div className="thread commentthread">
          <div className="threadtitle">
            {thread.title}
          </div>
          <div className="threadcontent commentpagethreadcontent">
            {thread.content}
            <div className="threadfooter">
            Posted by {thread.authorusername}
          </div>
          </div>
        </div>
    <div className="loginpage">
      <div className="loginformwhole">
        <div  className="commentform">
        <form onSubmit={postdata} className="form">
          <div className="loginlabel">
            <label htmlFor="content">
              Comment:
            </label>
          </div>
          <div className="loginbox commentbox">
            <textarea placeholder="What are your thoughts?" maxLength={1000} minLength={1} rows={5} cols={40} id="content" type="commenttext" value={content} onChange={handleContent}/>
            <br />
            <button className="formsubmitbutton">
              Submit
            </button>
          </div>
        </form>
        </div>
      </div>
    </div>
    <div className="bottomlink backbutton">
      <Link to = {`/${categoryid}/threads/${threadid}/comments`}>
        <button className="footerbutton">
          Back to Comments
        </button>
      </Link>
    </div>
  </div>
  )

}