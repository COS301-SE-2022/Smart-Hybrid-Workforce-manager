//Source code provided by Lester Fernandez
//https://github.com/lesterfernandez
import { Navigate, Outlet, useLocation } from "react-router-dom";
import { useState } from "react";
import { useContext } from "react";
import { UserContext } from "../App";
// 

const ProtectedRoute = () => {
    const { userData } = useContext(UserContext);    
    const location = useLocation();
    console.log(userData);
    console.log((userData==null));
    if(userData==null){
        
    }
    
    return (userData!=null)?<Outlet/>:<Navigate to="/login" replace state={{from: location}}/>
}

export default ProtectedRoute;