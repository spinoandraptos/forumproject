import React, { useContext, useEffect } from "react";
import { AuthContext } from "./Authenticate";
import { useNavigate }  from "react-router-dom";
import { useState } from "react";
import { Modal } from "react-bootstrap"
import "bootstrap/dist/css/bootstrap.min.css";

export default function Edituser() {
   
    const navigate = useNavigate();
    const [user, setUser] = useState({});
    const [modalOpen, setModalOpen] = useState(false);
    const [modal, setModal] = useState(0);
    const {flag, Fetchusername} = useContext(AuthContext);
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    function Closemodal(){
      setModalOpen(false)
    }

    function Usemodal1(input){
      input.preventDefault();
      setModalOpen(true)
      setModal(1)
    }

    function Usemodal2(input){
      input.preventDefault();
      setModalOpen(true)
      setModal(2)
    }

    const handleUsername = (input) => {
      setUsername(input.target.value);
    };

    const handlePassword = (input) => {
      setPassword(input.target.value);
    };
    
    useEffect(()=>
      Fetchusername(),
    [])

    useEffect(()=>{
      fetch(`http://localhost:3000/api/users/${JSON.parse(localStorage.getItem("jwt"))}`, {
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

    function Clickhomepage(){
        if (flag === true) {
          navigate("/authenticated")
        } else {
          navigate("/")
        }
      }

    function Deleteuser(input){

      input.preventDefault();

      fetch(`/api/users/${user.id}`, {
        method: "DELETE",
        credentials: "include",
        headers: { 'Content-Type': 'application/json' },
      })
      .then((response) => {
        if (response.ok) {
          Closemodal()
          localStorage.removeItem("jwt");
          alert("User Deleted Successfully!")
          navigate(`/`)
          window.location.reload() 
        } else if (response.status===401) {
          alert("Server Does Not Detect JWT")
        } else {
          alert("Error: User Cannot be Deleted")
        }
      })
    }

    function postdata(input){
        
      input.preventDefault();

        if (flag === true) {
        if(username === ""){
          fetch(`/api/users/${user.id}/password`, {
            method: "PUT",
            credentials: "include",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
              Password: password
            })
          })
          .then((response) => {
            if (response.ok) {
              Closemodal()
              console.log("Response:" + response)
              alert("Password Updated Successfully!")
              navigate(`/authenticated`)
            } else if (response.status===401) {
              alert("Server Does Not Detect JWT")
            } else {
              alert("Error: Password Cannot Update")
            }
          })
      } else if (password === ""){
        fetch(`/api/users/${user.id}/username`, {
          method: "PUT",
          credentials: "include",
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            Username: username
          })
        })
        .then((response) => {
          if (response.ok) {
            Closemodal()
            console.log("Response:" + response)
            alert("Username Updated Successfully!")
            localStorage.setItem("jwt", JSON.stringify(username))
            navigate(`/authenticated`)
          } else if (response.status===401) {
            alert("Server Does Not Detect JWT")
          } else {
            alert("Error: Username Cannot Update")
          }
        })
      } else {
        fetch(`/api/users/${user.id}`, {
          method: "PUT",
          credentials: "include",
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            Username: username,
            Password: password
          })
        })
        .then((response) => {
          if (response.ok) {
            console.log("Response:" + response)
            alert("User Updated Successfully!")
            localStorage.setItem("jwt", JSON.stringify(username))
            navigate(`/authenticated`)
          } else if (response.status===401) {
            alert("Server Does Not Detect JWT")
          } else {
            alert("Error: User Cannot Update")
          }
        })
      }
    } else {
        Closemodal()
        alert("Please Login First")
        navigate("/users/login")
      }
    }
    
    return (
        <div>
          
          <div className="modalcontent">
          <Modal size="lg" show={modalOpen} onHide={Closemodal}>
            {modal===1 && (
              <>
              <Modal.Header closeButton>
              <Modal.Title>
                Please Confirm
              </Modal.Title>
            </Modal.Header>
            <Modal.Body>
              <div className="warning">
                <div>
                  You are deleting your user account. Please double confirm.
                </div>
                <button onClick={Deleteuser} className="warningbutton">
                  Confirm
                </button>
              </div>
            </Modal.Body>
            <Modal.Footer>
              <button className="footerbutton2" onClick={Closemodal}>Abort Deletion</button>
            </Modal.Footer>
            </>
            )}
            {modal===2  && (
              <>
                <Modal.Header closeButton>
                  <Modal.Title>
                    Please Confirm Info Update
                  </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                  <div className="warning">
                    <button onClick={postdata} className="warningbutton2">
                      Confirm
                    </button>
                  </div>
                </Modal.Body>
                <Modal.Footer>
                  <button className="footerbutton2" onClick={Closemodal}>Abort Update</button>
                </Modal.Footer>
              </>
            )}
          </Modal>
          </div>

        <div className="herocontent">
          <div className="herotext">
            Bop Fish Nation 🦈 
          </div>
        </div>
        <header id = "homepageheader">
          <div className = "headerlinks">
              <button className="headerbutton" onClick={Clickhomepage}>
                Back to Homepage
              </button>
            <button className="headerbuttonwarning" onClick={Usemodal1}>
                Delete User
            </button>
          </div>
        </header>
        <div className="editcomment">
          <div className="userinfo">
            <div className="editusertitle">
                Update User Info
            </div>
            <form onSubmit={Usemodal2} className="editcommentform">
                <div className="loginbox commentbox">
                    <input id="username" placeholder="Update Username Here (Leave blank if no changes)" size={50} type="text" value={username} onChange={handleUsername} />
                    <br />
                    <input id="password" placeholder="Update Password Here (Leave blank if no changes)" size={50} type="password" value={password} onChange={handlePassword}/>
                    <br />
                    <button className="formsubmitbutton">
                       Submit
                    </button>
                </div>
            </form>
          </div>
        </div>
     </div>
    )
    
}