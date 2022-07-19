import useImage from 'use-image';
import meetingroom_grey from '../../img/meetingroom_grey.svg';
import { Image, Rect } from 'react-konva'
import { useRef, useEffect, Fragment } from 'react'
import { Transformer } from 'react-konva'

const MeetingRoom = ({ shapeProps, isSelected, onSelect, onChange}) =>
{
    const shapeRef = useRef(null);
    const imgRef = useRef(null);
    const transformRef = useRef(null);
    const [image] = useImage(meetingroom_grey);

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
                x = {shapeProps.x + shapeProps.width/2.0 - 65}
                y = {shapeProps.y + shapeProps.height/2.0 - 30}
                width = {130}
                height = {60}
                ref={imgRef}
                rotation = {shapeProps.rotation}
                draggable

                onClick={onSelect}
                onTap={onSelect}

                onDragMove={(e) =>
                {
                    onChange({
                        ...shapeProps,
                        x : e.target.x() + 65 - shapeProps.width/2.0,
                        y : e.target.y() + 30 - shapeProps.height/2.0
                    })
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
                    resizeEnabled = {true}
                    centeredScaling = {true}
                />
            )}
        </Fragment>
    );
}

export default MeetingRoom