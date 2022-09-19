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
        localStorage.removeItem("auth_data");
        console.log(userData);
        navigate("/login");
    }

    return (
        <Button className='button-user-profile' variant='primary' onClick={handleSubmit} data-testid='button-user-profile'>Log Out</Button>
    )
}

export default LogoutButton;