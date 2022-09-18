import styles from './kanban.module.css';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import { MdEdit, MdPersonAdd } from 'react-icons/md';
import { useRef, useState } from 'react';

const Kanban = () =>
{
    const columnsInit = 
    {
        ['col1']: 
        {
            name: 'Team 1',
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
        }       
    }

    const [columns, setColumns] = useState(columnsInit);

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
            <DragDropContext onDragEnd={result => onDragEnd(result, columns, setColumns)}>
                {Object.entries(columns).map(([id, col]) => {
                    return (
                        <Droppable key={id} droppableId={id}>
                            {(provided, snapshot) =>
                            {
                                return (
                                    <div {...provided.droppableProps} ref={provided.innerRef} className={styles.column}>
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
                                                                height: '20vh',
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
        </div>
    );
}

export default Kanban;