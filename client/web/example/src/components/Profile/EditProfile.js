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
    const [building, setBuilding] = useState('');
    const [buildings, SetBuildings] = useState([]);

    const buildingRef = useRef(null);
    const pictureInputRef = useRef(null);

    const homeRef = useRef(null);
    const parkingSRef = useRef(null);
    const parkingDRef = useRef(null);
    const parkingNRef = useRef(null);

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
                                fetch("http://localhost:8080/api/user/update", 
                                {
                                    method: "POST",
                                    mode: "cors",
                                    body: JSON.stringify({
                                        id: id,
                                        picture: url,
                                        work_from_home: workFromHome,
                                        parking: parking,
                                        office_days: officeDays,
                                        preferred_start_time: '0000-01-01T' + startTime + ':00.000Z',
                                        preferred_end_time: '0000-01-01T' + endTime + ':00.000Z',
                                        building_id: building
                                    }),
                                    headers:{
                                        'Content-Type': 'application/json',
                                        'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                                    }
                                }).then((res) =>
                                {
                                    if(res.status === 200)
                                    {
                                        alert("Profile Successfully Updated!");
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
            fetch("http://localhost:8080/api/user/update", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    id: id,
                    picture: picture,
                    work_from_home: workFromHome,
                    parking: parking,
                    office_days: officeDays,
                    preferred_start_time: '0000-01-01T' + startTime + ':00.000Z',
                    preferred_end_time: '0000-01-01T' + endTime + ':00.000Z',
                    building_id: building
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}`
                }
            }).then((res) =>
            {
                if(res.status === 200)
                {
                    alert("Profile Successfully Updated!");
                    edited(true);
                }
            });
        }
    }

    useEffect(() =>
    {
        if(user && Object.entries(user).length > 0)
        {
            setID(user.id);
            setParking(user.parking);
            setWorkFromHome(user.work_from_home);
            setOfficeDays(user.office_days);
            setStartTime(user.preferred_start_time.substring(11,16));
            setEndTime(user.preferred_end_time.substring(11,16));
            setPicture(user.picture);
            setBuilding(user.building_id)

            console.log(user);
        }
    }, [user]);

    useEffect(() =>
    {
        if(parkingSRef && parkingDRef && parkingNRef)
        {
            if(parking === 'STANDARD')
            {
                parkingSRef.current.checked = true;
            }
            else if(parking === 'DISABLED')
            {
                parkingDRef.current.checked = true;
            }
            else if(parking === 'NONE')
            {
                parkingNRef.current.checked = true;
            }
        }
    }, [parking]);

    useEffect(() =>
    {
        if(homeRef.current)
        {
            homeRef.current.checked = workFromHome;
        }
    }, [workFromHome]);

    useEffect(() => {
        fetch("http://localhost:8080/api/resource/building/information",
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                }),
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                }
            }).then((res) => res.json()).then(data => {
                SetBuildings(data);
            });
    }, [userData.token, userData.user_identifier]);

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
                <div className={styles.buildingSelectorContainer}>
                <div className={styles.formLabel}>Building</div>
                <select ref={buildingRef} className={styles.resourceSelector} name='building' onChange={(e) => setBuilding(e.target.value)}>
                    <option value='' disabled selected id='BuildingDefault'>--Select the building--</option>
                        {buildings.length > 0 && (
                            buildings.map(building => (
                                <option key={building.id} value={building.id}>{building.name + ' (' + building.location + ')'}</option>
                            ))
                        )}
                </select>
            </div>
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Work from home</div>
                <input ref={homeRef} type='checkbox' onChange={(e) => setWorkFromHome(e.target.checked)}></input>
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Parking type</div>

                <form>
                    <div className={styles.radioButtonContainer}>
                        <input ref={parkingSRef} type='radio' name='parking' value={'STANDARD'} onChange={(e) => setParking(e.target.value)}></input>
                        <label className={styles.radioLabel}>Standard</label>
                    </div>

                    <div className={styles.radioButtonContainer}>
                        <input ref={parkingDRef} type='radio' name='parking' value={'DISABLED'} onChange={(e) => setParking(e.target.value)}></input>
                        <label className={styles.radioLabel}>Disabled</label>
                    </div>

                    <div className={styles.radioButtonContainer}>
                        <input ref={parkingNRef} type='radio' name='parking' value={'NONE'} onChange={(e) => setParking(e.target.value)}></input>
                        <label className={styles.radioLabel}>None</label>
                    </div>
                </form>
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Office days per week</div>
                <input className={styles.formInput} type='number' min='1' placeholder="Office days" value={officeDays} onChange={(e) => setOfficeDays(e.target.value)}></input>
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Preferred start time</div>
                <input className={styles.formInput} type='time' value={startTime} onChange={(e) => setStartTime(e.target.value)}></input>
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Preferred end time</div>
                <input className={styles.formInput} type='time' value={endTime} onChange={(e) => setEndTime(e.target.value)}></input>
            </Form.Group>

            <Button className={styles.submit} onClick={EditProfileSubmit}>Save</Button>
        </div>
    );

}

export default EditProfile