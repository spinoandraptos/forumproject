import { useState, useEffect } from "react";
import "./HomepageLayout.css";

//define the function homepage that will serve as the component rendering the homepage upon entry
//first create the category state and state update function using the state hook useState
//category data will be stored in the local state, which is initialised to be an empty array
//then using the useEffect hook, fetch category data from API and update the state to contain the data
//finally, render the data to be a visible format and output it to client

export default function Homepage() {

    const [categories, setCategories] = useState([]);

    useEffect(() => {
        fetch("http://localhost:3000/", {
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
            console.log("Data:" + JSON.stringify(categorydata))
            setCategories(categorydata);
            console.log("Final:" + categories)
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
              <a href = "/register">
                <button className="headerbutton">
                  Register
                </button>
              </a>
              <a href = "/login">
                <button className="headerbutton">
                  Login
                </button>
              </a>
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
          
            </div>
          ))}
        </div>
      );
}