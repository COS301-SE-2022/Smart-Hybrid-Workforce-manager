import Button from 'react-bootstrap/Button';
import styles from './user.module.css';
import { useEffect, useState } from 'react';

const EditUser = ({userName, userPicture, userRoles}) =>
{
    const [name, setName] = useState(userName);
    const [picture, setPicture] = useState(userPicture);
    const [activeRoles, setActiveRoles] = useState(userRoles);

    const rolesInit = 
    {
        ['role1']:
        {
            name: 'Developer',
            color: '#09a2fb'
        },

        ['role2']:
        {
            name: 'Secretary',
            color: '#09a2fb'
        },

        ['role3']:
        {
            name: 'CEO',
            color: '#09a2fb'
        },

        ['role4']:
        {
            name: 'Engineer',
            color: '#09a2fb'
        },
    }

    const [roles, setRoles] = useState(rolesInit);

    const EditActiveRoles = (role) =>
    {
        if(activeRoles.includes(role))
        {
            setActiveRoles(activeRoles.filter((curr, _) =>
                curr !== role
            ));
        }
        else
        {
            setActiveRoles([...activeRoles, role]);
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

    useEffect(() =>
    {
        setActiveRoles(userRoles);
    }, [userRoles]);

    return (
        <div className={styles.editUserContainer}>
            <div className={styles.headerContainer}>
                <div className={styles.pictureContainer}>
                    <img className={styles.picture} src={picture} alt='User'></img>
                </div>

                <div className={styles.userName}>{name}</div>
            </div>

            <div className={styles.rolesContainer}>
                {Object.entries(roles).map(([id, role]) =>
                {
                    return (
                        <div key={id}>
                            <input type='checkbox' id={id} name={id} value={role.name} checked={activeRoles && activeRoles.includes(role.name) ? true : false} onChange={EditActiveRoles.bind(this, role.name)}></input>
                            <label className={styles.roleLabel} htmlFor={role.name}>{role.name}</label>
                        </div>
                    );
                })}
            </div>
        </div>
    );

}

export {EditUser as EditUserPanel}