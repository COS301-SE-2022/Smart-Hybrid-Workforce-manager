import { BrowserRouter as Router, Routes, Route, BrowserRouter } from 'react-router-dom'
import React, {useState} from 'react';
import Home from './pages/Home'
import Login from './pages/Login'
import Signup from './pages/Signup'

import Bookings from './pages/Bookings'
import BookingsDesk from './pages/BookingsDesk'
import BookingsDeskEdit from './pages/BookingsDeskEdit'
import BookingsMeeting from './pages/BookingsMeeting'

import Admin from './pages/Admin'

import Teams from './pages/Teams'
import CreateTeam from './pages/TeamsCreate'
import EditTeam from './pages/TeamsEdit'
import PermissionsTeam from './pages/TeamsPermissions'

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
import Users from './pages/Users'
import EditUser from './pages/UsersEdit'
import CreateUser from './pages/UsersCreate'

import Roles from './pages/Roles'
import CreateRole from './pages/RolesCreate'
import EditRole from './pages/RolesEdit'

import userContext from './store/userContext';
import ProtectedPage from './store/protectedPage';
import { UserProvidor } from './store/userContext';

function App()
{
  // const [isAuthenticated, userHasAuthenticated] = useState(true);
  const [userData, setUserData] = useState({});
  // if(sessionStorage.getItem("userData")){
  //   console.log("Set userdata");
  //   console.log(sessionStorage.getItem("userData").isLoggedIn);
  //   setUserData(sessionStorage.getItem("userData"));
  // }
  if(!userData || !userData.isLoggedIn)
    return <Login setUserData={setUserData}/>

  return(
    // <Home />
    // <BrowserRouter>
    //   <Route path="/test1" exact element={<Login/>} />
    // </BrowserRouter>
    <Router>
      <Routes>      
        <Route path="/test2" exact element={<Login/>} />
        <Route path="/" exact element={<Home/>} />
        <Route path="/login" exact element={<Login/>} />
        <Route path="/signup" exact element={<Signup/>} />
        <Route path="/bookings" exact element={<Bookings/>} />
        <Route path="/bookings-desk" exact element={<BookingsDesk/>} />
        <Route path="/bookings-meeting" exact element={<BookingsMeeting />} />
        <Route path="/admin" exact element={<Admin />} />

        <Route path="/users" exact element={<Users />} />
        <Route path="/user-edit" exact element={<EditUser />} />
        <Route path="/user-create" exact element={<CreateUser />} />
        
        <Route path="/team" exact element={<Teams />} />
        <Route path="/team-create" exact element={<CreateTeam />} />
        <Route path="/team-edit" exact element={<EditTeam />} />
        <Route path="/team-permissions" exact element={<PermissionsTeam />} />

        <Route path="/bookings-desk-edit" exact element={<BookingsDeskEdit/>} />
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

        <Route path="/role" exact element={<Roles/>} />
        <Route path="/role-create" exact element={<CreateRole/>} />
        <Route path="/role-edit" exact element={<EditRole/>} />
      </Routes>
    </Router>
  );
}

export default App;
