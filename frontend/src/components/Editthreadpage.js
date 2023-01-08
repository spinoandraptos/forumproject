import React, { useContext, useEffect } from "react";
import { AuthContext } from "./Authenticate";
import { Link, useNavigate, useParams }  from "react-router-dom";
import { useState } from "react";

export default function Editthread() {
   
    const navigate = useNavigate();
    const {categoryid, threadid} = useParams();
    const [thread, setThread] = useState({});
    const {flag, Fetchusername} = useContext(AuthContext);
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");

    const handleContent = (input) => {
      setContent(input.target.value);
    };

    const handleTitle = (input) => {
      setTitle(input.target.value);
    };
    
    useEffect(()=>
      Fetchusername(),
    [])

    useEffect(()=>{
      console.log(categoryid, threadid)
      fetch(`https://bopfishforum2.onrender.com/api/${categoryid}/threads/${threadid}`, {
        method: "GET",
        credentials: "include",
    })
    .then((threadresponse)=> {
      if(threadresponse.ok){
        return threadresponse.json()
      }
    })
    .then((threaddata)=>{
      console.log("Data1:" + threaddata)
      setThread(threaddata)
      console.log("Data1:" + JSON.stringify(thread))
    })
  },[])

    function Clickhomepage(){
        if (flag === true) {
          navigate("/authenticated")
        } else {
          navigate("/")
        }
      }

      function postdata(input){

        if (flag === true) {
        input.preventDefault();

        if(title === ""){
          fetch(`https://bopfishforum2.onrender.com/api/${categoryid}/threads/${threadid}/content`, {
            method: "PUT",
            credentials: "include",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
              Content: content
            })
          })
          .then((response) => {
            if (response.ok) {
              console.log("Response:" + response)
              alert("Thread Updated Successfully!")
              navigate(`/${categoryid}`)
            } else if (response.status===401) {
              alert("Server Does Not Detect JWT")
            } else {
              alert("Error: Thread Cannot Update")
            }
          })
      } else if (content === ""){
        fetch(`https://bopfishforum2.onrender.com/api/${categoryid}/threads/${threadid}/title`, {
          method: "PUT",
          credentials: "include",
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            Title: title
          })
        })
        .then((response) => {
          if (response.ok) {
            console.log("Response:" + response)
            alert("Thread Updated Successfully!")
            navigate(`/${categoryid}`)
          } else if (response.status===401) {
            alert("Server Does Not Detect JWT")
          } else {
            alert("Error: Thread Cannot Update")
          }
        })
      } else {
        fetch(`https://bopfishforum2.onrender.com/api/${categoryid}/threads/${threadid}`, {
          method: "PUT",
          credentials: "include",
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            Title: title,
            Content: content
          })
        })
        .then((response) => {
          if (response.ok) {
            console.log("Response:" + response)
            alert("Thread Updated Successfully!")
            navigate(`/${categoryid}`)
          } else if (response.status===401) {
            alert("Server Does Not Detect JWT")
          } else {
            alert("Error: Thread Cannot Update")
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
            <Link to = {`/${categoryid}`}>
              <button className="headerbutton">
                Back to Threads
              </button>
            </Link>
          </div>
        </header>
        <div className="editcomment">
          <div className="updatethread">
            <div className="comment" key={thread.id}>
              <div className="threadtitle2">
                {thread.title}
              </div>
              <div className="threadcontent2">
                {thread.content}
              </div>
            </div>
            <div>
              <form onSubmit={postdata} className="editcommentform">
                <div className="loginbox commentbox">
                  <textarea placeholder="Update thread title here (leave blank if not changing)" rows={2} cols={50} maxLength={500} minLength={1} id="title" type="commenttext" value={title} onChange={handleTitle}/>
                  <br />
                  <textarea placeholder="Update thread description here (leave blank if not changing)" rows={5} cols={50} maxLength={1000} minLength={1} id="content" type="commenttext" value={content} onChange={handleContent}/>
                  <br />
                  <button className="formsubmitbutton">
                    Submit
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    )
    
}