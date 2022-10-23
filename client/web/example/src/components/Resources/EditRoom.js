import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import styles from './resources.module.css';
import { useContext, useEffect, useState } from 'react';
import { UserContext } from '../../App';

const EditRoom = ({id, edited}) =>
{
    const [buildingID, setBuildingID] = useState('');
    const [name, setName] = useState('');
    const [floor, setFloor] = useState('');
    const [x, setX] = useState(0);
    const [y, setY] = useState(0);
    
    const {userData} = useContext(UserContext);

    const EditRoomSubmit = async () =>
    {
        fetch("http://deskflow.co.za:8080/api/resource/room/create", 
        {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                id: id,
                building_id: buildingID,
                name: name,
                xcoord: x,
                ycoord: y,
                zcoord: parseInt(floor),
                dimension: ''
            }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}`
            }
        }).then((res) =>
        {
            alert("Room Successfully Updated!");
            edited(true);
        });
    }

    useEffect(() =>
    {
        if(id)
        {
            fetch("http://deskflow.co.za:8080/api/resource/room/information", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    id: id
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}`
                }
            }).then((res) => res.json()).then(data =>
            {
                setBuildingID(data[0].building_id);
                setName(data[0].name);
                setFloor(data[0].zcoord);
                setX(parseFloat(data[0].xcoord));
                setY(parseFloat(data[0].ycoord));
            });
        }

    }, [id, userData.token]);

    return (
        <div className={styles.form}>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Name</div>
                <input className={styles.formInput} type='text' placeholder="Name" value={name} onChange={(e) => setName(e.target.value)}></input>
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Floor</div>
                <input className={styles.formInput} type='text' placeholder='Floor' value={floor} onChange={(e) => setFloor(e.target.value)}></input>
            </Form.Group>

            <Button className={styles.submit} onClick={EditRoomSubmit}>Update</Button>
        </div>
    );

}

export {EditRoom as EditRoomForm}