import { Path, Transformer } from 'react-konva';
import { useRef, useEffect, Fragment, useState } from 'react';

const MeetingRoom = ({ shapeProps, isSelected, onSelect, onChange, draggable, transform}) =>
{
    const shapeRef = useRef(null);
    const transformRef = useRef(null);
    const [booked, setBooked] = useState(shapeProps.booked);
    const [width, setWidth] = useState(shapeProps.width);
    const [height, setHeight] = useState(shapeProps.height);

    useEffect(() =>
    {
        if(isSelected && transform)
        {
            transformRef.current.nodes([shapeRef.current]);
            transformRef.current.getLayer().batchDraw();
        }

        if(!isSelected && booked)
        {
            shapeRef.current.fill('#7780e4');
        }
        else if(!isSelected && !booked)
        {
            shapeRef.current.fill('#e8e8e8');
        }
    }, [isSelected, transform, booked]);

    useEffect(() =>
    {
        if(booked)
        {
            shapeRef.current.fill('#7780e4');
        }
        else
        {
            shapeRef.current.fill('#e8e8e8');
        }
    },[booked])

    useEffect(() =>
    {
        setBooked(shapeProps.booked);
    },[shapeProps.booked]);

    return (
        <Fragment> 
            <Path
                {...shapeProps}
                ref={shapeRef}

                data = 'M 0 0 h 600 a 20 20 0 0 1 20 20 v 80 a 20 20 0 0 1 -20 20 h -600 a 20 20 0 0 1 -20 -20 v -80 a 20 20 0 0 1 20 -20 Z M 50 -10 h 10 v -10 h 80 v 10 h 10 v -20 a 10 10 0 0 0 -100 0 Z m 133 0 h 10 v -10 h 80 v 10 h 10 v -20 a 10 10 0 0 0 -100 0 Z m 133 0 h 10 v -10 h 80 v 10 h 10 v -20 a 10 10 0 0 0 -100 0 Z m 133 0 h 10 v -10 h 80 v 10 h 10 v -20 a 10 10 0 0 0 -100 0 Z M 50 130 h 10 v 10 h 80 v -10 h 10 v 20 a 10 10 0 0 1 -100 0 Z m 133 0 h 10 v 10 h 80 v -10 h 10 v 20 a 10 10 0 0 1 -100 0 Z m 133 0 h 10 v 10 h 80 v -10 h 10 v 20 a 10 10 0 0 1 -100 0 Z m 133 0 h 10 v 10 h 80 v -10 h 10 v 20 a 10 10 0 0 1 -100 0 Z M -30 10 v 10 h -10 v 80 h 10 v 10 h -20 a 10 10 0 0 1 0 -100 Z M 630 10 v 10 h 10 v 80 h -10 v 10 h 20 a 10 10 0 0 0 0 -100 Z'

                fill='#e8e8e8'

                scaleX={0.3}
                scaleY={0.3}

                draggable = {draggable}

                onMouseDown={(e) =>
                {
                    onSelect();
                }}

                onTap={onSelect}

                onDragEnd={(e) =>
                {
                    onChange({
                        ...shapeProps,
                        x : e.target.x(),
                        y : e.target.y(),
                        edited : true
                    })
                }}

                onTransformEnd={(e) =>
                {
                    onChange({
                        ...shapeProps,
                        x : e.target.x(),
                        y : e.target.y(),
                        rotation : e.target.rotation(),
                        edited : true
                    });
                }}

                onMouseEnter={(e) =>
                {
                    e.target.getStage().container().style.cursor = transform ? 'move' : 'pointer';
                    e.target.fill('#09a2fb');
                }}

                onMouseLeave={(e) =>
                {
                    e.target.getStage().container().style.cursor = 'grab';
                    if(isSelected)
                    {
                        e.target.fill('#09a2fb');
                    }
                    else if(!isSelected && booked)
                    {
                        e.target.fill('#7780e4');
                    }
                    else if(!isSelected && !booked)
                    {
                        e.target.fill('#e8e8e8');
                    }
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
}

export default MeetingRoom