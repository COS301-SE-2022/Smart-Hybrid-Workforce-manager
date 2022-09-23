import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import styles from './team.module.css';
import { useContext, useEffect, useState } from 'react';
import { UserContext } from '../../App';
import { useNavigate } from 'react-router-dom';

const EditTeam = ({teamName, teamColor, teamPriority, teamPicture, teamID}) =>
{
    const [id, setID] = useState(teamID);
    const [color, setColor] = useState(teamColor);
    const [name, setName] = useState(teamName);
    const [priority, setPriority] = useState(teamPriority);
    const [picture, setPicture] = useState('');

    const {userData} = useContext(UserContext);
    const navigate = useNavigate();

    const CheckPriority = (value) =>
    {
        if(value < 0)
        {
            setPriority(0);
        }
        else
        {
            setPriority(value);
        }
    }

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

    const EditTeamSubmit = async () =>
    {
        try
        {
            let res = await fetch("http://localhost:8080/api/team/create", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    id: id,
                    name: name,
                    priority: priority
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                }
            });

            if(res.status === 200)
            {
                alert("Team Successfully Updated!");
                navigate("/admin");
            }
        }
        catch(err)
        {

        }
    }

    useEffect(() =>
    {
        setID(teamID);
    }, [teamID]);

    useEffect(() =>
    {
        setName(teamName);
    }, [teamName]);

    useEffect(() =>
    {
        setColor(teamColor);
    }, [teamColor]);

    useEffect(() =>
    {
        setPriority(teamPriority);
    }, [teamPriority]);

    useEffect(() =>
    {
        setPicture('');
    }, [teamPicture]);


    return (
        <Form className={styles.form} onSubmit={EditTeamSubmit}>
            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <Form.Label className={styles.formLabel}>Team name</Form.Label>
                <Form.Control className={styles.formInput} type='text' placeholder="Team name" value={name} onChange={(e) => setName(e.target.value)} />
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <Form.Label className={styles.formLabel}>Team color</Form.Label>
                <Form.Control className={styles.colorPicker} type='color' value={color} onChange={(e) => setColor(e.target.value)} />
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <Form.Label className={styles.formLabel}>Team priority</Form.Label>
                <Form.Control className={styles.formInput} type='number' placeholder="Team priority" value={priority} onChange={(e) => CheckPriority(e.target.value)} />
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <Form.Label className={styles.formLabel}>Team picture</Form.Label>
                <Form.Control className={styles.fileUpload} type='file' value={picture} onChange={(e) => CheckPicture(e.target.value)} />
            </Form.Group>


            <Button className={styles.submit} type='submit'>Edit Team</Button>
        </Form>
    );

}

export {EditTeam as EditTeamForm}