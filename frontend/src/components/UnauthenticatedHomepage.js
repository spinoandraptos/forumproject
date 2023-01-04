import React from "react"
import { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import "./Forum.css";

//define the function homepage that will serve as the component rendering the homepage upon entry
//first create the category state and state update function using the state hook useState
//category data will be stored in the local state, which is initialised to be an empty array
//then using the useEffect hook, fetch category data from API and update the state to contain the data
//finally, render the data to be a visible format and output it to client

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
      fetch(`http://localhost:3000/search`, {
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
        fetch("http://localhost:3000/", {
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