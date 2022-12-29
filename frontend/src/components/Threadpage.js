import React from "react"
import { useState, useEffect, useContext } from "react";
import { AuthContext } from "./Authenticate";
import { Link, useParams, useNavigate } from "react-router-dom";

export default function Threadpage() {

    const navigate = useNavigate();
    const [thread, setThread] = useState([]);
    const [comments, setComments] = useState([]);
    const {categoryid} = useParams();
    const {threadid} = useParams();
    const {flag, Fetchusername} = useContext(AuthContext);

    useEffect(()=>
      Fetchusername(),
    [])

    useEffect(() => {
      Promise.all([
        fetch(`http://localhost:3000/${categoryid}/threads/${threadid}`, {
            method: "GET",
            credentials: "include",
        }),
        fetch(`http://localhost:3000/${categoryid}/threads/${threadid}/comments`, {
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
      if (flag == true) {
        navigate("/authenticated")
      } else {
        navigate("/")
      }
    }

    function Clickpostcomment(){
      if (flag == true) {
        navigate(`/${categoryid}/threads/${threadid}/comments/create`)
      } else {
        alert("Please Login First")
      }
    }

    return (
        <div className = "allthreads">
        <div className="herocontent">
          <div className="herotext">
            Bop Fish Nation 
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
        {comments?.map(comment => (
            <div className="comments">
                <div className="comment" key={comment.id}>
                    <div className="commentheader">
                        {comment.authorusername}
                    </div>
                    <div className="commentcontent">
                        {comment.content}
                    </div>
                </div>
            </div>
        ))}
     </div>
    )
}