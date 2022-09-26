import { Stage, Layer } from 'react-konva';
import { useRef, useState, useEffect, useCallback, useContext, Fragment } from 'react';
import Desk from './Desk';
import MeetingRoom from './MeetingRoom';
import { FaSave, FaQuestion } from 'react-icons/fa';
import { MdEdit, MdAdd } from 'react-icons/md';
import { BsThreeDotsVertical } from 'react-icons/bs';
import desk_white from '../../img/desk_white.svg';
import meetingroom_white from '../../img/meetingroom_white.svg';
import { UserContext } from '../../App';
import { useNavigate } from 'react-router-dom';
import styles from './map.module.css';

const Creator = () =>
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
    const helpRef = useRef(null);
    const helpToolRef = useRef(null);

    //Building and rooms
    const buildingMenuRef = useRef(null);
    const roomMenuRef = useRef(null);

    //Panel states
    const [propertiesPanel, setPropertiesPanel] = useState(0.985*window.innerWidth);
    const [resourceName, setResourceName] = useState('');
    const [resourceXCoord, setResourceXCoord] = useState('');
    const [resourceYCoord, setResourceYCoord] = useState('');

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
            mode: 'cors',
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
            mode: 'cors',
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

            if(buildingMenuRef.current && buildingMenuRef.current.style.display === 'block')
            {
                buildingMenuRef.current.style.display = 'none';
            }

            if(roomMenuRef.current && roomMenuRef.current.style.display === 'block')
            {
                roomMenuRef.current.style.display = 'none';
            }
        }
    }

    const ShowBuildingMenu = () =>
    {
        if(buildingMenuRef.current)
        {
            buildingMenuRef.current.style.display = 'block';
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

    const ShowRoomMenu = () =>
    {
        if(roomMenuRef.current)
        {
            roomMenuRef.current.style.display = 'block';
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

        if(selectedId)
        {
            setPropertiesPanel(0.65*window.innerWidth);
        }
        else
        {
            setPropertiesPanel(0.85*window.innerWidth);
        }
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
                    resource_type: 'MEETINGROOM',
                    decorations: `{"capacity": ${currMeetingRoom.capacity}}`,
                })
            }
        }

        try
        {
            let res = await fetch("http://localhost:8080/api/resource/batch-create", 
            {
                method: "POST",
                mode: 'cors',
                body: JSON.stringify(resources),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}`
                }
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

        fetch("http://localhost:8080/api/resource/building/information", 
        {
            method: "POST",
            mode: 'cors',
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
    },[userData.token]);

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

    //Check if properties are open or closed
    useEffect(() =>
    {
        if(selectedId)
        {
            setPropertiesPanel(0.65*window.innerWidth);
        }
        else
        {
            setPropertiesPanel(0.85*window.innerWidth);
        }
    },[selectedId, deskProps])

    const ChangeName = (name) =>
    {
        for(var i = 0; i < deskProps.length; i++)
        {
            if(deskProps[i].key === selectedId)
            {
                const newProps = deskProps.slice();
                newProps[i].name = name;
                SetDeskProps(newProps);
                setResourceName(name);
                break;
            }
        }
    }

    const ChangeXCoord = (x) =>
    {
        for(var i = 0; i < deskProps.length; i++)
        {
            if(deskProps[i].key === selectedId)
            {
                const newProps = deskProps.slice();
                newProps[i].x = parseInt(x);
                SetDeskProps(newProps);
                setResourceXCoord(x);
                break;
            }
        }
    }

    const ChangeYCoord = (y) =>
    {
        for(var i = 0; i < deskProps.length; i++)
        {
            if(deskProps[i].key === selectedId)
            {
                const newProps = deskProps.slice();
                newProps[i].y = parseInt(y);
                SetDeskProps(newProps);
                setResourceYCoord(y);
                break;
            }
        }
    }

    return (
            <Fragment>
                <FaQuestion className='help' size={20} onClick={ViewHelp} />
                <div ref={helpRef} className='help-container'>
                    <span ref={helpToolRef} className='help-tooltip' onClick={CloseHelp}>
                        -Welcome to the office layout creation page.<br></br>
                        -Use the toolbar on the left to add desks and meeting rooms to the floor plan.<br></br>
                        -The properties pane on the right is used to choose the building and room that you are working in.<br></br>
                        -New buildings and rooms can be added using the plus icon by each label respectively. The pencil icon allows for editing the currently selected building/room.<br></br><br></br>
                        -Desks/Meeting Rooms can be moved around by clicking and dragging them.<br></br>
                        -Left click a desk/meeting room once to bring up the transformation gizmo which allows rotation and scaling.<br></br>
                        -Click and drag on the canvas to pan the entire view.<br></br>
                        -The scroll wheel can be used to zoom in and out.<br></br><br></br>
                        Click this help box to close it. It can be reopened by clicking the question mark at the top.

                    </span>
                </div>

                <div className={styles.mapHeadingContainer}>
                    <div className={styles.mapHeading}>Office creator</div>
                </div>

                <div className={styles.propertiesPanel} style={{left: propertiesPanel}}>
                    <p>Name</p>
                    <input type='text' placeholder='Name' value={resourceName} onChange={(e) => ChangeName(e.target.value)}></input>

                    <p>X Coordinate</p>
                    <input type='number' placeholder='X Coordinate' value={resourceXCoord} onChange={(e) => ChangeXCoord(e.target.value)}></input>

                    <p>Y Coordinate</p>
                    <input type='number' placeholder='Y Coordinate' value={resourceYCoord} onChange={(e) => ChangeYCoord(e.target.value)}></input>
                </div>

                <div className='actions-pane'>
                    <div className='save-container' onClick={SaveLayout}>
                        <FaSave className='save-icon' size={30}/>
                    </div>

                    <div className='add-desk-container' onClick={AddDesk}>
                        <img src={desk_white} alt='Add Desk' className='add-desk-img'></img>
                    </div>

                    <div className='add-meetingroom-container' onClick={AddMeetingRoom}>
                        <img src={meetingroom_white} alt='Add Meeting Room' className='add-meetingroom-img'></img>
                    </div>
                </div>                                       

                <div ref={canvasRef} className={styles.canvasContainer}>
                    <Stage width={stage.width} height={stage.height} onMouseDown={CheckDeselect} onTouchStart={CheckDeselect} draggable onWheel={ZoomInOut} ref={stageRef}>
                        <Layer>
                            {deskProps.map((desk, i) => (
                                <Desk
                                    key = {desk.key}
                                    shapeProps = {desk}

                                    isSelected = {desk.key === selectedId}
                                    
                                    onSelect = {() => 
                                    {
                                        SelectShape(desk.key);
                                        setResourceName(deskProps[i].name);
                                        setResourceXCoord(deskProps[i].x);
                                        setResourceYCoord(deskProps[i].y);
                                    }}
                                    
                                    onChange = {(newProps) => 
                                    {
                                        const newDeskProps = deskProps.slice();
                                        newDeskProps[i] = newProps;
                                        setResourceName(newDeskProps[i].name);
                                        setResourceXCoord(newDeskProps[i].x);
                                        setResourceYCoord(newDeskProps[i].y);
                                        SetDeskProps(newDeskProps);
                                    }}

                                    draggable = {true}

                                    transform = {true}
                                />
                            ))}

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

                                        draggable = {true}
                                    />
                                ))
                            )}                             
                        </Layer>
                    </Stage>
                </div>

                <div className={styles.buildingSelectorContainer}>
                    <select className={styles.resourceSelector} name='building' defaultValue={''} onChange={UpdateRooms.bind(this)}>
                        <option value='' disabled id='BuildingDefault'>--Select the building--</option>
                            {buildings.map(building => (
                                <option key={building.id} value={building.id}>{building.name + ' (' + building.location + ')'}</option>
                            ))}
                    </select>

                    <div className={styles.threeDotsContainer}>
                        <BsThreeDotsVertical className={styles.threeDots} onClick={() => ShowBuildingMenu()} />
                    </div>

                    <div ref={buildingMenuRef} className={styles.menu}>
                        <div className={styles.editResource} onMouseDown={EditBuilding.bind(this, currBuilding)}>Edit building</div>
                        <div className={styles.deleteResource}>Remove building</div>
                    </div>
                    
                </div>

                <div className={styles.roomSelectorContainer}>
                    <select className={styles.resourceSelector} name='room' defaultValue={''} onChange={UpdateResources.bind(this)}>
                        <option value='' disabled id='RoomDefault'>--Select the room--</option>
                            {rooms.length > 0 && (
                                rooms.map(room => (
                                    <option key={room.id} value={room.id}>{room.name + ' (' + room.location + ')'}</option>
                                ))
                            )}
                    </select>

                    <div className={styles.threeDotsContainer}>
                        <BsThreeDotsVertical className={styles.threeDots} onClick={() => ShowRoomMenu()} />
                    </div>

                    <div ref={roomMenuRef} className={styles.menu}>
                        <div className={styles.editResource} onMouseDown={EditRoom.bind(this, currRoom)}>Edit room</div>
                        <div className={styles.deleteResource}>Remove room</div>
                    </div>
                </div>
            </Fragment>
    )
}

export default Creator