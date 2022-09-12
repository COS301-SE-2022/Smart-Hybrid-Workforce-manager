import { Stage, Layer } from 'react-konva'
import { useRef, useState, useEffect, useCallback, useContext } from 'react'
import Desk from '../components/Map/Desk'
import MeetingRoom from '../components/Map/MeetingRoom'
import { FaSave, FaQuestion } from 'react-icons/fa'
import { MdEdit, MdAdd } from 'react-icons/md'
import desk_white from '../img/desk_white.svg';
import meetingroom_white from '../img/meetingroom_white.svg';
import { UserContext } from '../App'
import { useNavigate } from 'react-router-dom'
import ProfileBar from '../components/Navbar/ProfileBar.js';
import Navbar from '../components/Navbar/Navbar.js'

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
    const deletedResources = useRef([]);
    const propertiesPaneRef = useRef(true);
    const helpRef = useRef(null);
    const helpToolRef = useRef(null);

    //Pane states
    const [propertiesPaneLeft, SetPropertiesPaneLeft] = useState(0.85*window.innerWidth);

    //Desk and meeting room prop arrays
    const [deskProps, SetDeskProps] = useState([]);
    const [meetingRoomProps, SetMeetingRoomProps] = useState([]);
    const [stage, SetStage] = useState({width : 100, height : 100});
    const [selectedId, SelectShape] = useState(null);

    //API fetch variables
    const [buildings, SetBuildings] = useState([]);
    const [currBuilding, SetCurrBuilding] = useState("");
    const [rooms, SetRooms] = useState([]);
    const [currRoom, SetCurrRoom] = useState("");
    const [resources, SetResources] = useState([]);

    const {userData} = useContext(UserContext)
    const navigate = useNavigate();

    //POST requests
    const FetchBuildings = () =>
    {
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
    }

    /*useEffect(() =>
    {
        console.log("Resources: " + resources)
        console.log(deskProps);
        console.log(deskPropsRef.current);
    },[resources, deskProps])*/

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
                SetCurrRoom("");
                SetCurrBuilding(e.target.value);
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
                SetCurrRoom(e.target.value);
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

    //Collapse Properties pane
    const PropertiesCollapse = () =>
    {
        if(propertiesPaneRef.current)
        {
            SetPropertiesPaneLeft(0.985*window.innerWidth);
            propertiesPaneRef.current = false;
        }
        else
        {
            SetPropertiesPaneLeft(0.85*window.innerWidth);
            propertiesPaneRef.current = true;
        }
    }

    //Add building
    const AddBuilding = () =>
    {
        navigate("/building");
    }

    //Edit selected building
    let EditBuilding = async (e) =>
    {
        if(currBuilding !== "")
        {
            e.preventDefault();
            window.sessionStorage.setItem("BuildingID", currBuilding);
            navigate("/building-edit");
        }
        else
        {
            window.alert("Please select a building to edit");
        }
    }

    const AddRoom = () =>
    {
        if(currBuilding !== "")
        {
            window.sessionStorage.setItem("BuildingID", currBuilding);
            navigate("/room");
        }
        else
        {
            alert("Please select a building");
        }
    }

    let EditRoom = async (e) =>
    {
        if(currRoom !== "")
        {
            e.preventDefault();
            window.sessionStorage.setItem("RoomID", currRoom);
            window.sessionStorage.setItem("BuildingID", currBuilding);
            navigate("/room-edit");
        }
        else
        {
            window.alert("Please select a room to edit");
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

    //Add a new desk to the state
    const AddDesk = () =>
    {
        /*if(currBuilding === "" || currRoom === "")
        {
            window.alert("Please select a building and room");
            return;
        }*/

        if(stageRef.current !== null)
        {
            SetDeskProps(
            [
                ...deskProps,
                {
                    key : "desk" + deskCount.current,
                    id : null,
                    name : "Desk " + deskCount.current,
                    x : (-stageRef.current.x() + stageRef.current.width() / 2.0) / stageRef.current.scaleX(),
                    y : (-stageRef.current.y() + stageRef.current.height() / 2.0) / stageRef.current.scaleY(),
                    width : 60,
                    height : 55,
                    rotation : 0,
                    edited : false
                }
            ]);
        }
    };

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

    //Add a new desk to the state
    const AddMeetingRoom = () =>
    {
        /*if(currBuilding === "" || currRoom === "")
        {
            window.alert("Please select a building and room");
            return;
        }*/

        if(stageRef.current !== null)
        {
            SetMeetingRoomProps(
            [
                ...meetingRoomProps,
                {
                    key : "meetingroom" + meetingRoomCount.current,
                    id : null,
                    name : "Meeting Room " + meetingRoomCount.current,
                    x : (-stageRef.current.x() + stageRef.current.width() / 2.0) / stageRef.current.scaleX(),
                    y : (-stageRef.current.y() + stageRef.current.height() / 2.0) / stageRef.current.scaleY(),
                    width : 200,
                    height : 200,
                    rotation : 0,
                    edited : true
                }
            ]);
        }
    };

    //Function to monitor when a key is pressed. Returns true if target key is pressed and false when target key is released
    const deletePressed = useKeyPress("Delete");
    function useKeyPress(targetKey)
    {
        // State for keeping track of whether key is pressed
        const [keyPressed, SetKeyPressed] = useState(false);
        
        //Event listeners for key press
        useEffect(() =>
        {
            // If pressed key is our target key then set to true
            function downHandler({key})
            {
                if (key === targetKey)
                {
                    SetKeyPressed(true);
                }
            };

            // If released key is our target key then set to false
            function upHandler({key})
            {
                if (key === targetKey)
                {
                    SetKeyPressed(false);
                }
            };

            window.addEventListener("keydown", downHandler);
            window.addEventListener("keyup", upHandler);

            // Remove event listeners on cleanup
            return () => 
            {
                window.removeEventListener("keydown", downHandler);
                window.removeEventListener("keyup", upHandler);
            };
        }, [targetKey]);

        return keyPressed;
    };

    const HandleDelete = useCallback(() =>
    {
        if(selectedId !== null)
        {
            if(selectedId.includes("desk"))
            {
                for(var i = 0; i < deskProps.length; i++)
                {
                    if(deskProps[i].key === selectedId)
                    {
                        if(deskProps[i].id !== null)
                        {
                            deletedResources.current.push(deskProps[i]);
                        }

                        var newDesk = [...deskProps];
                        newDesk.splice(i, 1);
                        SetDeskProps(newDesk);
                        SelectShape(null);
                        break;
                    }
                }
            }
            else
            {
                for(i = 0; i < meetingRoomProps.length; i++)
                {
                    if(meetingRoomProps[i].key === selectedId)
                    {
                        if(meetingRoomProps[i].id !== null)
                        {
                            deletedResources.current.push(meetingRoomProps[i]);
                        }

                        var newMeetingRoom = [...meetingRoomProps];
                        newMeetingRoom.splice(i, 1);
                        SetMeetingRoomProps(newMeetingRoom);
                        SelectShape(null);
                        break;
                    }
                }
            }
        }
    }, [deskProps, meetingRoomProps, selectedId])

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

    //Saves the current layout to the database
    const SaveLayout = async () =>
    {
        if(currBuilding === "" || currRoom === "")
        {
            window.alert("Please select a building and room");
            return;
        }

        var resources = [];

        for(var i = 0; i < deskProps.length; i++)
        {
            var currDesk = deskProps[i];
            
            if(currDesk.edited)
            {
                resources.push(
                {
                    id : currDesk.id,
                    room_id: currRoom,
                    name: currDesk.name,
                    xcoord: currDesk.x,
                    ycoord: currDesk.y,
                    width: currDesk.width,
                    height: currDesk.height,
                    rotation: currDesk.rotation,
                    role_id: null,
                    resource_type: 'DESK',
                    decorations: '{"computer": true}',
                })
            }
        }

        for(i = 0; i < meetingRoomProps.length; i++)
        {
            var currMeetingRoom = meetingRoomProps[i];
            
            if(currMeetingRoom.edited)
            {
                resources.push(
                {
                    id : currMeetingRoom.id,
                    room_id: currRoom,
                    name: currMeetingRoom.name,
                    xcoord: currMeetingRoom.x,
                    ycoord: currMeetingRoom.y,
                    width: currMeetingRoom.width,
                    height: currMeetingRoom.height,
                    rotation: currMeetingRoom.rotation,
                    role_id: null,
                    resource_type: 'MEETINGROOM',
                    decorations: '{}',
                })
            }
        }

        try
        {
            let res = await fetch("http://localhost:8080/api/resource/batch-create", 
            {
                method: "POST",
                body: JSON.stringify(resources)
            });

            if(res.status === 200)
            {
                alert("Saved!");
            }
        }
        catch(err)
        {
            console.log(err);
        }

        if(deletedResources.current.length > 0)
        {
            console.log(deletedResources.current + currBuilding);
        }
    }

    const ViewHelp = () =>
    {
        helpRef.current.style.visibility = 'visible';
        helpToolRef.current.style.visibility = 'visible';
    }

    const CloseHelp = () =>
    {
        helpRef.current.style.visibility = 'hidden';
        helpToolRef.current.style.visibility = 'hidden';
    }

    //Effect on the loading of the web page
    useEffect(() =>
    {
        SetStage({width : canvasRef.current.offsetWidth, height : canvasRef.current.offsetHeight});
        FetchBuildings();
    },[]);

    //Effect to monitor if delete key is pressed
    
    useEffect(() =>
    {
        if(deletePressed)
        {
            console.log("D")
            HandleDelete();
        }
    }, [deletePressed, HandleDelete]);

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


    return (
        <div className='page-container'>
            <div className='content'>
                <ProfileBar />
                <Navbar />

                <div className='main-container'>
                    <div className='properties-pane' style={{left: propertiesPaneLeft}}>
                        <div className='properties-pane-label-container' onClick={PropertiesCollapse} >
                            <p>Properties</p>
                        </div>

                        <div className='building-pane'>
                            <p className='building-label'>Buildings</p>
                            <MdAdd className='add-building-img' size={35} onClick={AddBuilding} />
                            <MdEdit className='edit-building-img' size={25} onClick={EditBuilding} />

                                <select className='list-box-building' name='building' size='10' onChange={UpdateRooms.bind(this)}>
                                    <option value='' disabled selected id='BuildingDefault'>--Select the building--</option>
                                    {buildings.length > 0 && (
                                        buildings.map(building => (
                                            <option value={building.id}>{building.name + ' (' + building.location + ')'}</option>
                                        ))
                                    )}
                                </select>
                        </div>

                        <div className='room-pane'>
                            <p className='room-label'>Rooms</p>
                            <MdAdd className='add-room-img' size={35} onClick={AddRoom} />
                            <MdEdit className='edit-room-img' size={25} onClick={EditRoom} />

                                <select className='list-box-room' name='room' size="10" onChange={UpdateResources.bind(this)}>
                                    <option value='' disabled selected id='RoomDefault'>--Select the room--</option>
                                    {rooms.length > 0 && (
                                        rooms.map(room => (
                                            <option value={room.id}>{room.name + ' (' + room.location + ')'}</option>
                                        ))
                                    )}
                                </select>
                        </div>
                    </div>                                     

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
                </div>
            </div>  
        </div>
    )
}

export default Home