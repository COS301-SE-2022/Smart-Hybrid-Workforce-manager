import styles from './kanban.module.css';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import { MdEdit, MdPersonAdd, MdClose } from 'react-icons/md';
import { useEffect, useRef, useState } from 'react';
import { IoIosArrowBack, IoIosArrowForward } from 'react-icons/io';
import { AiOutlineUsergroupAdd } from 'react-icons/ai';
import { FaSave } from 'react-icons/fa';
import { BsThreeDotsVertical } from 'react-icons/bs';
import { EditTeamForm } from '../Team/EditTeam';
import { AddTeamForm } from '../Team/AddTeam';

const Kanban = () =>
{
    const columnsContainerRef = useRef(null);
    const rightIntervalRef = useRef(null);
    const leftIntervalRef = useRef(null);

    const [editTeamName, setEditTeamName] = useState('Default');
    const [editTeamColor, setEditTeamColor] = useState('#ffffff');
    const [editTeamPicture, setEditTeamPicture] = useState('');
    const [currTeam, setCurrTeam] = useState('');

    const columnsInit = 
    {
        ['col1']: 
        {
            name: 'Team 1',
            color: '#09a2fb',
            picture: 'https://seeklogo.com/images/B/breaking-bad-logo-032E797C11-seeklogo.com.jpg',
            users: [
                {
                    id: 't1u1',
                    name: 'Walter White',
                    picture: 'https://upload.wikimedia.org/wikipedia/en/thumb/0/03/Walter_White_S5B.png/220px-Walter_White_S5B.png'
                },
                {
                    id: 't1u2',
                    name: 'Jesse Pinkman',
                    picture: 'https://upload.wikimedia.org/wikipedia/en/thumb/c/c6/Jesse_Pinkman_S5B.png/220px-Jesse_Pinkman_S5B.png'
                }
            ]
        },

        ['col2']:
        {
            name: 'Team 2',
            color: '#ff3e30',
            users: [
                {
                    id: 't2u1',
                    name: 'User 1'
                },
                {
                    id: 't2u2',
                    name: 'User 2'
                },
                {
                    id: 't2u3',
                    name: 'User 3'
                },
                {
                    id: 't2u4',
                    name: 'User 4'
                }
            ]
        },
        
        ['col3']:
        {
            name: 'Team 3',
            color: '#86ff30',
            users: [
                {
                    id: 't3u1',
                    name: 'User 1'
                },
                {
                    id: 't3u2',
                    name: 'User 2'
                },
                {
                    id: 't3u3',
                    name: 'User 3'
                },
                {
                    id: 't3u4',
                    name: 'User 4'
                },
                {
                    id: 't3u5',
                    name: 'User 5'
                },
                {
                    id: 't3u6',
                    name: 'User 6'
                },
                {
                    id: 't3u7',
                    name: 'User 7'
                }
            ]
        },

        ['col4']:
        {
            name: 'Team 4',
            color: '#d230ff',
            users: [
                {
                    id: 't4u1',
                    name: 'User 1'
                },
                {
                    id: 't4u2',
                    name: 'User 2'
                },
                {
                    id: 't4u3',
                    name: 'User 3'
                },
                {
                    id: 't4u4',
                    name: 'User 4'
                },
                {
                    id: 't4u5',
                    name: 'User 5'
                },
                {
                    id: 't4u6',
                    name: 'User 6'
                },
                {
                    id: 't4u7',
                    name: 'User 7'
                }
            ]
        }
    }

    const [columns, setColumns] = useState(columnsInit);

    const ShowSaveHint = () =>
    {
        document.getElementById('SaveHint').style.display = 'block';
    }

    const HideSaveHint = () =>
    {
        document.getElementById('SaveHint').style.display = 'none';
    }

    const AddUser = (col) =>
    {
        console.log(col);
    }

    const EditTeam = (col) =>
    {
        setEditTeamName(columns[col].name);
        setEditTeamColor(columns[col].color);
        setEditTeamPicture(columns[col].picture);

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
        document.getElementById('BackgroundDimmer').style.display = 'block';
        document.getElementById('AddTeam').style.display = 'block';
    }

    const CloseAddTeam = () =>
    {
        document.getElementById('BackgroundDimmer').style.display = 'none';
        document.getElementById('AddTeam').style.display = 'none';
    }

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

    const ShowUserMenu = () =>
    {
        window.alert("Yo");
    }

    const ShowTeamMenu = (id) =>
    {
        if(document.getElementById('TeamMenu').style.display === 'none')
        {
            document.getElementById('TeamMenu').style.display = 'block';
            document.getElementById('TeamMenu').style.left = document.getElementById(id + 'ColumnActions').getBoundingClientRect().left - 0.22*window.innerWidth + 'px';
            document.getElementById('TeamMenu').style.top = document.getElementById(id + 'ColumnActions').getBoundingClientRect().top - 0.10*window.innerHeight + 'px';
            setCurrTeam(id);
        }
        else
        {
            document.getElementById('TeamMenu').style.display = 'none';
        }
    }

    document.addEventListener('mousedown', function ClickOutside(event)
    {
        if(document.getElementById('TeamMenu').style.display === 'block')
        {
            document.getElementById('TeamMenu').style.display = 'none';
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
                <div className={styles.editTeam} onMouseDown={EditTeam.bind(this, currTeam)}>Edit team</div>
                <div className={styles.deleteTeam}>Delete team</div>
            </div>

            <div id='BackgroundDimmer' className={styles.backgroundDimmer}></div>

            <div id='EditTeam' className={styles.formTeamContainer}>
                <div className={styles.formTeamClose} onClick={CloseEditTeam}><MdClose /></div>
                <EditTeamForm teamName={editTeamName} teamColor={editTeamColor} teamPriority={3} teamPicture={editTeamPicture} />
            </div>

            <div id='AddTeam' className={styles.formTeamContainer}>
                <div className={styles.formTeamClose} onClick={CloseAddTeam}><MdClose /></div>
                <AddTeamForm />
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
                                                        <BsThreeDotsVertical className={styles.menu} onMouseUp={ShowTeamMenu.bind(this, id)} />
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

                                                {col.users.length > 0 && col.users.map((user, index) => (
                                                    <Draggable key={user.id} draggableId={user.id} index={index}>
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
                                                                            <div className={styles.role}>Developer</div>
                                                                            <div className={styles.role}>Secretary</div>
                                                                            <div className={styles.role}>Team Lead</div>
                                                                            <div className={styles.role}>CEO</div>
                                                                        </div>
                                                                    </div>

                                                                    <div className={styles.userMenuContainer}>
                                                                        <BsThreeDotsVertical className={styles.menu} onClick={ShowUserMenu} />
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

                <div className={styles.addColumn}>
                    <div className={styles.addTeamContainer} onClick={AddTeam}>
                        <AiOutlineUsergroupAdd />
                        Add team
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Kanban;