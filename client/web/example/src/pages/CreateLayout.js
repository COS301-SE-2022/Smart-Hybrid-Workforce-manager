import { Stage, Layer } from 'react-konva'
import { useRef, useState, useEffect, useCallback } from 'react'
import Desk from '../components/Map/Desk'
import MeetingRoom from '../components/Map/MeetingRoom'

const Layout = () =>
{
    const canvasRef = useRef(null);
    const stageRef = useRef(null);
    const scaleFactor = 1.3;

    const [deskProps, setDeskProps] = useState([]);
    const [meetingRoomProps, setMeetingRoomProps] = useState([]);
    const [stage, setStage] = useState({width : 100, height : 100});
    const [count, setCount] = useState(0);
    const [selectedId, selectShape] = useState(null);

    const checkDeselect = (e) =>
    {
        const clickedEmpty = e.target === e.target.getStage();
        if(clickedEmpty)
        {
            selectShape(null);
        }
    }

    const AddDesk = () =>
    {
        setDeskProps(
        [
            ...deskProps,
            {
                key : "desk" + count,
                x : (-stageRef.current.x() + stageRef.current.width() / 2.0) / stageRef.current.scaleX(),
                y : (-stageRef.current.y() + stageRef.current.height() / 2.0) / stageRef.current.scaleY(),
                width : 60,
                height : 55,
                rotation : 0
            }
        ]);

        setCount(count + 1);
    }

    const AddMeetingRoom = () =>
    {
        setMeetingRoomProps(
        [
            ...meetingRoomProps,
            {
                key : "meetingroom" + count,
                x : (-stageRef.current.x() + stageRef.current.width() / 2.0) / stageRef.current.scaleX(),
                y : (-stageRef.current.y() + stageRef.current.height() / 2.0) / stageRef.current.scaleY(),
                width : 200,
                height : 200,
                rotation : 0
            }
        ]);

        setCount(count + 1);
    }

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
        }, [downHandler, upHandler]); // Empty array ensures that effect is only run on mount and unmount
        return keyPressed;
    }

    const handleDelete = useCallback(() =>
    {
        if(selectedId !== null)
        {
            if(selectedId.includes("desk"))
            {
                for(var i = 0; i < deskProps.length; i++)
                {
                    if(deskProps[i].key === selectedId)
                    {
                        var newDesk = [...deskProps];
                        newDesk.splice(i, 1);
                        setDeskProps(newDesk);
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
    }, [deskProps, meetingRoomProps, selectedId])

    const handleResize = () =>
    {
        setStage({width : canvasRef.current.offsetWidth, height : canvasRef.current.offsetHeight});
    }

    window.addEventListener('resize', handleResize);

    const canvasDrag = () =>
    {

    }

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

    const SaveLayout = () =>
    {
        window.alert("Saved");
    }

    useEffect(() =>
    {
        setStage({width : canvasRef.current.offsetWidth, height : canvasRef.current.offsetHeight});
        
        if(deletePressed)
        {
            handleDelete();
        }

    }, [deletePressed, handleDelete])

    return (
        <div className='page-container'>
            <div className='canvas-content'>
                <button onClick={AddDesk}>Add Desk</button><br></br>
                <button onClick={AddMeetingRoom}>Add Meeting Room</button>
                <div ref={canvasRef} className='canvas-container'>
                    <Stage width={stage.width} height={stage.height} onMouseDown={checkDeselect} onTouchStart={checkDeselect} draggable onDragEnd={canvasDrag} onWheel={zoomInOut} ref={stageRef}>
                        <Layer>
                            {deskProps.length > 0 && (
                                deskProps.map((desk, i) => (
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
                                            const newDeskProps = deskProps.slice();
                                            newDeskProps[i] = newProps;
                                            setDeskProps(newDeskProps)
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