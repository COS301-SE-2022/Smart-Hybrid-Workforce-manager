import styles from './kanban.module.css';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import { MdEdit, MdPersonAdd, MdClose } from 'react-icons/md';
import { useContext, useEffect, useRef, useState } from 'react';
import { IoIosArrowBack, IoIosArrowForward } from 'react-icons/io';
import { AiOutlineUsergroupAdd } from 'react-icons/ai';
import { FaSave } from 'react-icons/fa';
import { BsThreeDotsVertical } from 'react-icons/bs';
import { EditTeamForm } from '../Team/EditTeam';
import { AddTeamForm } from '../Team/AddTeam';
import { EditUserPanel } from '../User/EditUser';
import { UserContext } from '../../App';

const Kanban = () =>
{
    const columnsContainerRef = useRef(null);
    const rightIntervalRef = useRef(null);
    const leftIntervalRef = useRef(null);

    const [columns, setColumns] = useState({});

    const [editTeam, setEditTeam] = useState(null);
    const [teamEdited, setTeamEdited] = useState(true);
    const [addTeam, setAddTeam] = useState(0);

    const [currUser, setCurrUser] = useState({id: '', name: '', picture: ''});
    const [userPanelLeft, setUserPanelLeft] = useState(0.85*window.innerWidth);

    const {userData} = useContext(UserContext);

    const [roles, setRoles] = useState([]);

    const ShowSaveHint = () =>
    {
        document.getElementById('SaveHint').style.display = 'block';
    }

    const HideSaveHint = () =>
    {
        document.getElementById('SaveHint').style.display = 'none';
    }

    //Teams
    const ShowTeamMenu = (col) =>
    {
        if(document.getElementById('TeamMenu').style.display === 'none')
        {
            document.getElementById('TeamMenu').style.display = 'block';
            document.getElementById('TeamMenu').style.left = document.getElementById(col.id + 'ColumnActions').getBoundingClientRect().left - 0.22*window.innerWidth + 'px';
            document.getElementById('TeamMenu').style.top = document.getElementById(col.id + 'ColumnActions').getBoundingClientRect().top - 0.10*window.innerHeight + 'px';
            setEditTeam(col);
        }
        else
        {
            document.getElementById('TeamMenu').style.display = 'none';
        }
    }

    const EditTeam = () =>
    {
        document.getElementById('BackgroundDimmer').style.display = 'block';
        document.getElementById('EditTeam').style.display = 'block';
    }

    const CloseEditTeam = () =>
    {
        document.getElementById('BackgroundDimmer').style.display = 'none';
        document.getElementById('EditTeam').style.display = 'none';
    }

    const AddTeam = () =>
    {
        if(addTeam === 0)
        {
            setAddTeam(1);
        }
        else
        {
            setAddTeam(0);
        }

        document.getElementById('BackgroundDimmer').style.display = 'block';
        document.getElementById('AddTeam').style.display = 'block';
    }

    const CloseAddTeam = () =>
    {
        document.getElementById('BackgroundDimmer').style.display = 'none';
        document.getElementById('AddTeam').style.display = 'none';
    }

    //Navigation
    const StartScrollLeft = () =>
    {
        if(columnsContainerRef.current)
        {
            leftIntervalRef.current = setInterval(() =>
            {
                columnsContainerRef.current.scrollLeft -= 10;
            }, 10);
        }
    }

    const StopScrollLeft = () =>
    {
        if(leftIntervalRef.current)
        {
            clearInterval(leftIntervalRef.current);
            leftIntervalRef.current = null;
        }
    }

    const StartScrollRight = () =>
    {
        if(columnsContainerRef.current)
        {
            rightIntervalRef.current = setInterval(() =>
            {
                columnsContainerRef.current.scrollLeft += 10;
            }, 10);
        }
    }

    const StopScrollRight = () =>
    {
        if(rightIntervalRef.current)
        {
            clearInterval(rightIntervalRef.current);
            rightIntervalRef.current = null;
        }
    }

    //Users
    const ShowUserMenu = (user) =>
    {
        if(document.getElementById('UserMenu').style.display === 'none')
        {
            document.getElementById('UserMenu').style.display = 'block';
            document.getElementById('UserMenu').style.left = document.getElementById(user.id + 'UserActions').getBoundingClientRect().left - 0.23*window.innerWidth + 'px';
            document.getElementById('UserMenu').style.top = document.getElementById(user.id + 'UserActions').getBoundingClientRect().top - 0.10*window.innerHeight + 'px';
            setCurrUser(user);
        }
        else
        {
            document.getElementById('UserMenu').style.display = 'none';
        }
    }

    const AddUser = (col) =>
    {
        document.getElementById('BackgroundDimmer').style.display = 'block';
        document.getElementById('EditUser').style.display = 'block';
    }

    const EditUser = (col) =>
    {
        document.getElementById('BackgroundDimmer').style.display = 'block';
        setUserPanelLeft(0.65*window.innerWidth);
    }

    const CloseUserPanel = () =>
    {
        document.getElementById('BackgroundDimmer').style.display = 'none';
        setUserPanelLeft(0.85*window.innerWidth);
    }

    document.addEventListener('mousedown', function ClickOutside(event)
    {
        if(document.getElementById('TeamMenu').style.display === 'block')
        {
            document.getElementById('TeamMenu').style.display = 'none';
        }

        if(document.getElementById('UserMenu').style.display === 'block')
        {
            document.getElementById('UserMenu').style.display = 'none';
        }
    });

    const onDragEnd = (result, columns, setColumns) =>
    {
        if(!result.destination)
        {
            return;
        }

        const {source, destination} = result; //Source and destination is position in column

        if(source.droppableId !== destination.droppableId)
        {
            const sourceColumn = columns[source.droppableId]; //Gets current column
            const destinationColumn = columns[destination.droppableId]; //Gets new column

            const sourceItems = [...sourceColumn.users]; //Copies items from current column
            const [removed] = sourceItems.splice(source.index, 1); //Removes item from the source index

            const destinationItems = [...destinationColumn.users]; //Copies items from new column
            destinationItems.splice(destination.index, 0, removed); //Adds item to the destination index

            setColumns({
                ...columns,
                [source.droppableId]:
                {
                    ...sourceColumn,
                    users: sourceItems
                },

                [destination.droppableId]:
                {
                    ...destinationColumn,
                    users: destinationItems
                }
            });
        }
        else
        {
            const {source, destination} = result; //Source and destination is position in column
            const column = columns[source.droppableId]; //Gets current column
            const copiedItems = [...column.users]; //Copies items from current column
            const [removed] = copiedItems.splice(source.index, 1); //Removes item from the source index
            copiedItems.splice(destination.index, 0, removed); //Adds item to the destination index
            setColumns({
                ...columns,
                [source.droppableId]:
                {
                    ...column,
                    users: copiedItems
                }
            });
        }
    }

    const GetData = () =>
    {
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
            let teams = {};
            data.forEach(function CreateTeamsObject(team)
            {
                teams = 
                {
                    ...teams,
                    [team.id]:
                    {
                        id: team.id,
                        name: team.name,
                        color: team.color,
                        picture: team.picture,
                        priority: team.priority,
                        lead: team.lead,
                        users:
                        [

                        ]
                    }
                }
            })

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
                let usersData = [];
                data.forEach(function CreateUsersObject(user)
                {
                    usersData.push(
                        {
                            id: user.id,
                            name: user.first_name + ' ' + user.last_name,
                            picture: user.picture,
                            roles:
                            [

                            ]
                        }
                    );
                });

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
                    let rolesData = {};
                    let rolesForEdit = [];
                    data.forEach(function CreateRolesObject(role)
                    {
                        rolesForEdit.push(role.name);
                        rolesData =
                        {
                            ...rolesData,
                            [role.id]:
                            {
                                name: role.name,
                                color: role.color
                            }
                        };
                    });

                    setRoles(rolesForEdit);

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
                        data.forEach(function AddRoles(role)
                        {
                            for(var i = 0; i < usersData.length; i++)
                            {
                                if(usersData[i].id === role.user_id)
                                {
                                    usersData[i].roles.push(rolesData[role.role_id].name);
                                    break;
                                }
                            }
                        });

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
                            data.forEach(function AddTeamMembers(teamMember)
                            {
                                for(var i = 0; i < usersData.length; i++)
                                {
                                    if(usersData[i].id === teamMember.user_id)
                                    {
                                        teams[teamMember.team_id].users.push(usersData[i]);
                                        break;
                                    }
                                }
                            });

                            setColumns(teams);
                        });
                    });
                });
            });
        });
    }

    useEffect(() =>
    {
        if(teamEdited)
        {
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
                let teams = {};
                data.forEach(function CreateTeamsObject(team)
                {
                    teams = 
                    {
                        ...teams,
                        [team.id]:
                        {
                            id: team.id,
                            name: team.name,
                            color: team.color,
                            capacity: team.capacity,
                            picture: team.picture,
                            priority: team.priority,
                            lead: team.lead,
                            users:
                            [

                            ]
                        }
                    }
                })

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
                    let usersData = [];
                    data.forEach(function CreateUsersObject(user)
                    {
                        usersData.push(
                            {
                                id: user.id,
                                name: user.first_name + ' ' + user.last_name,
                                picture: user.picture,
                                roles:
                                [

                                ]
                            }
                        );
                    });

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
                        let rolesData = {};
                        let rolesForEdit = [];
                        data.forEach(function CreateRolesObject(role)
                        {
                            rolesForEdit.push(role.name);
                            rolesData =
                            {
                                ...rolesData,
                                [role.id]:
                                {
                                    name: role.name,
                                    color: role.color
                                }
                            };
                        });

                        setRoles(rolesForEdit);

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
                            data.forEach(function AddRoles(role)
                            {
                                for(var i = 0; i < usersData.length; i++)
                                {
                                    if(usersData[i].id === role.user_id)
                                    {
                                        usersData[i].roles.push(rolesData[role.role_id].name);
                                        break;
                                    }
                                }
                            });

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
                                data.forEach(function AddTeamMembers(teamMember)
                                {
                                    for(var i = 0; i < usersData.length; i++)
                                    {
                                        if(usersData[i].id === teamMember.user_id)
                                        {
                                            teams[teamMember.team_id].users.push(usersData[i]);
                                            break;
                                        }
                                    }
                                });

                                setColumns(teams);
                                setTeamEdited(false);
                            });
                        });
                    });
                });
            });
        }

    }, [userData.token, teamEdited]);

    return (
        <div className={styles.kanbanContainer}>
            <div className={styles.kanbanHeadingContainer}>
                <div className={styles.kanbanHeading}>Team and User Management</div>
            </div>

            <div className={styles.saveIcon} onMouseEnter={ShowSaveHint} onMouseLeave={HideSaveHint}><FaSave /></div>
            <div id='SaveHint' className={styles.saveHint}>Save</div>

            <div className={styles.leftArrow} onMouseDown={StartScrollLeft} onMouseUp={StopScrollLeft} onMouseLeave={StopScrollLeft}><IoIosArrowBack /></div>
            <div className={styles.rightArrow} onMouseDown={StartScrollRight} onMouseUp={StopScrollRight} onMouseLeave={StopScrollRight}><IoIosArrowForward /></div>

            <div id='TeamMenu' className={styles.teamMenu}>
                <div className={styles.editTeam} onMouseDown={EditTeam.bind(this)}>Edit team</div>
                <div className={styles.deleteTeam}>Delete team</div>
            </div>

            <div id='UserMenu' className={styles.userMenu}>
                <div className={styles.editUser} onMouseDown={EditUser.bind(this, currUser)}>Edit user</div>
                <div className={styles.deleteUser}>Remove user</div>
            </div>

            <div id='BackgroundDimmer' className={styles.backgroundDimmer}></div>

            <div id='AddTeam' className={styles.formContainer}>
                <div className={styles.formClose} onClick={CloseAddTeam}><MdClose /></div>
                <AddTeamForm makeDefault={addTeam} />
            </div>

            <div id='EditTeam' className={styles.formContainer}>
                <div className={styles.formClose} onClick={CloseEditTeam}><MdClose /></div>
                <EditTeamForm team={editTeam} edited={setTeamEdited} />
            </div>

            <div id='EditUser' className={styles.userPanel} style={{left: userPanelLeft}}>
                <div className={styles.userPanelClose} onClick={CloseUserPanel}><MdClose /></div>
                <EditUserPanel userName={currUser.name} userPicture={currUser.picture} userRoles={currUser.roles} allRoles={roles} />
            </div>
            

            <div ref={columnsContainerRef} className={styles.columnsContainer}>
                <DragDropContext onDragEnd={result => onDragEnd(result, columns, setColumns)}>
                    {Object.entries(columns).map(([id, col]) => {
                        return (                             
                            <Droppable key={id} droppableId={id}>
                                {(provided, snapshot) =>
                                {
                                    return (
                                        <div {...provided.droppableProps} ref={provided.innerRef} className={styles.column}
                                        style={{
                                            background: "linear-gradient(180deg, " + col.color + "66  0%, rgba(255,255,255,0.4) 20%)"
                                        }}>
                                            <div className={styles.columnHeaderContainer}>
                                                <div className={styles.columnHeaderTop}>
                                                    <div className={styles.columnPicture}>
                                                        <img className={styles.image} src={col.picture} alt='Team'></img>
                                                    </div>

                                                    <div id={id + 'ColumnActions'} className={styles.columnActions}>
                                                        <BsThreeDotsVertical className={styles.menu} onMouseUp={ShowTeamMenu.bind(this, col)} />
                                                    </div>
                                                </div>
                                                
                                                <div className={styles.columnHeader}>
                                                    {col.name}
                                                </div>
                                            </div>

                                            <div className={styles.itemsContainer}>
                                                <div className={styles.addUser}>
                                                    <div className={styles.addUserContainer} onClick={AddUser.bind(this, id)}>
                                                        <AiOutlineUsergroupAdd />
                                                        Add user
                                                    </div>
                                                </div>

                                                {col.users.map((user, index) => (
                                                    <Draggable key={user.id} draggableId={id + user.id} index={index}>
                                                        {(provided, snapshot) =>
                                                        {
                                                            return (
                                                                <div {...provided.draggableProps} {...provided.dragHandleProps} ref={provided.innerRef} className={styles.userCard}
                                                                style={{
                                                                    backgroundColor: snapshot.isDragging ? '#09a2fb55' : 'white',
                                                                    ...provided.draggableProps.style
                                                                }}>
                                                                    <div className={styles.userPictureContainer}>
                                                                        <img className={styles.image} src={user.picture} alt='user'></img>
                                                                    </div>

                                                                    <div className={styles.userDetailsContainer}>
                                                                        <div className={styles.userName}>{user.name}</div>
                                                                        <div className={styles.userRolesContainer}>
                                                                            {user.roles.map((role) =>
                                                                            {
                                                                                return (
                                                                                    <div key={role} className={styles.role}>{role}</div>
                                                                                );
                                                                                
                                                                            })}
                                                                        </div>
                                                                    </div>

                                                                    <div className={styles.userMenuContainer}>
                                                                        <BsThreeDotsVertical id={user.id + 'UserActions'} className={styles.menu} onMouseUp={ShowUserMenu.bind(this, user)} />
                                                                    </div>
                                                                </div>
                                                            )
                                                        }}
                                                    </Draggable>
                                                ))}
                                            </div>
                                            
                                            {provided.placeholder}
                                        </div>
                                    )
                                }}
                            </Droppable>
                        )
                    })}
                </DragDropContext>

                <div className={styles.addColumn} onClick={AddTeam}>
                    <div className={styles.addTeamContainer}>
                        <AiOutlineUsergroupAdd />
                        Add team
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Kanban;