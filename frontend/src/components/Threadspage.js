// this will serve as the component rendering the threads page for a category
// we will create states to manage the data of retrieved category and threads
// upon render, the data for category and threads (displayed on screen) are fetched via fetch API from Go backend
// by checking the author username of each of the threads with that of current user, we separately provide additional functionalities like edit and delete for threads the current user posted
// link and navigate allow users to navigate to other components

import React from "react"
import { useState, useEffect, useContext } from "react";
import { AuthContext } from "./Authenticate";
import { Link, useParams, useNavigate } from "react-router-dom";
import moment from "moment";

export default function Threadspage() {

    const navigate = useNavigate();
    const [category, setCategory] = useState([]);
    const [threads, setThreads] = useState([]);
    const {categoryid} = useParams();
    const {flag, Fetchusername} = useContext(AuthContext)
    console.log(categoryid);

    useEffect(()=>
      Fetchusername(),
    [])

    useEffect(() => {
      Promise.all([
        fetch(`https://bopfishforum2.onrender.com/api/${categoryid}`, {
            method: "GET",
            credentials: "include",
        }),
        fetch(`https://bopfishforum2.onrender.com/api/${categoryid}/threads`, {
            method: "GET",
            credentials: "include",
        })
      ])
        .then(([categoryresponse, threadsresponse]) => 
        Promise.all([categoryresponse.json(), threadsresponse.json()])
        )
        .then(([categorydata, threadsdata]) => {
            console.log("Data1:" + categorydata)
            setCategory(categorydata);
            console.log("Final1:" + JSON.stringify(category))
            console.log("Data2:" + threadsdata)
            setThreads(threadsdata);
            console.log("Final2:" + JSON.stringify(threads))
        });
    }, []);

    function Clickhomepage(){
      if (flag === true) {
        navigate("/authenticated")
      } else {
        navigate("/")
      }
    }

    function Clickpostthread(){
      if (flag === true) {
        navigate(`/${category.id}/threads/create`)
      } else {
        alert("Please Login First")
      }
    }

    function Clickdeletethread(value){
      fetch(`https://bopfishforum2.onrender.com/api/${categoryid}/threads/${value}`, {
        method: "DELETE",
        credentials: "include",
        headers: { 'Content-Type': 'application/json' }
      })
      .then((response) => {
        if (response.ok) {
            console.log("Response:" + response)
            alert("Thread Deletion Successful!")
            window.location.reload()
        } else if (!response.ok) {
          alert("Thread Deletion Failed")
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
              <button className="headerbutton" onClick={Clickpostthread}>
                Post A Thread
              </button>
          </div>
        </header>
        <div className="threadcategory">
          <div className="threadcategorytitle">
            {category.title}
          </div>
          <div className="threadcategorydescription">
            {category.description}
          </div>
        </div>
        {threads?.map(thread => (thread.authorusername === JSON.parse(localStorage.getItem("jwt"))? (
            <div className="threads">
                <div key={thread.id}>
                    <div className="threadtitle">
                        {thread.title}
                    </div>
                    <div className="threadcontent">
                        {thread.content}
                    </div>
                    <div className="threadfooter">
                      <div>
                        Posted by {thread.authorusername} at {moment(thread.createdat).format('DD-MM-YYYY hh:mm A')}
                      </div>
                      <div className="threadfooterbutton">
                      <Link to = {`/${category.id}/threads/${thread.id}`}>
                          <button className="footerbutton threadbutton2">
                            Edit Thread
                          </button>
                        </Link>
                        <button className="footerbutton threadbutton2" onClick={()=>{Clickdeletethread(thread.id)}}>
                            Delete Thread
                          </button>
                        <Link to = {`/${category.id}/threads/${thread.id}/comments`}>
                          <button className="footerbutton threadbutton2">
                            See Comments
                          </button>
                        </Link>
                      </div>
                    </div>
                </div>
            </div>
          )
        : (
          <div className="threads">
                <div key={thread.id}>
                    <div className="threadtitle">
                        {thread.title}
                    </div>
                    <div className="threadcontent">
                        {thread.content}
                    </div>
                    <div className="threadfooter">
                      <div>
                        Posted by {thread.authorusername} at {moment(thread.CreatedAt).format('YYYY-MM-DD hh:mm')}
                      </div>
                      <div className="threadfooterbutton">
                        <Link to = {`/${category.id}/threads/${thread.id}/comments`}>
                          <button className="footerbutton threadbutton2">
                            See Comments
                          </button>
                        </Link>
                      </div>
                    </div>
                </div>
            </div>
          )
        ))}
     </div>
    )
}