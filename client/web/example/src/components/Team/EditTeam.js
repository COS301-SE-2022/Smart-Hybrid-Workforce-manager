import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import styles from './team.module.css';
import { HuePicker } from 'react-color';
import { useState } from 'react';

const EditTeam = ({teamName, teamColor}) =>
{
    const [color, setColor] = useState(teamColor);

    const ChangeColor = (color) =>
    {
        setColor(color);
    }

    const EditTeamSubmit = () =>
    {

    }

    return (
        <Form className={styles.form} onSubmit={EditTeamSubmit}>
            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <Form.Label className={styles.formLabel}>Team name</Form.Label>
                <Form.Control className={styles.formInput} type='text' placeholder="Enter your team name" value={teamName} />
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.colorLabel}>
                    <Form.Label className={styles.formLabel}>Team color</Form.Label>
                    <div className={styles.colorSwatch} style={{backgroundColor: color.hex}}></div>
                </div>
                <HuePicker className={styles.colorPicker} color={color} onChange={ChangeColor} width='20vw' />
            </Form.Group>


            <Button className='button-submit' variant='primary' type='submit'>Create Booking</Button>
        </Form>
    );

}

export {EditTeam as EditTeamForm}