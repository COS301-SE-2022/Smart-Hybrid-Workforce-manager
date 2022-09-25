import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import styles from '../Team/team.module.css';
import { useContext, useEffect, useRef, useState } from 'react';
import { UserContext } from '../../App';
import { storage } from '../../firebase';
import { ref, uploadBytes, listAll, getDownloadURL } from 'firebase/storage';

const EditProfile = ({user, edited}) =>
{
    const [id, setID] = useState('');
    const [parking, setParking] = useState('');
    const [workFromHome, setWorkFromHome] = useState(false);
    const [officeDays, setOfficeDays] = useState(1);
    const [startTime, setStartTime] = useState('');
    const [endTime, setEndTime] = useState('');
    const [picture, setPicture] = useState('');
    const [pictureUpload, setPictureUpload] = useState(null);

    const pictureInputRef = useRef(null);

    const {userData} = useContext(UserContext);

    const ClickPictureInput = () =>
    {
        if(pictureInputRef)
        {
            pictureInputRef.current.click();
        }
    }

    const ChangePicture = (e) =>
    {
        if(e.target.files[0])
        {
            setPictureUpload(e.target.files[0]);
            
            var reader = new FileReader();
            reader.readAsDataURL(e.target.files[0]);
            reader.onloadend = function()
            {
                setPicture(reader.result);
            }
        }
    }

    const EditProfileSubmit = async () =>
    {
        if(pictureUpload !== null)
        {            
            const lastPeriod = pictureUpload.name.lastIndexOf(".");
            const newName = id + pictureUpload.name.substring(lastPeriod);
            const pictureRef = ref(storage, `users/${newName}`);

            uploadBytes(pictureRef, pictureUpload).then(() =>
            {
                const pictureListRef = ref(storage, 'users/')
                listAll(pictureListRef).then((response) =>
                {
                    response.items.forEach((picture) =>
                    {
                        getDownloadURL(picture).then((url) =>
                        {
                            if(url.includes(newName))
                            {
                                fetch("http://localhost:8080/api/user/create", 
                                {
                                    method: "POST",
                                    mode: "cors",
                                    body: JSON.stringify({
                                        id: id,
                                        name: name,
                                        color: color,
                                        capacity: parseInt(capacity),
                                        picture: url,
                                        priority: priority,
                                        team_lead_id: lead
                                    }),
                                    headers:{
                                        'Content-Type': 'application/json',
                                        'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                                    }
                                }).then((res) =>
                                {
                                    if(res.status === 200)
                                    {
                                        alert("Team Successfully Updated!");
                                        edited(true);
                                    }
                                });
                            }
                        })
                    });
                });
            });
        }
        else
        {
            fetch("http://localhost:8080/api/team/create", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    id: id,
                    name: name,
                    color: color,
                    capacity: parseInt(capacity),
                    picture: picture,
                    priority: priority,
                    team_lead_id: lead
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                }
            }).then((res) =>
            {
                if(res.status === 200)
                {
                    alert("Team Successfully Updated!");
                    edited(true);
                }
            });
        }
    }

    useEffect(() =>
    {
        if(user)
        {
            setID(user.id);
            setParking(user.parking);
            setWorkFromHome(user.work_from_home);
            setOfficeDays(user.office_days);
            setStartTime(user.preferred_start_time.substring(11,16));
            setEndTime(user.preferred_end_time.substring(11,16));
            setPicture(user.picture);
        }
    }, [user]);

    return (
        <div className={styles.form}>
            <div className={styles.pictureEditContainer}>
                <div className={styles.pictureContainer}>
                    <img className={styles.picture} src={picture} alt='user'></img>
                </div>

                <Form.Control ref={pictureInputRef} style={{display: 'none'}} type='file' accept='image/png, image/jpeg, image/jpg' onChange={ChangePicture.bind(this)} />
                <div className={styles.fileUpload} onClick={ClickPictureInput}>Change picture</div>
            </div>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Name</div>
                <input className={styles.formInput} type='text' placeholder="Team name" value={id} onChange={(e) => setName(e.target.value)}></input>
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Color</div>
                <Form.Control className={styles.colorPicker} type='color' value={officeDays} onChange={(e) => setColor(e.target.value)} />
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Capacity</div>
                <input className={styles.formInput} type='number' min='2' placeholder="Capacity" value={capacity} onChange={(e) => setCapacity(e.target.value)}></input>
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


            <Button className={styles.submit} onClick={EditProfileSubmit}>Save</Button>
        </div>
    );

}

export default EditProfile