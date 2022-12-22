import React from "react"
import { Link } from "react-router-dom";

export default function Threadpage() {
    return (
      <div>
        <h1>All Comments Page</h1>
        <p>Check out all the Comments under this category!</p>
          <Link to = "/">
            Return to mainpage
          </Link>
      </div>
    )
}