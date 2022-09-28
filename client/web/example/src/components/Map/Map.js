import { Stage, Layer } from 'react-konva';
import { useRef, useState, useEffect, useCallback, useContext, Fragment } from 'react';
import Desk from './Desk';
import MeetingRoom from './MeetingRoom';
import Wall from './Wall';
import { UserContext } from '../../App';
import styles from './map.module.css';

const Map = () =>
{
    //Canvas references
    const canvasRef = useRef(null);
    const stageRef = useRef(null);
    const deskPropsRef = useRef([]);
    const meetingRoomPropsRef = useRef([]);
    const wallPropsRef = useRef([]);
    const dateRef = useRef(null);
    const scaleFactor = 1.3;

    //Selector refs
    const buildingRef = useRef(null);
    const roomRef = useRef(null);

    const preferenceRef = useRef(null);

    //Date
    const [date, setDate] = useState('');

    //Side panel
    const [sidePanel, setSidePanel] = useState(0.85*window.innerWidth);
    const [currResource, setCurrResource] = useState(null);
    const [allUsers, setAllUsers] = useState({});
    const [currUsers, setCurrUsers] = useState({});
    const [user, setUser] = useState({});
    const prefRef = useRef(null);

    //Desk and meeting room prop arrays
    const [deskProps, SetDeskProps] = useState([]);
    const [meetingRoomProps, SetMeetingRoomProps] = useState([]);
    const [wallProps, SetWallProps] = useState([]);
    const [stage, SetStage] = useState({width : 100, height : 100});
    const [selectedId, SelectShape] = useState(null);

    //API fetch variables
    const [buildings, SetBuildings] = useState([]);
    const [rooms, SetRooms] = useState([]);
    const [resources, SetResources] = useState([]);
    const [bookings, setBookings] = useState([]);
    const [currBookings, setCurrBookings] = useState([]);

    const {userData} = useContext(UserContext);

    //POST requests
    const UpdateRooms = (e) =>
    {
        SelectShape(null);
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
        SelectShape(null);
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
        e.target.getStage().container().style.cursor = 'grabbing';
        const clickedEmpty = e.target === e.target.getStage();
        if(clickedEmpty)
        {
            SelectShape(null);
        }
    }

    //Load desks from the database
    const LoadDesk = useCallback((id, name, x, y, width, height, rotation, booked, user) =>
    {
        //Uses a reference array to prevent state dependency and infinite loop
        if(stageRef.current !== null)
        {
            deskPropsRef.current =
            [
                ...deskPropsRef.current,
                {
                    key : id,
                    id : id,
                    name : name,
                    x : x,
                    y : y,
                    width : width,
                    height : height,
                    rotation : rotation,
                    edited : false,
                    booked : booked,
                    user: user
                }
            ];

            //Set the state using the reference array
            SetDeskProps(deskPropsRef.current);
        }
    },[]);

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
                    key : id,
                    id : id,
                    name : name,
                    x : x,
                    y : y,
                    width : width,
                    height : height,
                    rotation : rotation,
                    capacity: capacity,
                    edited : true,
                    booked: false
                }
            ];

            //Set the state using the reference array
            SetMeetingRoomProps(meetingRoomPropsRef.current);
        }
    },[]);

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

    //Adjusts the canvas size for difference screen sizes
    const HandleResize = () =>
    {
        SetStage({width : canvasRef.current.offsetWidth, height : canvasRef.current.offsetHeight});
        if(selectedId)
        {
            setSidePanel(0.65*window.innerWidth);
        }
        else
        {
            setSidePanel(0.85*window.innerWidth);
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

    const setPreference = () =>
    {
        if(preferenceRef.current)
        {
            fetch("http://localhost:8080/api/user/update", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    ...user,
                    preferred_desk: preferenceRef.current.checked ? currResource.id : null,
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                }
            }).then((res) =>
            {
                if(res.status === 200)
                {
                    alert("Preferred Desk Set");
                }

                fetch("http://localhost:8080/api/user/information", 
                {
                    method: "POST",
                    mode: "cors",
                    body: JSON.stringify({
                        identifier : userData.user_identifier
                    }),
                    headers:{
                        'Content-Type': 'application/json',
                        'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                    }
                }).then((res) => res.json()).then(data =>
                {
                    setUser(data[0]);
                });
            });  
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

        fetch("http://localhost:8080/api/user/information", 
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
            //Get info for all users
            var users = {};
            data.forEach((user) =>
            {
                users =  
                {
                    ...users,
                    [user.id]:
                    {
                        id: user.id,
                        name: user.first_name + " " + user.last_name,
                        picture: user.picture,
                        roles: [],
                        teams: []
                    }
                }
            });

            //Get info for all roles
            fetch("http://localhost:8080/api/role/information", 
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
                var roles = {};
                data.forEach((role) =>
                {
                    roles =
                    {
                        ...roles,
                        [role.id]:
                        {
                            id: role.id,
                            name: role.name,
                            color: role.color
                        }
                    }
                });

                //Add role to user based on role association
                fetch("http://localhost:8080/api/role/user/information", 
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
                    data.forEach((roleUser) =>
                    {
                        users[roleUser.user_id].roles.push(roles[roleUser.role_id]);
                    });
                });

                //Get info for all teams
                fetch("http://localhost:8080/api/team/information", 
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
                    var teams = {};
                    data.forEach((team) =>
                    {
                        teams =
                        {
                            ...teams,
                            [team.id]:
                            {
                                id: team.id,
                                name: team.name,
                                color: team.color
                            }
                        }
                    });

                    //Add team to user based on team association
                    fetch("http://localhost:8080/api/team/user/information", 
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
                        data.forEach((teamUser) =>
                        {
                            users[teamUser.user_id].teams.push(teams[teamUser.team_id]);
                        });

                        setAllUsers(users);
                    });
                });
            });
        });

        fetch("http://localhost:8080/api/user/information", 
        {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                identifier : userData.user_identifier
            }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
            }
        }).then((res) => res.json()).then(data =>
        {
            setUser(data[0]);
        });

    }, [userData.token, userData.user_identifier]);

    //Loads desks and meeting rooms from database after room is selected
    useEffect(() =>
    {
        //Reset reference array and counters
        deskPropsRef.current = [];
        meetingRoomPropsRef.current = [];
        wallPropsRef.current = [];

        SetDeskProps(deskPropsRef.current);
        SetMeetingRoomProps(meetingRoomPropsRef.current);
        SetWallProps(wallPropsRef.current);

        //Loop through resources and load desks and meeting rooms respectively
        for(var i = 0; i < resources.length; i++)
        {
            if(resources[i].resource_type === "DESK")
            {
                LoadDesk(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation, false, false);
            }
            else if(resources[i].resource_type === "MEETINGROOM")
            {
                LoadMeetingRoom(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation, JSON.parse(resources[i].decorations).capacity, false);
            }
            else if(resources[i].resource_type === "WALL")
            {
                LoadWall(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation);
            }
        }

    }, [resources, LoadDesk, LoadMeetingRoom, LoadWall]);


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

        if(dateRef.current)
        {
            dateRef.current.value = `${date.getFullYear()}-${month}-${day}`;
        }
    },[]);

    useEffect(() =>
    {
        if(bookings)
        {
            var currBookings = [];
            bookings.forEach((booking) =>
            {
                if(booking.start.includes(date) && booking.booked)
                {
                    currBookings.push(booking);
                }
            });

            setCurrBookings(currBookings);
        }
    },[date, bookings]);

    useEffect(() =>
    {
        if(currBookings)
        {
            var bookedIDs = [];
            currBookings.forEach((booking) =>
            {
                bookedIDs.push(booking.resource_id);
                console.log(booking)
            });

            deskPropsRef.current = [];
            SetDeskProps(deskPropsRef.current);

            for(let i = 0; i < resources.length; i++)
            {
                var booked = false;
                for(let j = 0; j < currBookings.length; j++)
                {
                    if(currBookings[j].resource_id === resources[i].id)
                    {
                        if(resources[i].resource_type === 'DESK' && currBookings[j].user_id === userData.user_id)
                        {
                            LoadDesk(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation, true, true);
                        }
                        else if(resources[i].resource_type === 'DESK')
                        {
                            LoadDesk(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation, true, false);
                        }
                        else if(resources[i].resource_type === 'MEETINGROOM')
                        {
                            LoadMeetingRoom(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation, JSON.parse(resources[i].decorations).capacity, true);
                        }

                        booked = true;
                        break;
                    }
                }

                if(!booked)
                {
                    if(resources[i].resource_type === 'DESK')
                    {
                        LoadDesk(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation, false, false);
                    }
                    else if(resources[i].resource_type === 'MEETINGROOM')
                    {
                        LoadMeetingRoom(resources[i].id, resources[i].name, resources[i].xcoord, resources[i].ycoord, resources[i].width, resources[i].height, resources[i].rotation, JSON.parse(resources[i].decorations).capacity, false);
                    }
                }
                
            }            
        }
    },[currBookings, allUsers, resources, LoadDesk, LoadMeetingRoom, LoadWall]);

    useEffect(() =>
    {
        if(selectedId)
        {
            setSidePanel(0.65*window.innerWidth);
            preferenceRef.current.checked = false;
            resources.forEach((resource) =>
            {
                if(resource.id === selectedId)
                {
                    setCurrResource(resource);

                    if(resource.id === user.preferred_desk && preferenceRef.current)
                    {
                        preferenceRef.current.checked = true;
                    }

                    if(resource.resource_type === 'DESK' && prefRef.current)
                    {
                        prefRef.current.style.display = 'block';
                    }
                    else if(prefRef.current)
                    {
                        prefRef.current.style.display = 'none';
                    }
                }
            })
        }
        else
        {
            setSidePanel(0.85*window.innerWidth);
        }
    },[selectedId, resources, user.preferred_desk]);

    useEffect(() =>
    {
        setCurrUsers({});

        if(currResource)
        {
            currBookings.forEach((booking) =>
            {
                if(booking.resource_id === currResource.id)
                {
                    setCurrUsers((prev) =>
                    ({
                        ...prev, 
                        [booking.user_id]: 
                        {
                            ...allUsers[booking.user_id],
                            start: booking.start.substr(booking.start.indexOf("T", 0) + 1, 5),
                            end: booking.end.substr(booking.end.indexOf("T", 0) + 1, 5)
                        }
                    }));
                }
            });
        }
    },[currResource, currBookings, allUsers]);

    return (
        <Fragment>
            <div className={styles.mapHeadingContainer}>
                <div className={styles.mapHeading}>Office map</div>
            </div>

            <div className={styles.sidePanel} style={{left: sidePanel}}>
                <div className={styles.resourceHeading}>
                    {currResource ? currResource.name : 'Default name'}
                </div>

                <div className={styles.resourceType}>
                    Type: {currResource ? currResource.resource_type : 'Default type'}
                </div>

                <div ref={prefRef} className={styles.resourcePreference}>
                    <input ref={preferenceRef} type='checkbox' onChange={() => setPreference()}></input>
                    <label>Set as preferred desk</label>
                </div>

                <div className={styles.userCardContainer}>
                    {currUsers && Object.entries(currUsers).map(([id, user]) =>
                    (
                        <div className={styles.userCard}>
                            <div className={styles.userPictureContainer}>
                                <img className={styles.image} src={user.picture} alt='user'></img>
                            </div>

                            <div className={styles.userDetailsContainer}>
                                <div className={styles.userName}>{user.name}</div>

                                <div className={styles.userTime}>{user.start} - {user.end}</div>

                                <div className={styles.userRolesContainer}>
                                    <div>Roles:</div>
                                    {user.roles.map((role) =>
                                    (
                                        <div key={role.id} className={styles.pill} style={{backgroundColor: role.color}}>{role.name}</div>                                        
                                    ))}
                                </div>

                                <div className={styles.userTeamsContainer}>
                                    <div>Teams:</div>
                                    {user.teams.map((team) =>
                                    (
                                        <div key={team.id} className={styles.pill} style={{backgroundColor: team.color}}>{team.name}</div>                                        
                                    ))}
                                </div>
                            </div>
                        </div>
                    ))}
                    
                </div>
            </div>

            <div ref={canvasRef} className={styles.canvasContainer}>
                <Stage width={stage.width} height={stage.height} onMouseDown={CheckDeselect} onMouseUp={(e) => e.target.getStage().container().style.cursor = 'grab'} onTouchStart={CheckDeselect} draggable onWheel={ZoomInOut} ref={stageRef}>
                    <Layer>
                        {deskProps.map((desk, i) =>
                        (
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

                                }}

                                draggable = {false}

                                transform = {false}
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
                                }}
                                
                                onChange = {(newProps) => 
                                {
                                    
                                }}

                                draggable = {false}

                                transform = {false}
                            />
                        ))} 

                        {wallProps.map((wall, i) => (
                            <Wall
                                key = {wall.key}
                                shapeProps = {wall}

                                isSelected = {wall.key === selectedId}
                                
                                onSelect = {() => 
                                {
                                    
                                }}
                                
                                onChange = {(newProps) => 
                                {
                                    
                                }}

                                draggable = {false}

                                transform = {false}
                            />
                        ))}                              
                    </Layer>
                </Stage>

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
                                    <option key={room.id} value={room.id}>{room.name + ' (Floor ' + room.zcoord + ')'}</option>
                                ))
                            )}
                    </select>
                </div>

                <div className={styles.dateContainer}>
                    <input ref={dateRef} className={styles.resourceSelector} type='date' onChange={(e) => setDate(e.target.value)}></input>
                </div>
            </div>
        </Fragment>
    )
}

export default Map