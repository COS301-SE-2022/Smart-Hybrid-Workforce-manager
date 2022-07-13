import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import { Stage, Layer, Rect} from 'react-konva'
import { useRef, useState, useEffect } from 'react'

const Layout = () =>
{
    const [position, setPosition] = useState({isDragging : false, x : 50, y : 50});
    const [stage, setStage] = useState({width : 100, height : 100});
    const canvasRef = useRef(null);

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
                <div ref={canvasRef} className='canvas-container'>
                    <Stage width={stage.width} height={stage.height} >
                        <Layer>
                            <Rect
                                width={100}
                                height={50}
                                x={position.x}
                                y={position.y}
                                fill="white"
                                stroke="black"
                                draggable
                                onDragStart={() =>
                                {
                                    setPosition({isDragging : true});
                                }}
                                onDragEnd={(e) =>
                                {
                                    setPosition({isDragging : false, x : e.target.x(), y : e.target.y()});
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
                        </Layer>
                    </Stage>
                </div>
            </div>  
            <Footer />
        </div>
    )
}

export default Layout