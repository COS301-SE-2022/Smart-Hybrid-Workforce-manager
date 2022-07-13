import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import { Stage, Layer, Rect} from 'react-konva'
import { useRef, useState, useEffect } from 'react'

const Layout = () =>
{
    const [position, setPosition] = useState([]);
    const [stage, setStage] = useState({width : 100, height : 100});
    const [count, setCount] = useState(0);
    const canvasRef = useRef(null);

    const AddDesk = () =>
    {
        setPosition(
        [
            ...position,
            {
                id : count,
                isDragging : false,
                x : 0,
                y : 0
            }
        ]);

        setCount(count + 1);
    }

    const handleResize = () =>
    {
        setStage({width : canvasRef.current.offsetWidth, height : canvasRef.current.offsetHeight});
    }

    window.addEventListener('resize', handleResize);

    useEffect(() =>
    {
        setStage({width : canvasRef.current.offsetWidth, height : canvasRef.current.offsetHeight});
    }, [])

    return (
        <div className='page-container'>
            <div className='content'>
                <Navbar />
                <button onClick={AddDesk}>Add Desk</button>
                <div ref={canvasRef} className='canvas-container'>
                    <Stage width={stage.width} height={stage.height} >
                        <Layer>
                            {position.length > 0 && (
                                position.map(desk => (
                                    <Rect
                                        key={desk.id}
                                        width={100}
                                        height={50}
                                        cornerRadius={10}
                                        x={desk.x}
                                        y={desk.y}
                                        fill={desk.isDragging ? "#09A4FB" : "white"}
                                        stroke="black"
                                        draggable
                                        onDragStart={(e) =>
                                        {
                                            const id = e.target.id();
                                            setPosition(
                                                position.map((pos) => 
                                                    {
                                                        return {
                                                            ...pos,
                                                            isDragging : pos.id === id,
                                                        }
                                                    })
                                            );
                                        }}
                                        onDragEnd={(e) =>
                                        {
                                            const id = e.target.id();
                                            setPosition(
                                                position.map((pos) => 
                                                    {
                                                        if(pos.id === id)
                                                        {
                                                            return {
                                                                ...pos,
                                                                isDragging : false,
                                                                x : e.target.x(),
                                                                y : e.target.y()
                                                            }
                                                        }
                                                        else
                                                        {
                                                            return {
                                                                ...pos,
                                                                isDragging : false,
                                                            }
                                                        }
                                                        
                                                    })
                                            );
                                            //setPosition({isDragging : false, x : e.target.x(), y : e.target.y()});
                                        }}
                                        onMouseEnter={(e) =>
                                        {
                                            e.target.getStage().container().style.cursor = 'move';
                                        }}
                                        onMouseLeave={(e) =>
                                        {
                                            e.target.getStage().container().style.cursor = 'default';
                                        }}
                                    />
                                ))
                            )}
                            
                        </Layer>
                    </Stage>
                </div>
            </div>  
            <Footer />
        </div>
    )
}

export default Layout