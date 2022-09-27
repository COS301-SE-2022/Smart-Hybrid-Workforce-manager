import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import styles from './resources.module.css';
import { useContext, useEffect, useState } from 'react';
import { UserContext } from '../../App';

const AddRoom = ({makeDefault, edited, buildingID}) =>
{
    const [name, setName] = useState('');
    const [floor, setFloor] = useState('0');

    const {userData} = useContext(UserContext);

    const AddRoomSubmit = async () =>
    {
        fetch("http://localhost:8080/api/resource/room/create", 
        {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                id: null,
                building_id: buildingID,
                name: name,
                xcoord: 0,
                ycoord: 0,
                zcoord: parseInt(floor),
                dimension: ''
            }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}`
            }
        }).then((res) =>
        {
            if(res.status === 200)
            {
                alert("Room Successfully Created!");
                edited(true);
            }
        });
    }

    useEffect(() =>
    {
        setName('');
        setFloor('0');

    }, [makeDefault]);

    return (
        <div className={styles.form}>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Name</div>
                <input className={styles.formInput} type='text' placeholder="Name" value={name} onChange={(e) => setName(e.target.value)}></input>
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Floor</div>
                <input className={styles.formInput} type='number' placeholder='Floor' value={floor} onChange={(e) => setFloor(e.target.value)}></input>
            </Form.Group>

            <Button className={styles.submit} onClick={AddRoomSubmit}>Create</Button>
        </div>
    );

}

export {AddRoom as AddRoomForm}