import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import styles from './team.module.css';
import { useContext, useEffect, useRef, useState } from 'react';
import { UserContext } from '../../App';

const AddTeam = ({makeDefault, edited}) =>
{
    const [name, setName] = useState('');
    const [color, setColor] = useState('#000000');
    const [capacity, setCapacity] = useState('2');
    const [priority, setPriority] = useState(0);

    const priority0Ref = useRef(null);
    const priority1Ref = useRef(null);
    const priority2Ref = useRef(null);

    const {userData} = useContext(UserContext);

    const AddTeamSubmit = async () =>
    {
        fetch("http://deskflow.co.za:8080/api/team/create", 
        {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                name: name,
                color: color,
                capacity: parseInt(capacity),
                picture: 'https://firebasestorage.googleapis.com/v0/b/arche-6bd39.appspot.com/o/teams%2FTeamDefault.png?alt=media&token=66cbabd9-a01f-47b9-9861-89b7aa523697',
                priority: priority,
            }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}`
            }
        }).then((res) =>
        {
            if(res.status === 200)
            {
                alert("Team Successfully Created!");
                edited(true);
            }
        });
    }

    useEffect(() =>
    {
        setName('');
        setColor('#000000');
        setCapacity('2');
        setPriority(0);
    }, [makeDefault]);

    useEffect(() =>
    {
        if(priority0Ref && priority1Ref && priority2Ref)
        {
            if(priority === 0)
            {
                priority0Ref.current.checked = true;
            }
            else if(priority === 1)
            {
                priority1Ref.current.checked = true;
            }
            else if(priority === 2)
            {
                priority2Ref.current.checked = true;
            }
        }
    }, [priority]);

    return (
        <div className={styles.form}>
            <div className={styles.pictureEditContainer}>
                <div className={styles.pictureContainer}>
                    <img className={styles.picture} src={'https://firebasestorage.googleapis.com/v0/b/arche-6bd39.appspot.com/o/teams%2FTeamDefault.png?alt=media&token=66cbabd9-a01f-47b9-9861-89b7aa523697'} alt='Team'></img>
                </div>
            </div>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Name</div>
                <Form.Control className={styles.formInput} type='text' placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} />
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Color</div>
                <Form.Control className={styles.colorPicker} type='color' value={color} onChange={(e) => setColor(e.target.value)} />
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Capacity</div>
                <Form.Control className={styles.formInput} type='number' min='2' placeholder="Capacity" value={capacity} onChange={(e) => setCapacity(e.target.value)} />
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Priority</div>

                <form>
                    <input ref={priority0Ref} type='radio' name='priority' value={0} onChange={(e) => setPriority(e.target.value)}></input>
                    <label className={styles.radioLabel}>0</label>

                    <input ref={priority1Ref} type='radio' name='priority' value={1} onChange={(e) => setPriority(e.target.value)}></input>
                    <label className={styles.radioLabel}>1</label>

                    <input ref={priority2Ref} type='radio' name='priority' value={2} onChange={(e) => setPriority(e.target.value)}></input>
                    <label className={styles.radioLabel}>2</label>
                </form>
            </Form.Group>

            <Button className={styles.submit} onClick={AddTeamSubmit}>Create</Button>
        </div>
    );

}

export {AddTeam as AddTeamForm}