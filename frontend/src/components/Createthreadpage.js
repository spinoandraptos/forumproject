import React from "react";
import { Link }  from "react-router-dom";

export default function Createthread() {

    return(
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
        </div>
    )

}