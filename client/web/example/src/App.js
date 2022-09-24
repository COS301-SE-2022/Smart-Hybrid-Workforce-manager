import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'

import Login from './pages/Login'
import Signup from './pages/Signup'

import Home from './pages/Home'

import BookingsDesk from './pages/BookingsDesk'
import BookingsMeetingRoom from './pages/BookingsMeetingRoom'

import Calendar from './pages/Calendar'

import Admin from './pages/Admin'

import Teams from './pages/Teams'
import CreateTeam from './pages/TeamsCreate'
import EditTeam from './pages/TeamsEdit'
import PermissionsTeam from './pages/TeamsPermissions'

import Layout from './pages/CreateLayout'
import Resources from './pages/Resources'
import CreateBuilding from './pages/ResourcesBuildingCreate'
import EditBuilding from './pages/ResourcesBuildingEdit'
import CreateRoom from './pages/ResourcesRoomCreate'
import EditRoom from './pages/ResourcesRoomEdit'
import CreateDesk from './pages/ResourcesDeskCreate'
import EditDesk from './pages/ResourcesDeskEdit'
import EditMeetingRoom from './pages/ResourcesMeetingRoomEdit'
import CreateMeetingRoom from './pages/ResourcesMeetingRoomCreate'

import Profile from './pages/Profile'
import ProfileConfiguration from './pages/ProfileConfiguration'
import ResetPassword from './pages/ResetPassword'
import Users from './pages/Users'
import EditUser from './pages/UsersEdit'
import CreateUser from './pages/UsersCreate'
import PermissionsUser from './pages/UsersPermissions'

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

            <Route path="/users" exact element={<Users />} />
            <Route path="/user-edit" exact element={<EditUser />} />
            <Route path="/user-create" exact element={<CreateUser />} />
            <Route path="/user-permissions" exact element={<PermissionsUser />} />
            
            <Route path="/team" exact element={<Teams />} />
            <Route path="/team-create" exact element={<CreateTeam />} />
            <Route path="/team-edit" exact element={<EditTeam />} />
            <Route path="/team-permissions" exact element={<PermissionsTeam />} />

            <Route path="/layout" exact element={<Layout />} />
            <Route path="/resources" exact element={<Resources/>} />
            <Route path="/building" exact element={<CreateBuilding/>} />
            <Route path="/building-edit" exact element={<EditBuilding/>} />
            <Route path="/room" exact element={<CreateRoom/>} />
            <Route path="/room-edit" exact element={<EditRoom/>} />
            <Route path="/desk" exact element={<CreateDesk/>} />
            <Route path="/resources-desk-edit" exact element={<EditDesk/>} />
            <Route path="/resources-meeting-room-edit" exact element={<EditMeetingRoom/>} />
            <Route path="/meetingroom" exact element={<CreateMeetingRoom />} />

            <Route path="/profile" exact element={<Profile />} />
            <Route path="/profile-configuration" exact element={<ProfileConfiguration />} />
            <Route path="/reset-password" exact element={<ResetPassword />} />

            <Route path="/role" exact element={<Roles/>} />
            <Route path="/role-create" exact element={<CreateRole/>} />
            <Route path="/role-edit" exact element={<EditRole />} />
            <Route path="/role-permissions" exact element={<PermissionsRole />} /> 

            <Route path="/admin" exact element={<Admin />} />
          </Route>          
          <Route path="/login" exact element={<Login />} />
          <Route path="/signup" exact element={<Signup/>} />
        </Routes>          
      </UserContext.Provider>          
    </Router>
  );
}

export default App;
