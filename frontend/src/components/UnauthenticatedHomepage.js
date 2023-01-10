// this will serve as the component rendering the unauthenticated homepage
// we will create states to manage the data of retrieved categories, searchbar input, and thread retrieevd from search
// upon render, the data for categories (displayed on screen) is fetched via fetch API from Go backend 
// additionally, upon input in searchbar fetch api will attempt to retrieve the corresponding thread 
// link and navigate allow users to navigate to other components

import React from "react"
import { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import "./Forum.css";

export default function NHomepage() {

    const [categories, setCategories] = useState([]);
    const [thread, setThread] = useState({});
    const [search, setSearch] = useState("");
    const navigate = useNavigate();

    const handleSearch = (input) => {
      setSearch(input.target.value);
    }

    function postdata(input){
      input.preventDefault();
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
        return (response.json())
      }
    })
    .then((threadinfo) => {
      setThread(threadinfo)
    })
    }, 500);

    return () => clearTimeout(delay);
    }, [search])

    useEffect(() => {
        fetch("https://bopfishforum2.onrender.com/api/categories", {
            method: "GET",
            credentials:"include"
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