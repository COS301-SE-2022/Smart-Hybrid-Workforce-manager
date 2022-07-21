import useImage from 'use-image';
import meetingroom_grey from '../../img/meetingroom_grey.svg';
import { Image, Rect } from 'react-konva'
import { useRef, useEffect, Fragment, useState } from 'react'
import { Transformer } from 'react-konva'

const MeetingRoom = ({ shapeProps, isSelected, onSelect, onChange}) =>
{
    const shapeRef = useRef(null);
    const imgRef = useRef(null);
    const transformRef = useRef(null);
    const [image] = useImage(meetingroom_grey);
    const [center, setCenter] = useState([100,100]);

    const calculateCenter = (x, y, width, height, angle) =>
    {
        angle = angle * Math.PI / 180;
        const cX = x + ((width / 2) * Math.cos(-angle)) + ((height / 2) * Math.sin(-angle));
        const cY = y + ((width / 2) * Math.sin(angle)) + ((height / 2) * Math.cos(angle));

        setCenter([cX, cY]);
    }

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
                cornerRadius = {10}
                fill = {"#bfcbd6"}
                {...shapeProps}
                ref={shapeRef}
                draggable

                onClick={onSelect}
                onTap={onSelect}

                onDragMove={(e) =>
                {
                    onChange({
                        ...shapeProps,
                        x : e.target.x(),
                        y : e.target.y()
                    })

                    calculateCenter(e.target.x(), e.target.y(), e.target.width(), e.target.height(), e.target.getAbsoluteRotation());
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
                        width : e.target.width() * scaleX,
                        height : e.target.height() * scaleY,
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

            <Image
                image = {image}
                rotation = {shapeProps.rotation}
                x = {center[0]}
                y = {center[1]}
                width = {130}
                height = {60}
                offsetX = {65}
                offsetY = {30}
                ref={imgRef}

                onClick={onSelect}
                onTap={onSelect}

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
                    resizeEnabled = {true}
                    centeredScaling = {true}
                />
            )}
        </Fragment>
    );
}

export default MeetingRoom