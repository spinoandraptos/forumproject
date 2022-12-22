import React from "react"
import { useState, useEffect } from "react";
import { Link, useParams } from "react-router-dom";

export default function Threadspage() {

    const [category, setCategory] = useState([]);
    const [threads, setThreads] = useState([]);
    const {id} = useParams();
    console.log(id);

    useEffect(() => {
        fetch(`http://localhost:3000/${id}`, {
            method: "GET"
        })
        .then((response) => {
        if (response.ok) {
            console.log("Response:" + response)
            return response.json();
        }
        throw new Error("Fetch Error");
        })
        .then((categorydata) => {
            console.log("Data1:" + categorydata)
            setCategory(categorydata);
            console.log("Final1:" + JSON.stringify(category))
        });
    }, []);

    useEffect(() => {
        fetch(`http://localhost:3000/${id}/threads`, {
            method: "GET"
        })
        .then((response) => {
        if (response.ok) {
            console.log("Response:" + response)
            return response.json();
        }
        throw new Error("Fetch Error");
        })
        .then((threadsdata) => {
            console.log("Data2:" + threadsdata)
            setThreads(threadsdata);
            console.log("Final2:" + JSON.stringify(threads))
        });
    }, []);

    return (
        <div className = "allcategories">
        <div className="herocontent">
          <div className="herotext">
            Bop Fish Nation 
          </div>
        </div>
        <header id = "homepageheader">
          <div className = "headerlinks">
            <Link to = "/users/signup">
              <button className="headerbutton">
                Register
              </button>
            </Link>
            <Link to = "/users/login">
              <button className="headerbutton">
                Login
              </button>
            </Link>
          </div>
        </header>
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
                        Posted by {thread.authorid}
                    </div>
                </div>
            </div>
        ))}
        <Link to = "/">
              <button className="footerbutton">
                Back to Homepage
              </button>
            </Link>
     </div>
    )
}

/*
fetch(`http://localhost:3000/${id}`, {
            method: "GET"
        })
        .then((response) => {
        if (response.ok) {
            console.log("Response:" + response)
            return response.json();
        }
        throw new Error("Fetch Error");
        })
        .then((categorydata) => {
            console.log("Data1:" + categorydata)
            setCategory(categorydata);
            console.log("Final1:" + JSON.stringify(category))
        });
*/