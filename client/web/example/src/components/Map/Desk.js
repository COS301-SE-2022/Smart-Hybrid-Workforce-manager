import useImage from 'use-image';
import desk_grey from '../../img/desk_grey.svg';
import { Image } from 'react-konva'
import { useRef, useEffect, Fragment } from 'react'
import { Transformer } from 'react-konva'

const Desk = ({ shapeProps, isSelected, onSelect, onChange}) =>
{
    const shapeRef = useRef(null);
    const transformRef = useRef(null);
    const [image] = useImage(desk_grey);

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
                width = {70}
                height = {50}
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