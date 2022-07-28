import React, { useContext } from 'react'
import { UserContext } from '../../App'
import "../../App.css"
import Button from 'react-bootstrap/Button'
import { useNavigate } from 'react-router-dom'

const LogoutButton = () =>{
    const {userData,setUserData} = useContext(UserContext);
    const navigate = useNavigate();
    // const location = useLocation();

    let handleSubmit = async(e) =>{
        e.preventDefault();
        setUserData(null);
        console.log(userData);
        navigate("/login");
    }

    return (
        <Button className='button-user-profile' variant='primary' onClick={handleSubmit}>Log Out</Button>
    )
}

export default LogoutButton;