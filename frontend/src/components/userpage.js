import React from "react"

//create a React class component called Header by extending the Component class in the React library 
//render() then returns the JSX that is being rendered on the screen
//JSX is a XML like syntax extension to JavaScript, it shares similar syntax to HTML
//JSX is converted into plain JavaScript code that all browser can understand through Babel compiler
class HomePage extends React.Component {
  render() {
    return (
      <div>
        <h1>Hello from Create React App</h1>
        <p>I am in a React Component!</p>
      </div>
    )
  }
}
export default HomePage