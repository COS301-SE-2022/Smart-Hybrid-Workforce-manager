import React, { useContext } from 'react'
import { UserContext } from '../../App'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useLocation, useNavigate } from 'react-router-dom'

const LogoutButton = () =>{
    const {userData,setUserData} = useContext(UserContext);
    const navigate = useNavigate();
    const location = useLocation();

    let handleSubmit = async(e) =>{
        e.preventDefault();
        setUserData(null);
        console.log(userData);
        navigate("/login");
    }

    return (
        <Form className='form' onSubmit={handleSubmit}>
            <Button className='button-submit' variant='primary' type='submit'></Button>
        </Form>
    )
}

export default LogoutButton;