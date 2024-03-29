// this will serve as the component rendering the authenticated homepage upon login
// we will create states to manage the data of retrieved categories, current user, searchbar input, thread retrieevd from search, password input for user verification, and the status of modal
// upon render, the data for categories (displayed on screen) and current user (username displayed) are fetched via fetch API from Go backend 
// additionally, upon input in searchbar fetch api will attempt to retrieve the corresponding thread 
// upon choosing to edit user info a modal will open to prompt for user verification, and if successful, direct user to page to edit user info
// link and navigate allow users to navigate to other components

import React from "react"
import "bootstrap/dist/css/bootstrap.min.css";
import { useState, useEffect, useContext } from "react";
import { AuthContext } from "./Authenticate";
import { Modal } from "react-bootstrap"
import { Link, useNavigate } from "react-router-dom";
import "./Forum.css";

export default function AHomepage() {

    const [categories, setCategories] = useState([]);
    const {Fetchusername} = useContext(AuthContext);
    const [user, setUser] = useState({});
    const [thread, setThread] = useState({});
    const [search, setSearch] = useState("");
    const [password, setPassword] = useState("");
    const [modalOpen, setModalOpen] = useState(false);
    const navigate = useNavigate();

    const handlePassword = (input) => {
      setPassword(input.target.value);
    }

    const handleSearch = (input) => {
      setSearch(input.target.value);
    }

    function Showmodal(){
      setModalOpen(true)
    }

    function Closemodal(){
      setModalOpen(false)
    }

    function Redirecteditpage(input){
      input.preventDefault();
      if(password===user.password) {
        setModalOpen(false)
        navigate(`/users/${user.id}`)
      } else {
        setModalOpen(false)
        alert("Incorrect Password!")
      }
    }

    function postdata(){
      console.log(thread)
      if(thread.id && thread.categoryid){
        navigate(`/${thread.categoryid}/threads/${thread.id}/comments`)
      }
      else {
        alert("Error: Thread Cannot be Found (Remember to Input Full Title)")
      }
    }

    useEffect(()=>{
      const delay = setTimeout(() => {
      fetch(`https://bopfishforum2.onrender.com/api/search`, {
      method: "POST",
      credentials: "include",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        Title: search,
      })
    })
    .then((response) => {
      if (response.ok) {
        console.log("Response:" + response)
        return (response.json())
      }
    })
    .then((threadinfo) => {
      setThread(threadinfo)
    })
    }, 500);

    return () => clearTimeout(delay);
    }, [search])

    useEffect(()=>
      Fetchusername(),
    [])

    useEffect(()=>{
      fetch(`https://bopfishforum2.onrender.com/api/users/${JSON.parse(localStorage.getItem("jwt"))}`, {
      method: "POST",
      credentials: "include",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        Username: JSON.parse(localStorage.getItem("jwt"))
      })
      })
      .then((response)=> {
        if (response.ok) {
          return response.json();
        }
      })
      .then((userdata)=> {
        setUser(userdata)
      })
    },[])

    useEffect(() => {
        fetch("https://bopfishforum2.onrender.com/api/categories", {
            method: "GET",
            credentials: "include",
        })
        .then((response) => {
        if (response.ok) {
            return response.json();
        }
        throw new Error("Fetch Error");
        })
        .then((categorydata) => {
            setCategories(categorydata);
        });
    }, []);

    return (
        <div className = "allcategories">
          <div className="modalcontent">
          <Modal size="lg" show={modalOpen} onHide={Closemodal}>
            <Modal.Header closeButton>
              <Modal.Title>
                User Verification
              </Modal.Title>
            </Modal.Header>

            <Modal.Body>
            <form className="form" onSubmit={Redirecteditpage}>
              <div className="verifylabel">
                <label htmlFor="password">
                  Password:
                </label>
              </div>
              <div className="verifybox">
                <input id="password" type="password" value={password} onChange={handlePassword}/>
                <br />
                <button className="formsubmitbutton" type="submit">
                  Submit
                </button>
              </div>
            </form>
            </Modal.Body>

            <Modal.Footer>
              <button className="footerbutton2" onClick={Closemodal}>Close</button>
            </Modal.Footer>
          </Modal>
          </div>
          <div className="herocontent">
            <div className="herotext">
              Bop Fish Nation 🦈
            </div>
          </div>
          <header id = "homepageheader">
            <div className = "headerlinks">
            <div className="userwelcome">
              Welcome, {user.username}!
            </div>
            <form onSubmit={postdata}>
              <input
                type="search"
                placeholder="Search Thread by Title"
                id="searchbar"
                size={25}
                value={search}
                onChange={handleSearch}
              />
              <button className="searchbutton">
                🔍
              </button>
            </form>
              <button className="headerbutton" onClick={Showmodal}>
                Edit User Info
              </button>
              <Link to = "/users/logout">
                <button className="headerbutton">
                  Logout
                </button>
              </Link>
            </div>
          </header>
          {categories?.map(category => (
            <div className = "indivcategory" key={category.id}>
              <div className="categorytitle">
                {category.title}
              </div>
              <div className="categorycontent">
                {category.description}
              </div>
              <Link to = {`/${category.id}`}>
                <button className="categorybutton">
                  Check out the threads in this category!
                </button>
              </Link>
            </div>
          ))}
        </div>
      );
}