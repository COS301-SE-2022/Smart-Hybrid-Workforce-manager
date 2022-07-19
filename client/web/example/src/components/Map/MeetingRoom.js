import useImage from 'use-image';
//import meetingroom_grey from '../../img/meetingroom_grey.svg';
import { Image, Rect } from 'react-konva'
import { useRef, useEffect, Fragment } from 'react'
import { Transformer } from 'react-konva'

const MeetingRoom = ({ shapeProps, isSelected, onSelect, onChange}) =>
{
    const shapeRef = useRef(null);
    const transformRef = useRef(null);
    //const [image] = useImage(meetingroom_grey);

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
                        width : e.target.width(),
                        height : e.target.height(),
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

            {/*<Image
                image = {image}
                width = {60}
                height = {55}
                {...shapeProps}
                ref={shapeRef}
                draggable

                onClick={onSelect}
                onTap={onSelect}

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
            />*/}
            
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