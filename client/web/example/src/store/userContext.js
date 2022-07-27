import React, { useContext, useEffect, useState } from 'react'

const userContext = React.createContext();
export const useUser = () => {useContext(userContext)}


export function UserProvidor({children}){
    // const [UserData, setUserData] = useState(null);
    // setUserData()


    // const login = (auth_data) => {
    //     setIsLoggedIn(true);
    //     setToken(auth_data["token"]);
    //     setFirstName(auth_data["first_name"]);
    //     setLastName(auth_data["last_name"]);
    //     setEmail(auth_data["email"]);
    //     setExpr(auth_data["expr_time"]);
    // }
    

    // const logout = () => {
    //     setIsLoggedIn(false);
    //     setToken(auth_data["token"]);
    //     setFirstName(auth_data["first_name"]);
    //     setLastName(auth_data["last_name"]);
    //     setEmail(auth_data["email"]);
    //     setExpr(auth_data["expr_time"]);
    // }

    // useEffect(() => {
    //     console.log("Updated State", isLoggedIn);
    //     if(isLoggedIn)
    //         window.location.assign("./");
    // }, [isLoggedIn])

    return (
        <userContext.Provider value={{}}>
            {children}
        </userContext.Provider>
    )
}

export default userContext;


