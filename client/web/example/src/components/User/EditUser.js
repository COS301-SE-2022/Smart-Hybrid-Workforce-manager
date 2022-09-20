import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import styles from './user.module.css';
import { useEffect, useState } from 'react';

const EditUser = ({userName, userPicture}) =>
{
    const [name, setName] = useState(userName);
    const [picture, setPicture] = useState(userPicture);

    const CheckPicture = (value) =>
    {
        const extension = value.substring(value.length - 3, value.length);
        if(!(extension === 'png' || extension === 'jpg'))
        {
            setPicture('');
        }
        else
        {
            setPicture(value);
        }
    }

    const EditTeamSubmit = () =>
    {

    }

    useEffect(() =>
    {
        setName(userName);
    }, [userName]);

    useEffect(() =>
    {
        setPicture(userPicture);
    }, [userPicture]);


    return (
        <div className={styles.editUserContainer}>
            <div className={styles.headerContainer}>
                <div className={styles.pictureContainer}>
                    <img className={styles.picture} src={picture} alt='User'></img>
                </div>

                <div className={styles.userName}>{name}</div>
            </div>
        </div>
    );

}

export {EditUser as EditUserPanel}