import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import styles from './team.module.css';
import { useEffect, useState } from 'react';

const AddTeam = () =>
{
    const [color, setColor] = useState('#ffffff');
    const [name, setName] = useState('');
    const [priority, setPriority] = useState('');
    const [picture, setPicture] = useState('');

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

    const AddTeamSubmit = () =>
    {

    }


    return (
        <Form className={styles.form} onSubmit={AddTeamSubmit}>
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


            <Button className={styles.submit} type='submit'>Add Team</Button>
        </Form>
    );

}

export {AddTeam as AddTeamForm}