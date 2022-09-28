import { Path, Transformer } from 'react-konva';
import { useRef, useEffect, useState, Fragment } from 'react';

const Desk = ({ shapeProps, isSelected, onSelect, onChange, draggable, transform, user}) =>
{
    const shapeRef = useRef(null);
    const transformRef = useRef(null);
    const [booked, setBooked] = useState(shapeProps.booked);
    const [userBooked, setUser] = useState(user);

    useEffect(() =>
    {
        if(isSelected && transform)
        {
            transformRef.current.nodes([shapeRef.current]);
            transformRef.current.getLayer().batchDraw();
        }

        if(!isSelected && booked)
        {
            if(userBooked)
            {
                shapeRef.current.fill('#000000');
            }
            else
            {
                shapeRef.current.fill('#e8e8e8');
            }
        }
        else if(!isSelected && !booked)
        {
            shapeRef.current.fill('#374146');
        }
    }, [isSelected, transform, booked, userBooked]);

    useEffect(() =>
    {
        if(booked)
        {
            if(userBooked)
            {
                shapeRef.current.fill('#000000');
            }
            else
            {
                shapeRef.current.fill('#e8e8e8');
            }
        }
        else
        {
            shapeRef.current.fill('#374146');
        }
    },[booked, userBooked])

    useEffect(() =>
    {
        setBooked(shapeProps.booked);
    },[shapeProps.booked]);

    useEffect(() =>
    {
        setUser(user);
    },[user]);

    return (
        <Fragment> 
            <Path 
                {...shapeProps}
                ref={shapeRef}

                data='h 200 a 20 20 0 0 1 20 20 v 80 a 20 20 0 0 1 -20 20 h -200 a 20 20 0 0 1 -20 -20 v -80 a 20 20 0 0 1 20 -20 Z M 50 -10 h 10 v -10 h 80 v 10 h 10 v -20 a 10 10 0 0 0 -100 0 Z'

                fill='#374146'

                scaleX={0.3}
                scaleY={0.3}

                draggable={draggable}

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
                        if(userBooked)
                        {
                            e.target.fill('#000000');
                        }
                        else
                        {
                            e.target.fill('#e8e8e8');
                        }
                        
                    }
                    else if(!isSelected && !booked)
                    {
                        e.target.fill('#374146');
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

export default Desk