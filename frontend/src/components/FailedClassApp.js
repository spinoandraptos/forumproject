import React, { Component } from 'react';
import './App.css';

export default class App extends Component {

//we first use a constructor to initialise the local state of category data (super is needed to set state)
//after the component is created, retrieve category data from database through fetch api
//the fetch response need to be resolved to JSON format which will update the state
//finally we render the retrieved data to be displayed in the browser
//we use console.log to check if the data are being retrieved and stored correctly

constructor(props) {
  super(props);
  this.state = {
    Categories: []
  };
}

  componentDidMount() {
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
        this.setState({ 
          Categories: categorydata
       });
    });
  }

  render() {
    const { categories } = this.state;
    console.log(this.state)

    return (
      <div className = "displaycategories">
            <h1> Categories </h1> 
            <ul className = "allcategories">
            {
              categories?.map(c => (
                <li key={c.id}>
                { c.title } - {c.description}
                </li>
               )) 
            }
            </ul>       
      </div>
    )
  }
}
