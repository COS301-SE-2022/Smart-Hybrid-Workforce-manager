import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import styles from './resources.module.css';
import { useContext, useEffect, useState } from 'react';
import { UserContext } from '../../App';

const EditBuilding = ({id, edited}) =>
{
    const [name, setName] = useState('');
    const [location, setLocation] = useState('');

    const {userData} = useContext(UserContext);

    const EditBuildingSubmit = async () =>
    {
        fetch("http://deskflow.co.za:8080/api/resource/building/create", 
        {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                id: id,
                name: name,
                location: location,
                dimension: ''
            }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}`
            }
        }).then((res) =>
        {
            alert("Building Successfully Updated!");
            edited(true);
        });
    }

    useEffect(() =>
    {
        if(id)
        {
            fetch("http://deskflow.co.za:8080/api/resource/building/information", 
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
                setName(data[0].name);
                setLocation(data[0].location);
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
                <div className={styles.formLabel}>Location</div>
                <input className={styles.formInput} type='text' placeholder='Location' value={location} onChange={(e) => setLocation(e.target.value)}></input>
            </Form.Group>

            <Button className={styles.submit} onClick={EditBuildingSubmit}>Update</Button>
        </div>
    );

}

export {EditBuilding as EditBuildingForm}