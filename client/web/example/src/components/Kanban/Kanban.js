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
                                    backgroundColor: snapshot.isDraggingOver ? 'lightblue' : 'lightgray',
                                    padding: 4,
                                    width: '20vw',
                                    height: '90vh',
                                    marginLeft: '1vw',
                                }}>
                                    {col.items.length > 0 && col.items.map((item, index) => (
                                        <Draggable key={item.id} draggableId={item.id} index={index}>
                                            {(provided, snapshot) =>
                                            {
                                                return (
                                                    <div {...provided.draggableProps} {...provided.dragHandleProps} ref={provided.innerRef}
                                                    style={{
                                                        userSelect: 'none',
                                                        padding: 16,
                                                        margin: '0 0 8px 0',
                                                        minHeight: 50,
                                                        backgroundColor: snapshot.isDragging ? 'red' : 'white',
                                                        ...provided.draggableProps.style
                                                    }}>
                                                        {item.name}
                                                    </div>
                                                )
                                            }}
                                        </Draggable>
                                    ))}
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