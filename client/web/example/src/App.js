import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'

import Login from './pages/Login'
import Signup from './pages/Signup'

import Home from './pages/Home'

import BookingsDesk from './pages/BookingsDesk'
import BookingsMeetingRoom from './pages/BookingsMeetingRoom'

import Calendar from './pages/Calendar'

import Admin from './pages/Admin'

import OfficeCreator from './pages/OfficeCreator'

import Statistics from './pages/Statistics'

import Profile from './pages/Profile'
import ResetPassword from './pages/ResetPassword'

import Roles from './pages/Roles'
import CreateRole from './pages/RolesCreate'
import EditRole from './pages/RolesEdit'

import PermissionsRole from './pages/RolesPermissions'

import React, { useEffect } from 'react'
import { useState } from 'react'
import ProtectedRoute from './store/ProtectedRoute'


export const UserContext = React.createContext();


function App()
{
  const [userData, setUserData] = useState(() => {
    const sessionData = localStorage.getItem("auth_data");
    try{
      const val = JSON.parse(sessionData);
      return val;
    }catch(error){
      return sessionData;
    }
  });
  useEffect(() => {
    const stringVal = JSON.stringify(userData);
    localStorage.setItem("auth_data",stringVal);
  },[userData]);

  return(
    <Router>
      <UserContext.Provider value={{userData, setUserData}}>
        <Routes>
          <Route element={<ProtectedRoute/>}>
            <Route path="/" exact element={<Home/>} />
            <Route path="/bookings-desk" exact element={<BookingsDesk/>} />
            <Route path="/bookings-meetingroom" exact element={<BookingsMeetingRoom/>} />
            <Route path="/calendar" exact element={<Calendar />} />

            <Route path="/statistics" exact element={<Statistics />} />

            <Route path="/profile" exact element={<Profile />} />
            <Route path="/reset-password" exact element={<ResetPassword />} />

            <Route path="/role" exact element={<Roles/>} />
            <Route path="/role-create" exact element={<CreateRole/>} />
            <Route path="/role-edit" exact element={<EditRole />} />
            <Route path="/role-permissions" exact element={<PermissionsRole />} /> 

            <Route path="/admin" exact element={<Admin />} />
            <Route path="/office-creator" exact element={<OfficeCreator />} />
          </Route>          
          <Route path="/login" exact element={<Login />} />
          <Route path="/signup" exact element={<Signup/>} />
        </Routes>          
      </UserContext.Provider>          
    </Router>
  );
}

export default App;
