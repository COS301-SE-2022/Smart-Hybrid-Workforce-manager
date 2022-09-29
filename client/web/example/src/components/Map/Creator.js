import { Stage, Layer } from 'react-konva';
import { useRef, useState, useEffect, useCallback, useContext, Fragment } from 'react';
import Desk from './Desk';
import MeetingRoom from './MeetingRoom';
import Wall from './Wall';
import { FaSave } from 'react-icons/fa';
import { MdAdd, MdClose } from 'react-icons/md';
import { BsThreeDotsVertical } from 'react-icons/bs';
import { UserContext } from '../../App';
import styles from './map.module.css';
import { AddBuildingForm } from '../Resources/AddBuilding';
import { EditBuildingForm } from '../Resources/EditBuilding';
import { AddRoomForm } from '../Resources/AddRoom';
import { EditRoomForm } from '../Resources/EditRoom';

const Creator = () =>
{
    //Canvas references
    const canvasRef = useRef(null);
    const stageRef = useRef(null);
    const scaleFactor = 1.3;
    const deskPropsRef = useRef([]);
    const meetingRoomPropsRef = useRef([]);
    const wallPropsRef = useRef([]);
    const deskCount = useRef(0);
    const meetingRoomCount = useRef(0);
    const wallCount = useRef(0);
    const deletedResources = useRef([]);

    //Building and rooms
    const backgroundDimmerRef = useRef(null);

    const buildingMenuRef = useRef(null);
    const addBuildingRef = useRef(null);
    const editBuildingRef = useRef(null);
    const [addBuilding, setAddBuilding] = useState(false);
    const [buildingEdited, setBuildingEdited] = useState(true);

    const roomMenuRef = useRef(null);
    const addRoomRef = useRef(null);
    const editRoomRef = useRef(null);
    const [addRoom, setAddRoom] = useState(false);
    const [roomEdited, setRoomEdited] = useState(false);

    //Panel states
    const [propertiesPanel, setPropertiesPanel] = useState(0.985*window.innerWidth);
    const [resourceName, setResourceName] = useState('');
    const [resourceXCoord, setResourceXCoord] = useState('');
    const [resourceYCoord, setResourceYCoord] = useState('');
    const [resourceRotation, setResourceRotation] = useState('');
    const capacityLabelRef = useRef(null);
    const capacityInputRef = useRef(null);
    const [resourceCapacity, setResourceCapacity] = useState('');
    const lengthLabelRef = useRef(null);
    const lengthInputRef = useRef(null);
    const [wallLength, setWallLength] = useState('');

    //Desk and meeting room prop arrays
    const [deskProps, SetDeskProps] = useState([]);
    const [meetingRoomProps, SetMeetingRoomProps] = useState([]);
    const [wallProps, SetWallProps] = useState([]);
    const [stage, SetStage] = useState({width : 100, height : 100});
    const [selectedId, SelectShape] = useState(null);

    //API fetch variables
    const [buildings, SetBuildings] = useState([]);
    const [currBuilding, SetCurrBuilding] = useState("");
    const [rooms, SetRooms] = useState([]);
    const [currRoom, SetCurrRoom] = useState("");
    const [resources, SetResources] = useState([]);

    const {userData} = useContext(UserContext);

    //POST requests
    const UpdateRooms = (id) =>
    {
        fetch("http://deskflow.co.za:8080/api/resource/room/information", 
        {
            method: "POST",
            mode: 'cors',
            body: JSON.stringify({
                building_id: id
            }),
        headers:{
            'Content-Type': 'application/json',
            'Authorization': `bearer ${userData.token}`
        }
        }).then((res) => res.json()).then(data => 
        {
            SetRooms(data);
            document.getElementById("RoomDefault").selected = true;
            SetCurrRoom("");
            SetCurrBuilding(id);
            SetResources([]);
        });
    }

    const UpdateResources = (id) =>
    {
        fetch("http://deskflow.co.za:8080/api/resource/information", 
        {
            method: "POST",
            mode: 'cors',
            body: JSON.stringify({
                room_id: id
            }),
        headers:{
            'Content-Type': 'application/json',
            'Authorization': `bearer ${userData.token}`
        }
        }).then((res) => res.json()).then(data => 
        {
            SetResources(data);
            SetCurrRoom(id);
        });
    }

    //Canvas functions
    //Check if canvas is clicked and deselect the selected resource
    const CheckDeselect = (e) =>
    {
        e.target.getStage().container().style.cursor = 'grabbing';
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
    const OpenAddBuilding = () =>
    {
        if(backgroundDimmerRef.current && addBuildingRef.current && buildingMenuRef.current)
        {
            backgroundDimmerRef.current.style.display = 'block';
            addBuildingRef.current.style.display = 'block'
            buildingMenuRef.current.style.display = 'none';
            setAddBuilding(!addBuilding);
        }
    }

    const CloseAddBuilding = () =>
    {
        if(backgroundDimmerRef.current && addBuildingRef)
        {
            backgroundDimmerRef.current.style.display = 'none';
            addBuildingRef.current.style.display = 'none'
        }
    }
    
    //Edit selected building
    const OpenEditBuilding = () =>
    {
        if(backgroundDimmerRef.current && editBuildingRef.current && buildingMenuRef.current && currBuilding !== '')
        {
            backgroundDimmerRef.current.style.display = 'block';
            editBuildingRef.current.style.display = 'block'
            buildingMenuRef.current.style.display = 'none';
        }
    }

    const CloseEditBuilding = () =>
    {
        if(backgroundDimmerRef.current && editBuildingRef)
        {
            backgroundDimmerRef.current.style.display = 'none';
            editBuildingRef.current.style.display = 'none'
        }
    }

    const DeleteBuilding = () =>
    {
        if(currBuilding !== '' && buildingMenuRef.current)
        {
            buildingMenuRef.current.style.display = 'none';

            fetch("http://deskflow.co.za:8080/api/resource/building/remove", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    id: currBuilding
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}`
                }
            }).then((res) =>
            {
                setBuildingEdited(true);
            });   
        }
    }

    const ShowRoomMenu = () =>
    {
        if(roomMenuRef.current)
        {
            roomMenuRef.current.style.display = 'block';
        }
    }

    //Add room
    const OpenAddRoom = () =>
    {
        if(backgroundDimmerRef.current && addRoomRef.current && roomMenuRef.current && currBuilding !== '')
        {
            backgroundDimmerRef.current.style.display = 'block';
            addRoomRef.current.style.display = 'block'
            roomMenuRef.current.style.display = 'none';
            setAddRoom(!addRoom);
        }
    }

    const CloseAddRoom = () =>
    {
        if(backgroundDimmerRef.current && addRoomRef)
        {
            backgroundDimmerRef.current.style.display = 'none';
            addRoomRef.current.style.display = 'none'
        }
    }
    
    //Edit selected room
    const OpenEditRoom = () =>
    {
        if(backgroundDimmerRef.current && editRoomRef.current && roomMenuRef.current && currRoom !== '')
        {
            backgroundDimmerRef.current.style.display = 'block';
            editRoomRef.current.style.display = 'block'
            roomMenuRef.current.style.display = 'none';
        }
    }

    const CloseEditRoom = () =>
    {
        if(backgroundDimmerRef.current && editRoomRef)
        {
            backgroundDimmerRef.current.style.display = 'none';
            editRoomRef.current.style.display = 'none'
        }
    }

    const DeleteRoom = () =>
    {
        if(currRoom !== '' && currBuilding !== '' && roomMenuRef.current)
        {
            console.log(currRoom);
            roomMenuRef.current.style.display = 'none';

            fetch("http://deskflow.co.za:8080/api/resource/room/remove", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    id: currRoom,
                    building_id: currBuilding,
                    name: null,
                    xcoord: null,
                    ycoord: null,
                    zcoord: null,
                    dimension: null
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}`
                }
            }).then((res) =>
            {
                setRoomEdited(true);
            });   
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
        if(currBuilding === "" || currRoom === "")
        {
            window.alert("Please select a building and room");
            return;
        }

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
    const LoadMeetingRoom = useCallback((id, name, x, y, width, height, rotation, capacity) =>
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
                    capacity: capacity,
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
        if(currBuilding === "" || currRoom === "")
        {
            window.alert("Please select a building and room");
            return;
        }

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
                    width : 0.3,
                    height : 0.3,
                    rotation : 0,
                    capacity: 2,
                    edited : true
                }
            ]);
        }
    };

    //Load walls from the database
    const LoadWall = useCallback((id, name, x, y, width, height, rotation) =>
    {
        //Uses a reference array to prevent state dependency and infinite loop
        if(stageRef.current !== null)
        {
            wallPropsRef.current =
            [
                ...wallPropsRef.current,
                {
                    key : "wall" + id,
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
            SetWallProps(wallPropsRef.current);
        }
    },[]);

    //Add a new wall to the state
    const AddWall = () =>
    {
        if(currBuilding === "" || currRoom === "")
        {
            window.alert("Please select a building and room");
            return;
        }

        if(stageRef.current !== null)
        {
            SetWallProps(
            [
                ...wallProps,
                {
                    key : "wall" + wallCount.current,
                    id : null,
                    name : "Wall " + wallCount.current,
                    x : (-stageRef.current.x() + stageRef.current.width() / 2.0) / stageRef.current.scaleX(),
                    y : (-stageRef.current.y() + stageRef.current.height() / 2.0) / stageRef.current.scaleY(),
                    width : 200,
                    height : 40,
                    rotation : 0,
                    edited : false
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
                for(let i = 0; i < deskProps.length; i++)
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
            else if(selectedId.includes('meetingroom'))
            {
                for(let i = 0; i < meetingRoomProps.length; i++)
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
            else if(selectedId.includes("wall"))
            {
                for(let i = 0; i < wallProps.length; i++)
                {
                    if(wallProps[i].key === selectedId)
                    {
                        if(wallProps[i].id !== null)
                        {
                            deletedResources.current.push(wallProps[i]);
                        }

                        var newWall = [...wallProps];
                        newWall.splice(i, 1);
                        SetWallProps(newWall);
                        SelectShape(null);
                        break;
                    }
                }
            }
        }
    }, [deskProps, meetingRoomProps, wallProps, selectedId])

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

        for(let i = 0; i < deskProps.length; i++)
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

        for(let i = 0; i < meetingRoomProps.length; i++)
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

        for(let i = 0; i < wallProps.length; i++)
        {
            var currWall = wallProps[i];
            
            if(currWall.edited)
            {
                resources.push(
                {
                    id : currWall.id,
                    room_id: currRoom,
                    name: currWall.name,
                    xcoord: currWall.x,
                    ycoord: currWall.y,
                    width: currWall.width,
                    height: currWall.height,
                    rotation: currWall.rotation,
                    resource_type: 'WALL',
                    decorations: '{}',
                })
            }
        }

        try
        {
            let res = await fetch("http://deskflow.co.za:8080/api/resource/batch-create", 
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
                if(deletedResources.current.length > 0)
                {
                    var deletedArray = [];
                    for(let i = 0; i < deletedResources.current.length; i++)
                    {
                        var deleted = deletedResources.current[i];

                        deletedArray.push(
                        {
                            id : deleted.id
                        });
                    }

                    let res = await fetch("http://deskflow.co.za:8080/api/resource/batch-remove", 
                    {
                        method: "POST",
                        mode: 'cors',
                        body: JSON.stringify(deletedArray),
                        headers:{
                            'Content-Type': 'application/json',
                            'Authorization': `bearer ${userData.token}`
                        }
                    });

                    if(res.status === 200)
                    {
                        deletedResources.current = [];
                        alert("Saved!");
                        UpdateResources(currRoom);
                    }
                }
                else
                {
                    alert("Saved!");
                    UpdateResources(currRoom);
                }
            }
        }
        catch(err)
        {
            console.log(err);
        }
    }

    //Effect on the loading of the web page
    useEffect(() =>
    {
        SetStage({width : canvasRef.current.offsetWidth, height : canvasRef.current.offsetHeight});

        if(buildingEdited)
        {
            fetch("http://deskflow.co.za:8080/api/resource/building/information", 
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

            setBuildingEdited(false);
        }
    },[userData.token, buildingEdited]);

    useEffect(() =>
    {
        if(roomEdited && currBuilding !== '')
        {
            fetch("http://deskflow.co.za:8080/api/resource/room/information", 
            {
                method: "POST",
                mode: 'cors',
                body: JSON.stringify({
                    building_id: currBuilding
                }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}`
            }
            }).then((res) => res.json()).then(data => 
            {
                SetRooms(data);
            });

            setRoomEdited(false);
        }
    },[userData.token, currBuilding, roomEdited]);

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
        deletedResources.current = [];
        deskPropsRef.current = [];
        deskCount.current = 0;
        meetingRoomPropsRef.current = [];
        meetingRoomCount.current = 0;
        wallPropsRef.current = [];
        wallCount.current = 0;

        SetDeskProps(deskPropsRef.current);
        SetMeetingRoomProps(meetingRoomPropsRef.current);
        SetWallProps(wallPropsRef.current);

        //Loop through resources and load desks and meeting rooms respectively
        for(var i = 0; i < resources.length; i++)
        {
            if(resources[i].resource_type === "DESK")
            {
                LoadDesk(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation);
            }
            else if(resources[i].resource_type === "MEETINGROOM")
            {
                LoadMeetingRoom(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation, JSON.parse(resources[i].decorations).capacity);
            }
            else if(resources[i].resource_type === "WALL")
            {
                LoadWall(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation);
            }
        }

    }, [resources, LoadDesk, LoadMeetingRoom, LoadWall]);

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

    //Update the wall counter when a new wall is added or removed
    useEffect(() =>
    {
        wallCount.current = wallProps.length;
    }, [wallProps.length]);

    //Check if properties are open or closed
    useEffect(() =>
    {
        if(selectedId)
        {
            setPropertiesPanel(0.65*window.innerWidth);
            if(lengthLabelRef.current && lengthInputRef.current)
            {
                if(selectedId.includes('wall'))
                {
                    lengthLabelRef.current.style.display = 'block';
                    lengthInputRef.current.style.display = 'inline-block';
                }
                else
                {
                    lengthLabelRef.current.style.display = 'none';
                    lengthInputRef.current.style.display = 'none';
                }

                if(selectedId.includes('meetingroom'))
                {
                    capacityLabelRef.current.style.display = 'block';
                    capacityInputRef.current.style.display = 'inline-block';
                }
                else
                {
                    capacityLabelRef.current.style.display = 'none';
                    capacityInputRef.current.style.display = 'none';
                }
            }
        }
        else
        {
            setPropertiesPanel(0.85*window.innerWidth);
        }
    },[selectedId, wallLength])

    const ChangeName = (name) =>
    {
        if(selectedId.includes('desk'))
        {
            for(let i = 0; i < deskProps.length; i++)
            {
                if(deskProps[i].key === selectedId)
                {
                    const newProps = deskProps.slice();
                    newProps[i].name = name;
                    newProps[i].edited = true;
                    SetDeskProps(newProps);
                    setResourceName(name);
                    break;
                }
            }
        }
        else if(selectedId.includes('meetingroom'))
        {
            for(let i = 0; i < meetingRoomProps.length; i++)
            {
                if(meetingRoomProps[i].key === selectedId)
                {
                    const newProps = meetingRoomProps.slice();
                    newProps[i].name = name;
                    newProps[i].edited = true;
                    SetMeetingRoomProps(newProps);
                    setResourceName(name);
                    break;
                }
            }
        }
        else if(selectedId.includes('wall'))
        {
            for(let i = 0; i < wallProps.length; i++)
            {
                if(wallProps[i].key === selectedId)
                {
                    const newProps = wallProps.slice();
                    newProps[i].name = name;
                    newProps[i].edited = true;
                    SetWallProps(newProps);
                    setResourceName(name);
                    break;
                }
            }
        }
    }

    const ChangeXCoord = (x) =>
    {
        if(selectedId.includes('desk'))
        {
            for(let i = 0; i < deskProps.length; i++)
            {
                if(deskProps[i].key === selectedId)
                {
                    const newProps = deskProps.slice();
                    newProps[i].x = parseInt(x);
                    newProps[i].edited = true;
                    SetDeskProps(newProps);
                    setResourceXCoord(x);
                    break;
                }
            }
        }
        else if(selectedId.includes('meetingroom'))
        {
            for(let i = 0; i < meetingRoomProps.length; i++)
            {
                if(meetingRoomProps[i].key === selectedId)
                {
                    const newProps = meetingRoomProps.slice();
                    newProps[i].x = parseInt(x);
                    newProps[i].edited = true;
                    SetMeetingRoomProps(newProps);
                    setResourceXCoord(x);
                    break;
                }
            }
        }
        else if(selectedId.includes('wall'))
        {
            for(let i = 0; i < wallProps.length; i++)
            {
                if(wallProps[i].key === selectedId)
                {
                    const newProps = wallProps.slice();
                    newProps[i].x = parseInt(x);
                    newProps[i].edited = true;
                    SetWallProps(newProps);
                    setResourceXCoord(x);
                    break;
                }
            }
        }
    }

    const ChangeYCoord = (y) =>
    {
        if(selectedId.includes('desk'))
        {
            for(let i = 0; i < deskProps.length; i++)
            {
                if(deskProps[i].key === selectedId)
                {
                    const newProps = deskProps.slice();
                    newProps[i].y = parseInt(y);
                    newProps[i].edited = true;
                    SetDeskProps(newProps);
                    setResourceYCoord(y);
                    break;
                }
            }
        }
        else if(selectedId.includes('meetingroom'))
        {
            for(let i = 0; i < meetingRoomProps.length; i++)
            {
                if(meetingRoomProps[i].key === selectedId)
                {
                    const newProps = meetingRoomProps.slice();
                    newProps[i].y = parseInt(y);
                    newProps[i].edited = true;
                    SetMeetingRoomProps(newProps);
                    setResourceYCoord(y);
                    break;
                }
            }
        }
        else if(selectedId.includes('wall'))
        {
            for(let i = 0; i < wallProps.length; i++)
            {
                if(wallProps[i].key === selectedId)
                {
                    const newProps = wallProps.slice();
                    newProps[i].y = parseInt(y);
                    newProps[i].edited = true;
                    SetWallProps(newProps);
                    setResourceYCoord(y);
                    break;
                }
            }
        }
    }

    const ChangeRotation = (rotation) =>
    {
        if(selectedId.includes('desk'))
        {
            for(let i = 0; i < deskProps.length; i++)
            {
                if(deskProps[i].key === selectedId)
                {
                    const newProps = deskProps.slice();
                    newProps[i].rotation = parseInt(rotation);
                    newProps[i].edited = true;
                    SetDeskProps(newProps);
                    setResourceRotation(rotation);
                    break;
                }
            }
        }
        else if(selectedId.includes('meetingroom'))
        {
            for(let i = 0; i < meetingRoomProps.length; i++)
            {
                if(meetingRoomProps[i].key === selectedId)
                {
                    const newProps = meetingRoomProps.slice();
                    newProps[i].rotation = parseInt(rotation);
                    newProps[i].edited = true;
                    SetMeetingRoomProps(newProps);
                    setResourceRotation(rotation);
                    break;
                }
            }
        }
        else if(selectedId.includes('wall'))
        {
            for(let i = 0; i < wallProps.length; i++)
            {
                if(wallProps[i].key === selectedId)
                {
                    const newProps = wallProps.slice();
                    newProps[i].rotation = parseInt(rotation);
                    newProps[i].edited = true;
                    SetWallProps(newProps);
                    setResourceRotation(rotation);
                    break;
                }
            }
        }
    }

    const ChangeCapacity = (capacity) =>
    {
        if(selectedId.includes('meetingroom'))
        {
            for(let i = 0; i < meetingRoomProps.length; i++)
            {
                if(meetingRoomProps[i].key === selectedId)
                {
                    const newProps = meetingRoomProps.slice();
                    newProps[i].capacity = parseInt(capacity);
                    newProps[i].edited = true;
                    SetMeetingRoomProps(newProps);
                    setResourceCapacity(capacity);
                    break;
                }
            }
        }
    }

    const ChangeLength = (width) =>
    {
        if(selectedId.includes('wall'))
        {
            for(let i = 0; i < wallProps.length; i++)
            {
                if(wallProps[i].key === selectedId)
                {
                    const newProps = wallProps.slice();
                    newProps[i].width = parseInt(width);
                    newProps[i].edited = true;
                    SetWallProps(newProps);
                    setWallLength(width);
                    break;
                }
            }
        }
    }

    return (
            <Fragment>
                <div className={styles.mapHeadingContainer}>
                    <div className={styles.mapHeading}>Office creator</div>
                </div>

                <div ref={backgroundDimmerRef} className={styles.backgroundDimmer}></div>

                <div ref={addBuildingRef} className={styles.formContainer}>
                    <div className={styles.formClose} onClick={() => CloseAddBuilding()}><MdClose /></div>
                    <AddBuildingForm makeDefault={addBuilding} edited={setBuildingEdited} />
                </div>

                <div ref={editBuildingRef} className={styles.formContainer}>
                    <div className={styles.formClose} onClick={() => CloseEditBuilding()}><MdClose /></div>
                    <EditBuildingForm id={currBuilding} edited={setBuildingEdited} />
                </div>

                <div ref={addRoomRef} className={styles.formContainer}>
                    <div className={styles.formClose} onClick={() => CloseAddRoom()}><MdClose /></div>
                    <AddRoomForm makeDefault={addRoom} edited={setRoomEdited} buildingID={currBuilding} />
                </div>

                <div ref={editRoomRef} className={styles.formContainer}>
                    <div className={styles.formClose} onClick={() => CloseEditRoom()}><MdClose /></div>
                    <EditRoomForm id={currRoom} edited={setRoomEdited} />
                </div>

                <div className={styles.propertiesPanel} style={{left: propertiesPanel}}>
                    <div className={styles.propertiesHeading}>Properties</div>

                    <div className={styles.formLabel}>Name</div>
                    <input className={styles.formInput} type='text' placeholder='Name' value={resourceName} onChange={(e) => ChangeName(e.target.value)}></input>

                    <div className={styles.formLabel}>X Coordinate</div>
                    <input className={styles.formInput} type='number' placeholder='X Coordinate' value={Math.trunc(resourceXCoord)} onChange={(e) => ChangeXCoord(e.target.value)}></input>

                    <div className={styles.formLabel}>Y Coordinate</div>
                    <input className={styles.formInput} type='number' placeholder='Y Coordinate' value={Math.trunc(resourceYCoord)} onChange={(e) => ChangeYCoord(e.target.value)}></input>

                    <div ref={lengthLabelRef} className={styles.formLabel}>Length</div>
                    <input ref={lengthInputRef} className={styles.formInput} type='number' placeholder='Length' disabled={selectedId ? !selectedId.includes('wall') : true} value={Math.trunc(wallLength)} onChange={(e) => ChangeLength(e.target.value)}></input>

                    <div className={styles.formLabel}>Rotation</div>
                    <input className={styles.formInput} type='number' placeholder='Rotation' value={Math.trunc(resourceRotation)} onChange={(e) => ChangeRotation(e.target.value)}></input>

                    <div ref={capacityLabelRef} className={styles.formLabel}>Capacity</div>
                    <input ref={capacityInputRef} className={styles.formInput} type='number' placeholder='Capacity' disabled={selectedId ? !selectedId.includes('meetingroom') : true} value={Math.trunc(resourceCapacity)} onChange={(e) => ChangeCapacity(e.target.value)}></input>
                </div>

                <div className={styles.actions}>
                    <div className={styles.addResource} onClick={SaveLayout}>
                        <FaSave />{' Save'}
                    </div>

                    <div className={styles.editResource}  onClick={AddDesk}>
                        <MdAdd />{' Add desk'}
                    </div>

                    <div className={styles.editResource}  onClick={AddMeetingRoom}>
                        <MdAdd />{' Add meeting room'}
                    </div>

                    <div className={styles.deleteResource}  onClick={AddWall}>
                        <MdAdd />{' Add wall'}
                    </div>
                </div>                                       

                <div ref={canvasRef} className={styles.canvasContainer}>
                    <Stage width={stage.width} height={stage.height} onMouseDown={CheckDeselect} onMouseUp={(e) => e.target.getStage().container().style.cursor = 'grab'} onTouchStart={CheckDeselect} draggable onWheel={ZoomInOut} ref={stageRef}>
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
                                        setResourceRotation(deskProps[i].rotation);
                                        setResourceCapacity('');
                                    }}
                                    
                                    onChange = {(newProps) => 
                                    {
                                        const newDeskProps = deskProps.slice();
                                        newDeskProps[i] = newProps;
                                        setResourceName(newDeskProps[i].name);
                                        setResourceXCoord(newDeskProps[i].x);
                                        setResourceYCoord(newDeskProps[i].y);
                                        setResourceRotation(newDeskProps[i].rotation);
                                        setResourceCapacity('');
                                        SetDeskProps(newDeskProps);
                                    }}

                                    draggable = {true}

                                    transform = {true}
                                />
                            ))}

                            {meetingRoomProps.map((meetingRoom, i) => (
                                <MeetingRoom
                                    key = {meetingRoom.key}
                                    shapeProps = {meetingRoom}

                                    isSelected = {meetingRoom.key === selectedId}
                                    
                                    onSelect = {() => 
                                    {
                                        SelectShape(meetingRoom.key);
                                        setResourceName(meetingRoomProps[i].name);
                                        setResourceXCoord(meetingRoomProps[i].x);
                                        setResourceYCoord(meetingRoomProps[i].y);
                                        setResourceRotation(meetingRoomProps[i].rotation);
                                        setResourceCapacity(meetingRoomProps[i].capacity);
                                    }}
                                    
                                    onChange = {(newProps) => 
                                    {
                                        const newMeetingRoomProps = meetingRoomProps.slice();
                                        newMeetingRoomProps[i] = newProps;
                                        setResourceName(newMeetingRoomProps[i].name);
                                        setResourceXCoord(newMeetingRoomProps[i].x);
                                        setResourceYCoord(newMeetingRoomProps[i].y);
                                        setResourceRotation(newMeetingRoomProps[i].rotation);
                                        setResourceCapacity(meetingRoomProps[i].capacity);
                                        SetMeetingRoomProps(newMeetingRoomProps);
                                    }}

                                    draggable = {true}

                                    transform = {true}
                                />
                            ))}    

                            {wallProps.map((wall, i) => (
                                <Wall
                                    key = {wall.key}
                                    shapeProps = {wall}

                                    isSelected = {wall.key === selectedId}
                                    
                                    onSelect = {() => 
                                    {
                                        SelectShape(wall.key);
                                        setResourceName(wallProps[i].name);
                                        setResourceXCoord(wallProps[i].x);
                                        setResourceYCoord(wallProps[i].y);
                                        setResourceRotation(wallProps[i].rotation);
                                        setResourceCapacity('');
                                        setWallLength(wallProps[i].width);
                                    }}
                                    
                                    onChange = {(newProps) => 
                                    {
                                        const newWallProps = wallProps.slice();
                                        newWallProps[i] = newProps;
                                        setResourceName(newWallProps[i].name);
                                        setResourceXCoord(newWallProps[i].x);
                                        setResourceYCoord(newWallProps[i].y);
                                        setResourceRotation(newWallProps[i].rotation);
                                        setResourceCapacity('');
                                        setWallLength(newWallProps[i].width);
                                        SetWallProps(newWallProps);
                                    }}

                                    draggable = {true}

                                    transform = {true}
                                />
                            ))}                        
                        </Layer>
                    </Stage>
                </div>

                <div className={styles.buildingSelectorContainer}>
                    <select className={styles.resourceSelector} name='building' defaultValue={''} onChange={(e) => UpdateRooms(e.target.value)}>
                        <option value='' id='BuildingDefault'>--Select the building--</option>
                            {buildings.map(building => (
                                <option key={building.id} value={building.id}>{building.name + ' (' + building.location + ')'}</option>
                            ))}
                    </select>

                    <div className={styles.threeDotsContainer}>
                        <BsThreeDotsVertical className={styles.threeDots} onClick={() => ShowBuildingMenu()} />
                    </div>

                    <div ref={buildingMenuRef} className={styles.menu}>
                        <div className={styles.addResource} onClick={() => OpenAddBuilding()}>Add building</div>
                        <div className={styles.editResource} onClick={() => OpenEditBuilding()}>Edit building</div>
                        <div className={styles.deleteResource} onClick={() => DeleteBuilding()}>Remove building</div>
                    </div>
                    
                </div>

                <div className={styles.roomSelectorContainer}>
                    <select className={styles.resourceSelector} name='room' defaultValue={''} onChange={(e) => UpdateResources(e.target.value)}>
                        <option value='' id='RoomDefault'>--Select the room--</option>
                            {rooms.map(room =>
                            (
                                <option key={room.id} value={room.id}>{room.name + ' (Floor ' + room.zcoord + ')'}</option>
                            ))}
                    </select>

                    <div className={styles.threeDotsContainer}>
                        <BsThreeDotsVertical className={styles.threeDots} onClick={() => ShowRoomMenu()} />
                    </div>

                    <div ref={roomMenuRef} className={styles.menu}>
                    <div className={styles.addResource} onClick={() => OpenAddRoom()}>Add room</div>
                        <div className={styles.editResource} onClick={() => OpenEditRoom()}>Edit room</div>
                        <div className={styles.deleteResource} onClick={() => DeleteRoom()}>Remove room</div>
                    </div>
                </div>
            </Fragment>
    )
}

export default Creator