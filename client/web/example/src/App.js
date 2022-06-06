import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Home from './pages/Home'
import Login from './pages/Login'
import Signup from './pages/Signup'
import Bookings from './pages/Bookings'
import BookingsDesk from './pages/BookingsDesk'
import BookingsDeskEdit from './pages/BookingsDeskEdit'
import BookingsMeeting from './pages/BookingsMeeting'
import Teams from './pages/Teams'
import Meetings from './pages/Meetings'
import Resources from './pages/Resources'
import CreateBuilding from './pages/CreateBuilding'
import CreateRoom from './pages/CreateRoom'
import CreateDesk from './pages/CreateDesk'
import CreateMeetingRoom from './pages/CreateMeetingRoom'

function App()
{
  return(
    <Router>
      <Routes>
        <Route path="/" exact element={<Home/>} />
        <Route path="/login" exact element={<Login/>} />
        <Route path="/signup" exact element={<Signup/>} />
        <Route path="/bookings" exact element={<Bookings/>} />
        <Route path="/bookings-desk" exact element={<BookingsDesk/>} />
        <Route path="/bookings-meeting" exact element={<BookingsMeeting/>} />
        <Route path="/teams" exact element={<Teams/>} />
        <Route path="/meetings" exact element={<Meetings/>} />
        <Route path="/bookings-desk-edit" exact element={<BookingsDeskEdit/>} />
        <Route path="/resources" exact element={<Resources/>} />
        <Route path="/building" exact element={<CreateBuilding/>} />
        <Route path="/room" exact element={<CreateRoom/>} />
        <Route path="/desk" exact element={<CreateDesk/>} />
        <Route path="/meetingroom" exact element={<CreateMeetingRoom/>} />
      </Routes>
    </Router>
  );
}

export default App;
