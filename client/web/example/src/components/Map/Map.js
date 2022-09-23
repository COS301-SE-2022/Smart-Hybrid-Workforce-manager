import { Stage, Layer } from 'react-konva';
import { useRef, useState, useEffect, useCallback, useContext, Fragment } from 'react';
import Desk from './Desk';
import MeetingRoom from './MeetingRoom';
import { UserContext } from '../../App';
import { IoIosArrowBack, IoIosArrowForward } from 'react-icons/io';
import styles from './map.module.css';

const Map = () =>
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

    //Date
    const [date, setDate] = useState('');

    //Desk and meeting room prop arrays
    const [deskProps, SetDeskProps] = useState([]);
    const [meetingRoomProps, SetMeetingRoomProps] = useState([]);
    const [stage, SetStage] = useState({width : 100, height : 100});
    const [selectedId, SelectShape] = useState(null);

    //API fetch variables
    const [buildings, SetBuildings] = useState([]);
    const [rooms, SetRooms] = useState([]);
    const [resources, SetResources] = useState([]);
    const [bookings, setBookings] = useState([]);

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
        /*if(stageRef.current !== null)
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
        }        */
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

        fetch("http://localhost:8080/api/booking/information", 
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
            setBookings(data);
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

    useEffect(() =>
    {
        const date = new Date();

        var day;
        if(date.getDate() < 10)
        {
            day = `0${date.getDate()}`;
        }
        else
        {
            day = `${date.getDate()}`;
        }

        var month;
        if(date.getMonth() + 1 < 10)
        {
            month = `0${date.getMonth() + 1}`;
        }
        else
        {
            month = `${date.getMonth() + 1}`;
        }

        setDate(`${date.getFullYear()}-${month}-${day}`);
    },[]);

    useEffect(() =>
    {
        if(bookings)
        {
            bookings.forEach((booking) =>
            {
                if(booking.start.includes(date))
                {
                    console.log(booking);
                }
            });
        }
    },[date, bookings]);

    return (
        <Fragment>
            <div ref={canvasRef} className={styles.canvasContainer}>
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

                                    ShowUserCard = {(coords) =>
                                    {
                                        const card = document.getElementById('UserCard');
                                        card.style.display = 'block';

                                        const rad = coords.rotation * Math.PI / 180;

                                        const origin = 
                                        {
                                            x: Math.abs(coords.x + stageRef.current.x()) + Math.cos(rad)*coords.width/2 + Math.sin(rad)*coords.height*0.35,
                                            y: Math.abs(coords.y + stageRef.current.y()) + Math.cos(rad)*coords.height*0.35 + Math.sin(rad)*coords.width/2
                                        }

                                        if((coords.rotation >= -10 && coords.rotation <= 10) || (coords.rotation >= 350 && coords.rotation <= 370) || (coords.rotation >= -370 && coords.rotation <= -350)) //Top
                                        {
                                            card.style.left = origin.x + 'px';
                                            card.style.top = origin.y - 0.15*window.innerHeight + 'px';
                                        }
                                        else if((coords.rotation >= 80 && coords.rotation <= 100) || (coords.rotation >= -280 && coords.rotation <= -260)) //Right
                                        {
                                            card.style.left = origin.x + 0.02*window.innerHeight + 'px';
                                            card.style.top = origin.y - 0.10*window.innerHeight + 'px';
                                        }
                                        else if((coords.rotation >= 170 && coords.rotation <= 190) || (coords.rotation >= -190 && coords.rotation <= -170)) //Bottom
                                        {
                                            card.style.left = origin.x + 'px';
                                            card.style.top = origin.y + 0.05*window.innerHeight + 'px';
                                        }
                                        else if((coords.rotation >= 260 && coords.rotation <= 280) || (coords.rotation >= -100 && coords.rotation <= -80)) //Left
                                        {
                                            card.style.left = origin.x - 0.11*window.innerWidth + 'px';
                                            card.style.top = origin.y - 0.10*window.innerHeight + 'px';
                                        }

                                        console.log(coords.rotation);

                                        /*const rad = coords.rotation * Math.PI / 180;

                                        const x = Math.sin(rad) * 0.02*window.innerHeight;
                                        const y = Math.cos(rad) * 0.15*window.innerHeight;

                                        const origin = 
                                        {
                                            x: Math.abs(coords.x + stageRef.current.x()) + Math.cos(rad)*coords.width/2 + Math.sin(rad)*coords.height*0.35,
                                            y: Math.abs(coords.y + stageRef.current.y()) + Math.cos(rad)*coords.height*0.35 + Math.sin(rad)*coords.width/2
                                        }

                                        //if((coords.rotation >= 0 ) || ())

                                        card.style.left = origin.x + x + 'px';
                                        card.style.top = origin.y - y + 'px';

                                        console.log(card.offsetWidth);
                                        console.log(stageRef.current.x() + ' ' + stageRef.current.y());*/
                                    }}

                                    HideUserCard = {() =>
                                    {
                                        document.getElementById('UserCard').style.display = 'none';
                                    }}

                                    draggable = {false}

                                    transform = {false}
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

                                    transform = {false}
                                />
                            ))
                        )}                             
                    </Layer>
                </Stage>
            </div>

            <div className={styles.buildingSelectorContainer}>
                <select ref={buildingRef} className={styles.resourceSelector} name='building' onChange={UpdateRooms.bind(this)}>
                    <option value='' disabled selected id='BuildingDefault'>--Select the building--</option>
                        {buildings.length > 0 && (
                            buildings.map(building => (
                                <option key={building.id} value={building.id}>{building.name + ' (' + building.location + ')'}</option>
                            ))
                        )}
                </select>
            </div>

            <div className={styles.roomSelectorContainer}>
                <select ref={roomRef} className={styles.resourceSelector} name='room' onChange={UpdateResources.bind(this)}>
                    <option value='' disabled selected id='RoomDefault'>--Select the room--</option>
                        {rooms.length > 0 && (
                            rooms.map(room => (
                                <option key={room.id} value={room.id}>{room.name + ' (' + room.location + ')'}</option>
                            ))
                        )}
                </select>
            </div>

            <div className={styles.dateContainer}>
                <input className={styles.resourceSelector} type='date' id='BookingDate' onChange={(e) => setDate(e.target.value)}></input>
            </div>

            <div id='UserCard' className={styles.userCard}>E</div>

        </Fragment>
    )
}

export default Map