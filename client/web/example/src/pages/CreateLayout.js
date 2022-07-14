import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import { Stage, Layer, Rect, Transformer} from 'react-konva'
import { useRef, useState, useEffect, Fragment } from 'react'

const Layout = () =>
{
    const [position, setPosition] = useState([]);
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
    const canvasRef = useRef(null);

    const AddDesk = () =>
    {
        setPosition(
        [
            ...position,
            {
                key : "desk" + count,
                cornerRadius : 10,
                x : 0,
                y : 0,
                width : 100,
                height : 50,
                fill : "white",
                stroke : "black",
                rotation : 0
            }
        ]);

        setCount(count + 1);
    }

    const Rectangle = ({ shapeProps, isSelected, onSelect, onChange}) =>
    {
        const shapeRef = useRef(null);
        const transformRef = useRef(null);

        useEffect(() =>
        {
            if(isSelected)
            {
                transformRef.current.nodes([shapeRef.current]);
                transformRef.current.getLayer().batchDraw();
            }
        }, [isSelected]);

        return (
            <Fragment>
                <Rect
                    onClick={onSelect}
                    onTap={onSelect}
                    ref={shapeRef}
                    {...shapeProps}
                    {...console.log(position)}
                    draggable

                    onDragEnd={(e) =>
                    {
                        onChange({
                            ...shapeProps,
                            x : e.target.x(),
                            y : e.target.y()
                        })
                    }}

                    onTransformEnd={(e) =>
                    {
                        onChange({
                            ...shapeProps,
                            x : e.target.x(),
                            y : e.target.y(),
                            rotation : e.target.rotation()
                        });
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
                
                {isSelected && (
                    <Transformer 
                        ref = {transformRef}
                        rotationSnaps = {[0, 90, 180, 270]}
                        resizeEnabled = {false}
                        centeredScaling = {true}
                    />
                )}
            </Fragment>
        );
    };

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
                    <Stage width={stage.width} height={stage.height} onMouseDown={checkDeselect} onTouchStart={checkDeselect}>
                        <Layer>
                            {position.length > 0 && (
                                position.map((desk, i) => (
                                    <Rectangle
                                        key = {desk.key}
                                        shapeProps = {desk}

                                        isSelected = {desk.key === selectedId}
                                        
                                        onSelect = {() => 
                                        {
                                            selectShape(desk.key);
                                        }}
                                        
                                        onChange = {(newPos) => 
                                        {
                                            const positions = position.slice();
                                            positions[i] = newPos;
                                            setPosition(positions)
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