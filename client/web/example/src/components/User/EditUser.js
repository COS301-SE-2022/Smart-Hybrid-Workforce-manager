import Button from 'react-bootstrap/Button';
import styles from './user.module.css';
import { useEffect, useState } from 'react';

const EditUser = ({userName, userPicture, userRoles, allRoles}) =>
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

    const [roles, setRoles] = useState(allRoles);

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

    useEffect(() =>
    {
        setRoles(allRoles);
    }, [allRoles]);

    return (
        <div className={styles.editUserContainer}>
            <div className={styles.headerContainer}>
                <div className={styles.pictureContainer}>
                    <img className={styles.picture} src={picture} alt='User'></img>
                </div>

                <div className={styles.userName}>{name}</div>
            </div>

            <div className={styles.rolesContainer}>
                {roles.map((role) =>
                {
                    console.log(roles);
                    return (
                        <div key={role}>
                            <input type='checkbox' id={role} name={role} value={role} checked={activeRoles && activeRoles.includes(role) ? true : false} onChange={EditActiveRoles.bind(this, role)}></input>
                            <label className={styles.roleLabel} htmlFor={role}>{role}</label>
                        </div>
                    );
                })}
            </div>
        </div>
    );

}

export {EditUser as EditUserPanel}