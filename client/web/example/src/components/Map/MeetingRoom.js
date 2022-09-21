import useImage from 'use-image';
import meetingroom_img from '../../img/meetingroom_img.svg';
import { Image } from 'react-konva'
import { useRef, useEffect, Fragment } from 'react'
import { Transformer } from 'react-konva'

const MeetingRoom = ({ shapeProps, isSelected, onSelect, onChange, draggable}) =>
{
    const shapeRef = useRef(null);
    const transformRef = useRef(null);
    //const [image] = useImage(meetingroom_grey);
    const [image] = useImage(meetingroom_img);

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
            <Image
                image = {image}
                offsetX = {shapeProps.width / 2.0}
                offsetY = {shapeProps.height / 2.0}
                {...shapeProps}
                ref={shapeRef}
                draggable = {draggable}

                onClick={onSelect}
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

    /*return (
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
    );*/
}

export default MeetingRoom