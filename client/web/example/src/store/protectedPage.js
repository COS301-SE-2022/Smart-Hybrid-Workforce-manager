import React, { Component, useContext } from "react";
import { Route, Redirect, useLocation } from "react-router-dom";
import userProvidor from "./userContext";
import { userContext } from "./userContext";
import { useUser } from "./userContext.js";
import { Navigate } from "react-router-dom";

export default function ProtectedPage({children}){
    const {token, setToken} = useUser();

    const isLoggedIn = token;
    return(
        <>
            {!isLoggedIn ? (
                <Navigate to="/login"/>
            ):({children})}
        </>
    )
}

// export class ProtectedRoute extends Route{
//     render(){
//         const { token } = useContext(userContext);
//         if (token === "")
//             return <Redirect to='/login'/>;
//         return super.render()
//     }
// }
