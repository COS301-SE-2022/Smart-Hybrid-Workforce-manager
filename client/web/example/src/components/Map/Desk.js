import useImage from 'use-image';
import desk_grey from '../../img/desk_light.svg';
import { Image, Path, Transformer } from 'react-konva';
import { useRef, useEffect, Fragment } from 'react';

const Desk = ({ shapeProps, isSelected, onSelect, onChange, draggable, transform}) =>
{
    const shapeRef = useRef(null);
    const transformRef = useRef(null);
    const [image] = useImage(desk_grey);

    useEffect(() =>
    {
        if(isSelected && transform)
        {
            transformRef.current.nodes([shapeRef.current]);
            transformRef.current.getLayer().batchDraw();
        }
    }, [isSelected, transform]);

    return (
        <Fragment> 
            {/*<Image
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
            />*/}

            <Path 
                {...shapeProps}
                ref={shapeRef}

                data='h 200 a 20 20 0 0 1 20 20 v 80 a 20 20 0 0 1 -20 20 h -200 a 20 20 0 0 1 -20 -20 v -80 a 20 20 0 0 1 20 -20 Z M 50 -10 h 10 v -10 h 80 v 10 h 10 v -20 a 10 10 0 0 0 -100 0 Z'

                fill='#374146'

                scaleX={0.3}
                scaleY={0.3}

                draggable={draggable}

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