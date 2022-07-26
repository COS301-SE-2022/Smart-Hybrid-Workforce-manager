import { Stage, Layer } from 'react-konva'
import { useRef, useState, useEffect, useCallback, useReducer } from 'react'
import Desk from '../components/Map/Desk'
import MeetingRoom from '../components/Map/MeetingRoom'

const Layout = () =>
{
    //Reducer initialState
    const initialState = {
        counter : 0
    }

    //Reducer function
    const reducer = (state, action) =>
    {
        var newState;
        if(action.type === "increase")
        {
            newState = 
            {
                counter : state.counter + 1
            };
        }
        else
        {
            newState = 
            {
                counter : state.counter - 1
            };
        }

        return newState;
    }

    //Canvas references
    const canvasRef = useRef(null);
    const stageRef = useRef(null);
    const scaleFactor = 1.3;
    const deskProps = useRef([]);
    const count = useRef(0);

    //Desk and meeting room prop arrays
    //const [deskProps, setDeskProps] = useState([]);
    const [meetingRoomProps, setMeetingRoomProps] = useState([]);
    const [stage, setStage] = useState({width : 100, height : 100});
    const [state, dispatch] = useReducer(reducer, initialState);
    const [selectedId, selectShape] = useState(null);

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
    const checkDeselect = (e) =>
    {
        const clickedEmpty = e.target === e.target.getStage();
        if(clickedEmpty)
        {
            selectShape(null);
        }
    }

    //Add a desk to the array with default props
    const LoadDesk = useCallback((id, name, x, y, width, height, rotation) =>
    {
        if(stageRef.current !== null)
        {
            deskProps.current =
            [
                ...deskProps.current,
                {
                    key : "desk" + count.current,
                    id : id,
                    name : name,
                    x : x,
                    y : y,
                    width : width,
                    height : height,
                    rotation : rotation
                }
            ];

            count.current = count.current + 1;
        }
    },[]);

    const AddDesk = () =>
    {
        if(stageRef.current !== null)
        {
            console.log(deskProps.current);
            deskProps.current =
            [
                ...deskProps.current,
                {
                    key : "desk" + count.current,
                    id : null,
                    name : "Desk " + count.current,
                    x : (-stageRef.current.x() + stageRef.current.width() / 2.0) / stageRef.current.scaleX(),
                    y : (-stageRef.current.y() + stageRef.current.height() / 2.0) / stageRef.current.scaleY(),
                    width : 60,
                    height : 55,
                    rotation : 0
                }
            ];

            count.current = count.current + 1;
        }
    };

    //Add a meeting room to the array with default props
    const AddMeetingRoom = () =>
    {
        setMeetingRoomProps(
        [
            ...meetingRoomProps,
            {
                key : "meetingroom" + state.counter,
                x : (-stageRef.current.x() + stageRef.current.width() / 2.0) / stageRef.current.scaleX(),
                y : (-stageRef.current.y() + stageRef.current.height() / 2.0) / stageRef.current.scaleY(),
                width : 200,
                height : 200,
                rotation : 0
            }
        ]);

        dispatch({type : "increase"});
    }

    //Check if resource is selected and delete key is pressed
    const deletePressed = useKeyPress("Delete")

    function useKeyPress(targetKey)
    {
        // State for keeping track of whether key is pressed
        const [keyPressed, setKeyPressed] = useState(false);

        // If pressed key is our target key then set to true
        const downHandler = useCallback(({ key }) =>
        {
            if (key === targetKey)
            {
                setKeyPressed(true);
            }
        },[targetKey]);

        // If released key is our target key then set to false
        const upHandler = useCallback(({ key }) =>
        {
            if (key === targetKey)
            {
                setKeyPressed(false);
            }
        },[targetKey]);
        
        // Add event listeners
        useEffect(() =>
        {
            window.addEventListener("keydown", downHandler);
            window.addEventListener("keyup", upHandler);

            // Remove event listeners on cleanup
            return () => {
                window.removeEventListener("keydown", downHandler);
                window.removeEventListener("keyup", upHandler);
            };
        }, [downHandler, upHandler]);

        return keyPressed;
    }

    const handleDelete = useCallback(() =>
    {
        if(selectedId !== null)
        {
            if(selectedId.includes("desk"))
            {
                for(var i = 0; i < deskProps.current.length; i++)
                {
                    if(deskProps.current[i].key === selectedId)
                    {
                        var newDesk = [...deskProps.current];
                        newDesk.splice(i, 1);
                        //setDeskProps(newDesk);
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
                        setMeetingRoomProps(newMeetingRoom);
                    }
                }
            }
        }
    }, [meetingRoomProps, selectedId])

    //Adjusts the canvas size for difference screen sizes
    const handleResize = () =>
    {
        setStage({width : canvasRef.current.offsetWidth, height : canvasRef.current.offsetHeight});
    }

    window.addEventListener('resize', handleResize);

    const canvasDrag = () =>
    {

    }

    //Ensures that the zooming in/out is oriented with the center of viewable canvas
    const zoomInOut = (event) =>
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

    useEffect(() =>
    {
        setStage({width : canvasRef.current.offsetWidth, height : canvasRef.current.offsetHeight});
        FetchBuildings();
        
        if(deletePressed)
        {
            handleDelete();
        }

    }, [deletePressed, handleDelete])

    useEffect(() =>
    {
        deskProps.current = [];
        count.current = 0;
        for(var i = 0; i < resources.length; i++)
        {
            if(resources[i].resource_type === "DESK")
            {
                console.log("DESK " + resources[i].name);
                LoadDesk(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation);
            }
        }

    }, [resources, LoadDesk])

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
                    <Stage width={stage.width} height={stage.height} onMouseDown={checkDeselect} onTouchStart={checkDeselect} draggable onDragEnd={canvasDrag} onWheel={zoomInOut} ref={stageRef}>
                        <Layer>
                            {deskProps.current.length > 0 && (
                                deskProps.current.map((desk, i) => (
                                    <Desk
                                        key = {desk.key}
                                        shapeProps = {desk}

                                        isSelected = {desk.key === selectedId}
                                        
                                        onSelect = {() => 
                                        {
                                            selectShape(desk.key);
                                        }}
                                        
                                        onChange = {(newProps) => 
                                        {
                                            const newDeskProps = deskProps.current.slice();
                                            newDeskProps[i] = newProps;
                                            //setDeskProps(newDeskProps)
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
                                            selectShape(meetingRoom.key);
                                        }}
                                        
                                        onChange = {(newProps) => 
                                        {
                                            const newMeetingRoomProps = meetingRoomProps.slice();
                                            newMeetingRoomProps[i] = newProps;
                                            setMeetingRoomProps(newMeetingRoomProps)
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