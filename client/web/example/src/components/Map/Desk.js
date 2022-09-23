import useImage from 'use-image';
import desk_grey from '../../img/desk_light.svg';
import { Image } from 'react-konva'
import { useRef, useEffect, Fragment } from 'react'
import { Transformer } from 'react-konva'

const Desk = ({ shapeProps, isSelected, onSelect, onChange, draggable}) =>
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
                offsetX = {30}
                offsetY = {27.5}
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