import React, { useContext, useEffect } from "react";
import { AuthContext } from "./Authenticate";
import { Link, useNavigate, useParams }  from "react-router-dom";
import { useState } from "react";

export default function Editcomment() {
   
    const navigate = useNavigate();
    const {categoryid, threadid, commentid} = useParams();
    const [comment, setComment] = useState({});
    const {flag, Fetchusername} = useContext(AuthContext);
    const [content, setContent] = useState("");

    const handleContent = (input) => {
      setContent(input.target.value);
    };
    
    useEffect(()=>
      Fetchusername(),
    [])

    useEffect(()=>{
      console.log(categoryid, threadid, commentid)
      fetch(`http://localhost:3000/${categoryid}/threads/${threadid}/comments/${commentid}`, {
        method: "GET",
        credentials: "include",
    })
    .then((commentresponse)=> {
      if(commentresponse.ok){
        return commentresponse.json()
      }
    })
    .then((commentdata)=>{
      console.log("Data1:" + commentdata)
      setComment(commentdata)
      console.log("Data1:" + JSON.stringify(comment))
    })
  },[])

    function Clickhomepage(){
        if (flag == true) {
          navigate("/authenticated")
        } else {
          navigate("/")
        }
      }

      function postdata(input){

        if (flag == true) {
        input.preventDefault();
        fetch(`/${categoryid}/threads/${threadid}/comments/${commentid}`, {
          method: "PUT",
          credentials: "include",
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            Content: content,
          })
        })
        .then((response) => {
          if (response.ok) {
              console.log("Response:" + response)
              alert("Comment Updated Successfully!")
              navigate(`/${categoryid}/threads/${threadid}`)
          } else if (response.status===401) {
            alert("Server Does Not Detect JWT")
          } else {
            alert("Error: Comment Cannot Update")
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
            <Link to = {`/${categoryid}/threads/${threadid}`}>
              <button className="headerbutton">
                Back to Comments
              </button>
            </Link>
          </div>
        </header>
        <div className="editcomment">
          <div className="comments2">
            <div className="comment" key={comment.id}>
              <div className="commentheader2">
                  {comment.authorusername}
              </div>
              <div className="commentcontent2">
                  {comment.content}
              </div>
            </div>
            <div>
              <div>
                <div>
                  <form onSubmit={postdata} className="editcommentform">
                    <div className="loginbox commentbox">
                      <textarea placeholder="Update your comment" rows={5} cols={50} maxLength={1000} minLength={1} id="content" type="commenttext" value={content} onChange={handleContent}/>
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
        </div>
      </div>
    )
    
}