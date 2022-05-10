import Navbar from "./components/Navbar"
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Home from './pages/Home'
import Bookings from './pages/Bookings'
import Teams from './pages/Teams'
import Meetings from './pages/Meetings'

function App()
{
  return(
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" exact element={<Home/>} />
        <Route path="/bookings" exact element={<Bookings/>} />
        <Route path="/teams" exact element={<Teams/>} />
        <Route path="/meetings" exact element={<Meetings/>} />
      </Routes>
    </Router>
  );
}

export default App;
