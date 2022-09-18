import styles from './kanban.module.css';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';

const Kanban = () =>
{
    const columns = 
    [
        {
            id: 'col1',
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

        {
            id: 'col2',
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
    ]

    const onDragEnd = () =>
    {
        console.log('E');
    }

    return (
        <div className={styles.kanbanContainer}>
            <DragDropContext onDragEnd={onDragEnd}>
                {columns.length > 0 && columns.map(col => (
                    <Droppable key={col.id} droppableId={col.id}>
                        {(provided, snapshot) =>
                        {
                            return (
                                <div {...provided.droppableProps} ref={provided.innerRef}
                                style={{
                                    backgroundColor: snapshot.isDraggingOver ? 'rgba(218, 223, 254, 0.8)' : 'white',
                                    paddingTop: '2vh',
                                    paddingBottom: '5vh',
                                    width: '20vw',
                                    minHeight: '90vh',
                                    marginLeft: '3vw',
                                }}>
                                    <div className={styles.columnHeaderContainer}>
                                        <div className={styles.columnHeader}>
                                            {col.name}
                                        </div>
                                        <div className={styles.columnActions}>

                                        </div>
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
                                                            marginTop: '3vh',
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
                ))}
            </DragDropContext>
        </div>
    );
}

export default Kanban;