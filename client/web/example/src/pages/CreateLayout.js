import { Stage, Layer } from 'react-konva'
import { useRef, useState, useEffect, useCallback } from 'react'
import Desk from '../components/Map/Desk'
import MeetingRoom from '../components/Map/MeetingRoom'

const Layout = () =>
{
    //Canvas references
    const canvasRef = useRef(null);
    const stageRef = useRef(null);
    const scaleFactor = 1.3;
    const deskPropsRef = useRef([]);
    const meetingRoomPropsRef = useRef([]);
    const deskCount = useRef(0);
    const meetingRoomCount = useRef(0);
    const deletePressed = KeyPress("Delete");

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

    //POST requests
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
                    key : "desk" + deskCount.current,
                    id : id,
                    name : name,
                    x : x,
                    y : y,
                    width : width,
                    height : height,
                    rotation : rotation
                }
            ];

            //Set the state using the reference array
            SetDeskProps(deskPropsRef.current);
        }
    },[]);

    //Add a new desk to the state
    const AddDesk = () =>
    {
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
                    rotation : 0
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
                    rotation : rotation
                }
            ];

            //Set the state using the reference array
            SetMeetingRoomProps(meetingRoomPropsRef.current);
        }
    },[]);

    //Add a new desk to the state
    const AddMeetingRoom = () =>
    {
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
                    rotation : 0
                }
            ]);
        }
    };

    //Function to monitor when a key is pressed. Returns true if target key is pressed and false when target key is released
    const KeyPress = ((targetKey) =>
    {
        // State for keeping track of whether key is pressed
        const [keyPressed, SetKeyPressed] = useState(false);

        // If pressed key is our target key then set to true
        const downHandler = useCallback((key) =>
        {
            if (key === targetKey)
            {
                SetKeyPressed(true);
            }
        },[targetKey]);

        // If released key is our target key then set to false
        const upHandler = useCallback((key) =>
        {
            if (key === targetKey)
            {
                SetKeyPressed(false);
            }
        },[targetKey]);
        
        //Event listeners for key press
        useEffect(() =>
        {
            window.addEventListener("keydown", downHandler);
            window.addEventListener("keyup", upHandler);

            // Remove event listeners on cleanup
            return () => 
            {
                window.removeEventListener("keydown", downHandler);
                window.removeEventListener("keyup", upHandler);
            };
        }, [downHandler, upHandler]);

        return keyPressed;
    });

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
    const SaveLayout = () =>
    {
        window.alert("Building: " + currBuilding + "\nRoom: " + currRoom);
        console.log(resources);
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
            <div className='canvas-content'>
                <button onClick={AddDesk}>Add Desk</button><br></br>
                <button onClick={AddMeetingRoom}>Add Meeting Room</button>

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
                    </div>
                </div>                                          

                <div ref={canvasRef} className='canvas-container'>
                    <Stage width={stage.width} height={stage.height} onMouseDown={CheckDeselect} onTouchStart={CheckDeselect} draggable onDragEnd={canvasDrag} onWheel={ZoomInOut} ref={stageRef}>
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

                                        stage = {stageRef.current}
                                    />
                                ))
                            )}                             
                        </Layer>
                    </Stage>
                </div>
                <button onClick={SaveLayout}>Save</button>
            </div>  
        </div>
    )
}

export default Layout