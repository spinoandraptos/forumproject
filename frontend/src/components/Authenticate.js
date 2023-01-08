import { createContext, useState } from "react";

  export const AuthContext = createContext();
 
  export const AuthContextProvider = ({ children }) => {
    const [flag,setFlag] = useState(false)
    const [userid, setUserid] = useState("");

    const Fetchusername = () => {
    if (localStorage.getItem("jwt")) {
        setFlag(true)
        
      fetch(`https://bopfishforum2.onrender.com/api/users/${JSON.parse(localStorage.getItem("jwt"))}/id`, {
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
        setUserid(userdata)
      })
    }
  }
 
    return (
      <AuthContext.Provider value={{ flag, userid, Fetchusername}}>
        {children}
      </AuthContext.Provider>
    );
  };
