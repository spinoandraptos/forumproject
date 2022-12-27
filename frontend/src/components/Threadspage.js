import React from "react"
import { useState, useEffect } from "react";
import { Link, useParams } from "react-router-dom";

export default function Threadspage() {

    const [category, setCategory] = useState([]);
    const [threads, setThreads] = useState([]);
    const {id} = useParams();
    console.log(id);

    useEffect(() => {
      Promise.all([
        fetch(`http://localhost:8000/${id}`, {
            method: "GET"
        }),
        fetch(`http://localhost:8000/${id}/threads`, {
            method: "GET"
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

    return (
        <div className = "allthreads">
        <div className="herocontent">
          <div className="herotext">
            Bop Fish Nation 
          </div>
        </div>
        <header id = "homepageheader">
          <div className = "headerlinks">
            <Link to = {`/${category.id}/threads/create`}>
              <button className="headerbutton">
                Post A Thread
              </button>
            </Link>
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
        {threads?.map(thread => (
            <div className="threads">
                <div className="thread" key={thread.id}>
                    <div className="threadtitle">
                        {thread.title}
                    </div>
                    <div className="threadcontent">
                        {thread.content}
                    </div>
                    <div className="threadfooter">
                      <div>
                        Posted by {thread.authorusername}
                      </div>
                      <div className="threadfooterbutton">
                        <Link to = {`/${category.id}/threads/${thread.id}`}>
                          <button className="footerbutton threadbutton2">
                            Reply to this Thread
                          </button>
                        </Link>
                      </div>
                    </div>
                </div>
            </div>
        ))}
        <Link to = "/">
              <button className="footerbutton threadbutton">
                Back to Homepage
              </button>
            </Link>
     </div>
    )
}