import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import { useState, useEffect } from 'react';
import { MdDesktopWindows } from 'react-icons/md'
import Button from 'react-bootstrap/Button'

const Resources = () =>
{
  const [buildings, SetBuildings] = useState([])
  const [currBuilidng, SetCurrBuilding] = useState("")
  const [rooms, SetRooms] = useState([])
  const [currRoom, SetCurrRoom] = useState("")
  const [resources, SetResources] = useState([])

  //POST request
  const FetchBuildings = () =>
  {
    fetch("http://localhost:8100/api/resource/building/information", 
        {
          method: "POST",
          body: JSON.stringify({
          })
        }).then((res) => res.json()).then(data => 
          {
            SetBuildings(data);
          });
  }

  const UpdateRooms = (e) =>
  {
    fetch("http://localhost:8100/api/resource/room/information", 
        {
          method: "POST",
          body: JSON.stringify({
            building_id: e.target.value
          })
        }).then((res) => res.json()).then(data => 
          {
            SetRooms(data);
            SetCurrBuilding(e.target.value);
          });
  }

  const UpdateResources = (e) =>
  {
    fetch("http://localhost:8100/api/resource/information", 
        {
          method: "POST",
          body: JSON.stringify({
            room_id: e.target.value
          })
        }).then((res) => res.json()).then(data => 
          {
            SetResources(data);
            SetCurrRoom(e.target.value);
          });
  }

  const AddBuilding = () =>
  {
    window.location.assign("./building");
  }

  const AddRoom = () =>
  {
    if(currBuilidng !== "")
    {
      window.sessionStorage.setItem("BuildingID", currBuilidng);
      window.location.assign("./room");
    }
    else
    {
      alert("Please select a building");
    }
  }

  const AddDesk = () =>
  {
    if(currRoom !== "")
    {
      window.sessionStorage.setItem("RoomID", currRoom);
      window.location.assign("./desk");
    }
    else
    {
      alert("Please select a room");
    }
  }

  //Using useEffect hook. This will send the POST request once the component is mounted
  useEffect(() =>
  {
    FetchBuildings()
    window.sessionStorage.removeItem("BuildingID");
    window.sessionStorage.removeItem("RoomID");
  }, [])
  

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='combo-grid'>
          <div className='building-container'>
            <select className='combo-box' name='building' onChange={UpdateRooms.bind(this)}>
              <option value='' disabled selected>--Select the building--</option>
              {buildings.length > 0 && (
                buildings.map(building => (
                  <option value={building.id}>{building.name + ' (' + building.location + ')'}</option>
                ))
              )}
            </select>

            <Button className='button-resource' variant='primary' onClick={AddBuilding}>Add Building</Button>
          </div>

          <div className='room-container'>
            <select className='combo-box' name='room' onChange={UpdateResources.bind(this)}>
              <option value='' disabled selected>--Select the room--</option>
              {rooms.length > 0 && (
                rooms.map(room => (
                  <option value={room.id}>{room.name + ' (' + room.location + ')'}</option>
                ))
              )}
            </select>

            <Button className='button-resource' variant='primary' onClick={AddRoom}>Add Room</Button>
          </div>
        </div>

        <div className='resources-map'>
            {resources.length > 0 && (
              resources.map(resource => (
                <div className='resource-container'>
                  <MdDesktopWindows className='resource' size={50}/>
                </div>
              ))
            )}
        </div>

        <div className='button-resource-container'>
          <Button className='button-resource' variant='primary' onClick={AddDesk}>Add Desk</Button>
        </div>

      </div>  
      <Footer />
    </div>
  )
}

export default Resources