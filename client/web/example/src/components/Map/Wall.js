import { Path, Rect, Transformer } from 'react-konva';
import { useRef, useEffect, useState, Fragment } from 'react';

const Wall = ({ shapeProps, isSelected, onSelect, onChange, draggable, transform}) =>
{
    const shapeRef = useRef(null);
    const transformRef = useRef(null);

    useEffect(() =>
    {
        if(isSelected && transform)
        {
            transformRef.current.nodes([shapeRef.current]);
            transformRef.current.getLayer().batchDraw();
        }
        else if(!isSelected)
        {
            shapeRef.current.fill('#374146');
        }
    }, [isSelected, transform]);


    return (
        <Fragment> 
            <Rect 
                {...shapeProps}
                ref={shapeRef}

                //data='h 200 a 20 20 0 0 1 20 20 v 80 a 20 20 0 0 1 -20 20 h -200 a 20 20 0 0 1 -20 -20 v -80 a 20 20 0 0 1 20 -20 Z M 50 -10 h 10 v -10 h 80 v 10 h 10 v -20 a 10 10 0 0 0 -100 0 Z'

                fill='#374146'

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
                    const scaleX = e.target.scaleX();
                    const scaleY = e.target.scaleY();

                    e.target.scaleX(1);
                    e.target.scaleY(1);

                    onChange({
                        ...shapeProps,
                        x : e.target.x(),
                        y : e.target.y(),
                        width: e.target.width() * scaleX,
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
                    else
                    {
                        e.target.fill('#374146');
                    }
                }}
            />
            
            {isSelected && (
                <Transformer 
                    ref = {transformRef}
                    rotationSnaps = {[0, 90, 180, 270]}
                    resizeEnabled = {true}
                    centeredScaling = {false}
                />
            )}
        </Fragment>
    );
}

export default Wall