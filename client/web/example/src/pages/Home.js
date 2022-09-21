import { Stage, Layer } from 'react-konva'
import { useRef, useState, useEffect, useCallback, useContext } from 'react'
import Desk from '../components/Map/Desk'
import MeetingRoom from '../components/Map/MeetingRoom'
import { UserContext } from '../App'
import ProfileBar from '../components/Navbar/ProfileBar.js';
import Navbar from '../components/Navbar/Navbar.js'
import NavbarAdmin from '../components/Navbar/NavbarAdmin'

const Home = () =>
{
    //Canvas references
    const canvasRef = useRef(null);
    const stageRef = useRef(null);
    const scaleFactor = 1.3;
    const deskPropsRef = useRef([]);
    const meetingRoomPropsRef = useRef([]);
    const deskCount = useRef(0);
    const meetingRoomCount = useRef(0);

    //Selector refs
    const buildingRef = useRef(null);
    const roomRef = useRef(null);

    //Desk and meeting room prop arrays
    const [deskProps, SetDeskProps] = useState([]);
    const [meetingRoomProps, SetMeetingRoomProps] = useState([]);
    const [stage, SetStage] = useState({width : 100, height : 100});
    const [selectedId, SelectShape] = useState(null);

    //API fetch variables
    const [buildings, SetBuildings] = useState([]);
    const [rooms, SetRooms] = useState([]);
    const [resources, SetResources] = useState([]);

    const {userData} = useContext(UserContext);

    //POST requests
    const UpdateRooms = (e) =>
    {
        fetch("http://localhost:8080/api/resource/room/information", 
            {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                building_id: e.target.value
            }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
            }
            }).then((res) => res.json()).then(data => 
            {
                SetRooms(data);
                document.getElementById("RoomDefault").selected = true;
                SetResources([]);
            });
    }

    const UpdateResources = (e) =>
    {
        fetch("http://localhost:8080/api/resource/information", 
            {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                room_id: e.target.value
            }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
            }
            }).then((res) => res.json()).then(data => 
            {
                SetResources(data);
            });
    }

    //Canvas functions
    //Check if canvas is clicked and deselect the selected resource
    const CheckDeselect = (e) =>
    {
        const clickedEmpty = e.target === e.target.getStage();
        if(clickedEmpty)
        {
            SelectShape(null);
        }
    }

    //Load desks from the database
    const LoadDesk = useCallback((id, name, x, y, width, height, rotation) =>
    {
        //Uses a reference array to prevent state dependency and infinite loop
        if(stageRef.current !== null)
        {
            deskPropsRef.current =
            [
                ...deskPropsRef.current,
                {
                    key : "desk" + id,
                    id : id,
                    name : name,
                    x : x,
                    y : y,
                    width : width,
                    height : height,
                    rotation : rotation,
                    edited : false
                }
            ];

            //Set the state using the reference array
            SetDeskProps(deskPropsRef.current);
        }
    },[]);

    //Load desks from the database
    const LoadMeetingRoom = useCallback((id, name, x, y, width, height, rotation) =>
    {
        //Uses a reference array to prevent state dependency and infinite loop
        if(stageRef.current !== null)
        {
            meetingRoomPropsRef.current =
            [
                ...meetingRoomPropsRef.current,
                {
                    key : "meetingroom" + meetingRoomCount.current,
                    id : id,
                    name : name,
                    x : x,
                    y : y,
                    width : width,
                    height : height,
                    rotation : rotation,
                    edited : true
                }
            ];

            //Set the state using the reference array
            SetMeetingRoomProps(meetingRoomPropsRef.current);
        }
    },[]);

    //Adjusts the canvas size for difference screen sizes
    const HandleResize = () =>
    {
        SetStage({width : canvasRef.current.offsetWidth, height : canvasRef.current.offsetHeight});
    }

    window.addEventListener('resize', HandleResize);

    //Ensures that the zooming in/out is oriented with the center of viewable canvas
    const ZoomInOut = (event) =>
    {
        if(stageRef.current !== null)
        {
            const stage = stageRef.current;
            const oldScale = stage.scaleX();

            const stageCenter =
            {
                x : stage.width() / 2.0,
                y : stage.height() / 2.0
            }

            const newStageCenter = 
            {
                x : (stageCenter.x - stage.x()) / oldScale,
                y : (stageCenter.y - stage.y()) / oldScale,
            }

            var newScale;
            if(event.evt.deltaY < 0)
            {
                newScale = oldScale * scaleFactor;
            }
            else
            {
                newScale = oldScale / scaleFactor;
            }

            stage.scale({x : newScale, y : newScale});

            const newPos = 
            {
                x : stageCenter.x - newStageCenter.x * newScale,
                y : stageCenter.y - newStageCenter.y * newScale,
            }

            stage.position(newPos);
            stage.batchDraw();
        }        
    }

    //Effect on the loading of the web page
    useEffect(() =>
    {
        SetStage({width : canvasRef.current.offsetWidth, height : canvasRef.current.offsetHeight});

        fetch("http://localhost:8080/api/resource/building/information", 
        {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
            }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
            }
        }).then((res) => res.json()).then(data => 
        {
            SetBuildings(data);
        });

    }, [userData.token]);

    //Loads desks and meeting rooms from database after room is selected
    useEffect(() =>
    {
        //Reset reference array and counters
        deskPropsRef.current = [];
        deskCount.current = 0;
        meetingRoomPropsRef.current = [];
        meetingRoomCount.current = 0;

        SetDeskProps(deskPropsRef.current);
        SetMeetingRoomProps(meetingRoomPropsRef.current);

        //Loop through resources and load desks and meeting rooms respectively
        for(var i = 0; i < resources.length; i++)
        {
            if(resources[i].resource_type === "DESK")
            {
                LoadDesk(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation);
            }
            else if(resources[i].resource_type === "MEETINGROOM")
            {
                LoadMeetingRoom(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation);
            }
        }

    }, [resources, LoadDesk, LoadMeetingRoom]);

    //Update the desk counter when a new desk is added or removed
    useEffect(() =>
    {
        deskCount.current = deskProps.length;
    }, [deskProps.length]);

    //Update the meeting room counter when a new meeting room is added or removed
    useEffect(() =>
    {
        meetingRoomCount.current = meetingRoomProps.length;
    }, [meetingRoomProps.length]);

    useEffect(() =>
    {
        if(buildings.length > 0 && buildingRef.current)
        {
            buildingRef.current.value = buildings[0].id;
            fetch("http://localhost:8080/api/resource/room/information", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    building_id: buildings[0].id
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                }
            }).then((res) => res.json()).then(data => 
            {
                SetRooms(data);
                SetResources([]);
            });
        }
    }, [buildings, userData.token]);

    useEffect(() =>
    {
        if(rooms.length > 0 && roomRef.current)
        {
            roomRef.current.value = rooms[0].id;

            fetch("http://localhost:8080/api/resource/information", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    room_id: rooms[0].id
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}`
                }
            }).then((res) => res.json()).then(data => 
            {
                SetResources(data);
            });
        }
    }, [rooms, userData.token]);

    const showNavbar = () =>
    {
        if(!userData.user_identifier.includes("admin"))
        {
            return <Navbar />;
        }
        else
        {
            return <NavbarAdmin />;
        }
    };

    return (
        <div className='page-container'>
            <div className='content'>
                <ProfileBar />
                {showNavbar()}

                <div className='main-container'>
                    <div ref={canvasRef} className='canvas-container'>
                        <Stage width={stage.width} height={stage.height} onMouseDown={CheckDeselect} onTouchStart={CheckDeselect} draggable onWheel={ZoomInOut} ref={stageRef}>
                            <Layer>
                                {deskProps.length > 0 && (
                                    deskProps.map((desk, i) => (
                                        <Desk
                                            key = {desk.key}
                                            shapeProps = {desk}

                                            isSelected = {desk.key === selectedId}
                                            
                                            onSelect = {() => 
                                            {
                                                SelectShape(desk.key);
                                            }}
                                            
                                            onChange = {(newProps) => 
                                            {
                                                const newDeskProps = deskProps.slice();
                                                newDeskProps[i] = newProps;
                                                SetDeskProps(newDeskProps)
                                            }}

                                            draggable = {false}
                                        />
                                    ))
                                )}

                                {meetingRoomProps.length > 0 && (
                                    meetingRoomProps.map((meetingRoom, i) => (
                                        <MeetingRoom
                                            key = {meetingRoom.key}
                                            shapeProps = {meetingRoom}

                                            isSelected = {meetingRoom.key === selectedId}
                                            
                                            onSelect = {() => 
                                            {
                                                SelectShape(meetingRoom.key);
                                            }}
                                            
                                            onChange = {(newProps) => 
                                            {
                                                const newMeetingRoomProps = meetingRoomProps.slice();
                                                newMeetingRoomProps[i] = newProps;
                                                SetMeetingRoomProps(newMeetingRoomProps)
                                            }}

                                            draggable = {false}
                                        />
                                    ))
                                )}                             
                            </Layer>
                        </Stage>
                    </div>

                    <div className='building-selector-container'>
                        <select ref={buildingRef} className='building-selector' name='building' onChange={UpdateRooms.bind(this)}>
                            <option value='' disabled selected id='BuildingDefault'>--Select the building--</option>
                                {buildings.length > 0 && (
                                    buildings.map(building => (
                                        <option key={building.id} value={building.id}>{building.name + ' (' + building.location + ')'}</option>
                                    ))
                                )}
                        </select>
                    </div>

                    <div className='room-selector-container'>
                        <select ref={roomRef} className='room-selector' name='room' onChange={UpdateResources.bind(this)}>
                            <option value='' disabled selected id='RoomDefault'>--Select the room--</option>
                                {rooms.length > 0 && (
                                    rooms.map(room => (
                                        <option key={room.id} value={room.id}>{room.name + ' (' + room.location + ')'}</option>
                                    ))
                                )}
                        </select>
                    </div>
                </div>
            </div>  
        </div>
    )
}

export default Home