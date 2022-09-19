import styles from './kanban.module.css';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import { MdEdit, MdPersonAdd } from 'react-icons/md';
import { useRef, useState } from 'react';
import { IoIosArrowBack, IoIosArrowForward } from 'react-icons/io';
import { AiOutlineUsergroupAdd } from 'react-icons/ai';
import { FaSave } from 'react-icons/fa';
import { EditTeamForm } from '../Team/EditTeam';

const Kanban = () =>
{
    const columnsContainerRef = useRef(null);
    const rightIntervalRef = useRef(null);
    const leftIntervalRef = useRef(null);

    const columnsInit = 
    {
        ['col1']: 
        {
            name: 'Team 1',
            color: '#09a2fb',
            items: [
                {
                    id: 't1u1',
                    name: 'User 1'
                },
                {
                    id: 't1u2',
                    name: 'User 2'
                }
            ]
        },

        ['col2']:
        {
            name: 'Team 2',
            color: '#ff3e30',
            items: [
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
            items: [
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
            items: [
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

    const ShowAddUserHint = (col) =>
    {
        document.getElementById(col + 'AddUserHint').style.display = 'block';
    }

    const HideAddUserHint = (col) =>
    {
        document.getElementById(col + 'AddUserHint').style.display = 'none';
    }

    const EditTeam = (col) =>
    {
        console.log(col);
    }

    const ShowEditTeamHint = (col) =>
    {
        document.getElementById(col + 'EditTeamHint').style.display = 'block';

    }

    const HideEditTeamHint = (col) =>
    {
        document.getElementById(col + 'EditTeamHint').style.display = 'none';
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

            const sourceItems = [...sourceColumn.items]; //Copies items from current column
            const [removed] = sourceItems.splice(source.index, 1); //Removes item from the source index

            const destinationItems = [...destinationColumn.items]; //Copies items from new column
            destinationItems.splice(destination.index, 0, removed); //Adds item to the destination index

            setColumns({
                ...columns,
                [source.droppableId]:
                {
                    ...sourceColumn,
                    items: sourceItems
                },

                [destination.droppableId]:
                {
                    ...destinationColumn,
                    items: destinationItems
                }
            });
        }
        else
        {
            const {source, destination} = result; //Source and destination is position in column
            const column = columns[source.droppableId]; //Gets current column
            const copiedItems = [...column.items]; //Copies items from current column
            const [removed] = copiedItems.splice(source.index, 1); //Removes item from the source index
            copiedItems.splice(destination.index, 0, removed); //Adds item to the destination index
            setColumns({
                ...columns,
                [source.droppableId]:
                {
                    ...column,
                    items: copiedItems
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

            <div className={styles.backgroundDimmer}></div>

            <div className={styles.editTeamContainer}>
                <EditTeamForm teamName='Team 1' teamColor='#86ff30' />
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
                                                <div className={styles.columnHeader}>
                                                    {col.name}
                                                </div>
                                                <div className={styles.columnActions}>
                                                    <div className={styles.addUser} onClick={AddUser.bind(this, id)}><MdPersonAdd onMouseEnter={ShowAddUserHint.bind(this, id)} onMouseLeave={HideAddUserHint.bind(this, id)} /></div>
                                                    <div className={styles.editTeam} onClick={EditTeam.bind(this, id)}><MdEdit onMouseEnter={ShowEditTeamHint.bind(this, id)} onMouseLeave={HideEditTeamHint.bind(this, id)}/></div>
                                                </div>

                                                <div id={id + 'AddUserHint'} className={styles.addUserHint}>Add a new user</div>
                                                <div id={id + 'EditTeamHint'} className={styles.editTeamHint}>Edit team</div>
                                            </div>

                                            <div className={styles.itemsContainer}>
                                                {col.items.length > 0 && col.items.map((item, index) => (
                                                    <Draggable key={item.id} draggableId={item.id} index={index}>
                                                        {(provided, snapshot) =>
                                                        {
                                                            return (
                                                                <div {...provided.draggableProps} {...provided.dragHandleProps} ref={provided.innerRef}
                                                                style={{
                                                                    userSelect: 'none',
                                                                    paddingTop: '2vh',
                                                                    paddingLeft: '1vw',
                                                                    marginBottom: '3vh',
                                                                    height: '15vh',
                                                                    width: '18vw',
                                                                    borderRadius: '1vh',
                                                                    boxShadow: '0 4px 10px rgba(0, 0, 0, 0.2)',
                                                                    backgroundColor: snapshot.isDragging ? '#09a2fb55' : 'white',
                                                                    ...provided.draggableProps.style
                                                                }}>
                                                                    {item.name}
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