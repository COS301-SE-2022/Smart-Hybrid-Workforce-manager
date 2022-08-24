import React, { useContext, useState, useRef, useEffect } from 'react'

const Circle = ({size, top, left, color}, ref) =>
{
    const circleStyle =
    {
        backgroundColor: color,
        left: left,
        top: top,
        height: size,
        width: size,
        position: 'absolute',
        borderRadius: '50%'
    };

    return (
        <div ref={ref} className='circle' style={circleStyle}></div>
    )
}

export default React.forwardRef(Circle)