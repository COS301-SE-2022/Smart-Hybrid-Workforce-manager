import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import { useState, useEffect } from 'react';
import Button from 'react-bootstrap/Button'
import ResourceDesk from '../components/Resource/ResourceDesk';
import ResourceMeetingRoom from '../components/Resource/ResourceMeetingRoom';

const Resources = () =>
{
  const [buildings, SetBuildings] = useState([])
  const [currBuilding, SetCurrBuilding] = useState("")
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
            document.getElementById("RoomDefault").selected = true;
            SetCurrRoom("");
            SetCurrBuilding(e.target.value);
            SetResources([]);
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

  let EditBuilding = async (e) =>
  {
      e.preventDefault();
      window.sessionStorage.setItem("BuildingID", currBuilding);
      window.location.assign("./building-edit");
  }

  const AddRoom = () =>
  {
    if(currBuilding !== "")
    {
      window.sessionStorage.setItem("BuildingID", currBuilding);
      window.location.assign("./room");
    }
    else
    {
      alert("Please select a building");
    }
  }

  let EditRoom = async (e) =>
  {
      e.preventDefault();
      window.sessionStorage.setItem("RoomID", currRoom);
      window.sessionStorage.setItem("BuildingID", currBuilding);
      window.location.assign("./room-edit");
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

  const AddMeetingRoom = () =>
  {
    if(currRoom !== "")
    {
      window.sessionStorage.setItem("RoomID", currRoom);
      window.location.assign("./meetingroom");
    }
    else
    {
      alert("Please select a room");
    }
  }

  //Using useEffect hook. This will send the POST request once the component is mounted
  useEffect(() =>
  {
    FetchBuildings();
    document.getElementById("RoomDefault").selected = true;
    document.getElementById("BuildingDefault").selected = true;
    window.sessionStorage.removeItem("BuildingID");
    window.sessionStorage.removeItem("RoomID");
    window.sessionStorage.removeItem("DeskID");
    window.sessionStorage.removeItem("DeskLocation");
    window.sessionStorage.removeItem("DeskName");
  }, [])
  

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='combo-grid'>
          <div className='building-container'>
            <select className='combo-box' name='building' onChange={UpdateRooms.bind(this)}>
              <option value='' disabled selected id='BuildingDefault'>--Select the building--</option>
              {buildings.length > 0 && (
                buildings.map(building => (
                  <option value={building.id}>{building.name + ' (' + building.location + ')'}</option>
                ))
              )}
            </select>

            <Button className='button-resource' variant='primary' onClick={AddBuilding}>Add Building</Button>
            <Button className='button-resource' variant='primary' onClick={EditBuilding}>Edit Building</Button>
          </div>

          <div className='room-container'>
            <select className='combo-box' name='room' onChange={UpdateResources.bind(this)}>
              <option value='' disabled selected id='RoomDefault'>--Select the room--</option>
              {rooms.length > 0 && (
                rooms.map(room => (
                  <option value={room.id}>{room.name + ' (' + room.location + ')'}</option>
                ))
              )}
            </select>

            <Button className='button-resource' variant='primary' onClick={AddRoom}>Add Room</Button>
            <Button className='button-resource' variant='primary' onClick={EditRoom}>Edit Room</Button>
          </div>
        </div>

        <div className='resources-map'>
          {resources.length > 0 && (
            resources.map(resource => 
            {
              if (resource.resource_type === "DESK")
                return <ResourceDesk id={resource.id} name={resource.name} location={resource.location} roomId={resource.room_id}/>
              return <ResourceMeetingRoom id={resource.id} name={resource.name} location={resource.location} capacity={resource.capacity} roomId={resource.room_id}/>
            }
          )
          )}
        </div>

        <div className='button-resource-container'>
          <Button className='button-resource' variant='primary' onClick={AddDesk}>Add Desk</Button>
          <Button className='button-resource' variant='primary' onClick={AddMeetingRoom}>Add Meeting Room</Button>
        </div>

      </div>  
      <Footer />
    </div>
  )
}

export default Resources