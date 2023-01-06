import React from "react"
import moment from "moment";
import { useState, useEffect, useContext } from "react";
import { AuthContext } from "./Authenticate";
import { Link, useParams, useNavigate } from "react-router-dom";

export default function Threadpage() {

    const navigate = useNavigate();
    const [thread, setThread] = useState({});
    const [comments, setComments] = useState([]);
    const {categoryid, threadid} = useParams();
    const {flag, Fetchusername} = useContext(AuthContext);

    useEffect(()=>
      Fetchusername(),
    [])

    useEffect(() => {
      Promise.all([
        fetch(`/api/${categoryid}/threads/${threadid}`, {
            method: "GET",
            credentials: "include",
        }),
        fetch(`/api/${categoryid}/threads/${threadid}/comments`, {
            method: "GET",
            credentials: "include",
        })
      ])
        .then(([threadresponse, commentsresponse]) => 
        Promise.all([threadresponse.json(), commentsresponse.json()])
        )
        .then(([threaddata, commentsdata]) => {
            console.log("Data1:" + threaddata)
            setThread(threaddata);
            console.log("Final1:" + JSON.stringify(thread))
            console.log("Data2:" + commentsdata)
            setComments(commentsdata);
            console.log("Final2:" + JSON.stringify(comments))
        });
    }, []);

    function Clickhomepage(){
      if (flag === true) {
        navigate("/authenticated")
      } else {
        navigate("/")
      }
    }

    function Clickpostcomment(){
      if (flag === true) {
        navigate(`/${categoryid}/threads/${threadid}/comments/create`)
      } else {
        alert("Please Login First")
      }
    }

    function Clickdeletecomment(value){
        fetch(`/api/${categoryid}/threads/${threadid}/comments/${value}`, {
          method: "DELETE",
          headers: { 'Content-Type': 'application/json' }
        })
        .then((response) => {
          if (response.ok) {
              console.log("Response:" + response)
              alert("Comment Deletion Successful!")
              window.location.reload()
          } else if (!response.ok) {
            alert("Comment Deletion Failed")
          }
        })
      }

    return (
        <div className = "allthreads">
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
            <button className="headerbutton" onClick={Clickpostcomment}>
                Post A Comment
              </button>
          </div>
        </header>
        <div className="thread">
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
        {comments?.map(comment => (comment.authorusername === JSON.parse(localStorage.getItem("jwt"))? (
            <div className="comments">
                <div className="comment" key={comment.id}>
                    <div className="commentheader">
                        😎 {comment.authorusername} :
                    </div>
                    <div className="commentcontent">
                        {comment.content}
                    </div>
                    <div className="commentfooter">
                      <div className="footerdate">
                        Posted at {moment(comment.CreatedAt).format('YYYY-MM-DD hh:mm')}
                      </div>
                      <Link to = {`/${categoryid}/threads/${threadid}/comments/${comment.id}`}>
                        <button className="footerbutton2">
                          Edit Comment
                        </button>
                      </Link>
                      <button className="footerbutton3" onClick={()=>{Clickdeletecomment(comment.id)}}>
                        Delete Comment
                      </button>
                    </div>
                </div>
            </div>
        )
        : (
          <div className="comments">
                <div className="comment" key={comment.id}>
                    <div className="commentheader">
                        🙂 {comment.authorusername} :
                    </div>
                    <div className="commentcontent">
                        {comment.content}
                    </div>
                    <div className="commentfooter">
                      <div className="footerdate">
                        Posted at {moment(comment.CreatedAt).format('YYYY-MM-DD hh:mm')}
                      </div>
                    </div>
                    </div>
                </div>
         )
        ))}
     </div>
    )
}