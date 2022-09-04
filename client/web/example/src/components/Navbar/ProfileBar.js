import React, { useContext, useEffect, useRef, useState } from 'react'
import { useNavigate } from "react-router-dom"
import { FaCalendar, FaTicketAlt, FaMap, FaChartPie, FaDoorOpen } from 'react-icons/fa'
import { CgProfile } from 'react-icons/cg'
import { UserContext } from '../../App';

const ProfileBar = (props, ref) =>
{
    const navigate = useNavigate();
    const profileRef = useRef(null);

    const {userData,setUserData} = useContext(UserContext);
    const [currLocation, setCurrLocation] = useState("");

    const NavigateProfile = () =>
    {
        navigate("/profile");
    }

    const Logout = () =>
    { 
        setUserData(null);
        localStorage.removeItem("auth_data");
        navigate("/login");
    }

    useEffect(() =>
    {
        if(currLocation === "/profile")
        {
            profileRef.current.style.color = "#09a2fb";
        }

    },[currLocation])

    useEffect(() =>
    {
        setCurrLocation(window.location.pathname);
    },[])

    return (
        <div className='profilebar-container'>
            <div ref={profileRef} className='profilepic-container' onClick={NavigateProfile}>
                <CgProfile />
            </div>

            <div className='logout' onClick={Logout}>
                <FaDoorOpen />
                &nbsp;
                Logout
            </div>
            
        </div>
    )
}

export default React.forwardRef(ProfileBar)