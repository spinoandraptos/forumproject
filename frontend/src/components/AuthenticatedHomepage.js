import React from "react"
import { useState, useEffect, useContext } from "react";
import { AuthContext } from "./Authenticate";
import { Link } from "react-router-dom";
import "./Forum.css";
import "./Userpage"

//define the function homepage that will serve as the component rendering the homepage upon entry
//first create the category state and state update function using the state hook useState
//category data will be stored in the local state, which is initialised to be an empty array
//then using the useEffect hook, fetch category data from API and update the state to contain the data
//finally, render the data to be a visible format and output it to client

export default function AHomepage() {

    const [categories, setCategories] = useState([]);
    const {userid, Fetchusername} = useContext(AuthContext);
    const [user, setUser] = useState({})

    useEffect(()=>
      Fetchusername(),
    [])

    useEffect(()=>{
      fetch(`http://localhost:3000/users/${JSON.parse(localStorage.getItem("jwt"))}`, {
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
        fetch("http://localhost:3000/", {
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
            <Link to = {`/users/${userid}`}>
                <button className="headerbutton">
                  Edit User Info
                </button>
              </Link>
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