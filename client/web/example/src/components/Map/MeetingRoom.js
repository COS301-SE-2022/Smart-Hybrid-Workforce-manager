import useImage from 'use-image';
import meetingroom_grey from '../../img/meetingroom_grey.svg';
import { Image, Rect } from 'react-konva'
import { useRef, useEffect, Fragment, useState } from 'react'
import { Transformer } from 'react-konva'

const MeetingRoom = ({ shapeProps, isSelected, onSelect, onChange, stage}) =>
{
    const shapeRef = useRef(null);
    const imgRef = useRef(null);
    const transformRef = useRef(null);
    const [image] = useImage(meetingroom_grey);
    const [center, setCenter] = useState([(-stage.x() + stage.width() / 2.0) / stage.scaleX(), (-stage.y() + stage.height() / 2.0) / stage.scaleY()]);

    const calculateCenter = (x, offX, y, offY, width, height, angle) =>
    {
        angle = angle * Math.PI / 180;
        const cX = x + ((width / 2) * Math.cos(-angle)) + ((height / 2) * Math.sin(-angle)) - (offX * Math.cos(-angle)) - (offY * Math.sin(-angle));
        const cY = y + ((width / 2) * Math.sin(angle)) + ((height / 2) * Math.cos(angle)) - (offY * Math.cos(angle)) - (offX * Math.sin(angle));

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

    useEffect(() =>
    {
        calculateCenter(shapeRef.current.x(), shapeRef.current.offsetX(), shapeRef.current.y(), shapeRef.current.offsetY(), shapeRef.current.width(), shapeRef.current.height(), shapeRef.current.getAbsoluteRotation());
    }, []);

    return (
        <Fragment> 
            <Rect 
                cornerRadius = {10}
                fill = {"#bfcbd6"}
                {...shapeProps}
                ref={shapeRef}
                offsetX = {shapeProps.width / 2.0}
                offsetY = {shapeProps.height / 2.0}
                draggable

                onClick={onSelect}
                onTap={onSelect}

                onDragMove={(e) =>
                {
                    onChange({
                        ...shapeProps,
                        x : e.target.x(),
                        y : e.target.y(),
                        edited : true
                    })

                    calculateCenter(e.target.x(), e.target.offsetX(), e.target.y(), e.target.offsetY(), e.target.width(), e.target.height(), e.target.getAbsoluteRotation());
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
                        rotation : e.target.rotation(),
                        edited : true
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