import React from 'react'
import { FaLongArrowAltRight } from 'react-icons/fa'
import { useNavigate } from 'react-router-dom'

const BookingCard = ({name, description, path, image}) =>
{
    let navigate = useNavigate();
    const route = () =>
    {
        navigate(path);
    }

    return (
        <div>
            <div className="card" onClick={route}>
                <div className="card-image" 
                    style={{
                        gridArea : 'image',
                        background : 'linear-gradient(#fff0 0%, #fff0 70%, #1d1d1d 100%), url(' + image + ')',
                        'backgroundSize': 'cover',
                        'borderTopLeftRadius': '4vh',
                        'borderTopRightRadius': '4vh'
                        }}>
                </div>
                
                <div className="card-text">
                    <h2>{name}</h2>
                    <p>{description}</p>
                </div>
                <div className="card-arrow">
                    <FaLongArrowAltRight size={50}/>
                </div>
            </div>
        </div>
    )
}

export default BookingCard