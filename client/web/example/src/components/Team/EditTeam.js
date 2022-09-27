import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import styles from './team.module.css';
import { useContext, useEffect, useRef, useState } from 'react';
import { UserContext } from '../../App';
import { useNavigate } from 'react-router-dom';
import { storage } from '../../firebase';
import { ref, uploadBytes, listAll, getDownloadURL } from 'firebase/storage';

const EditTeam = ({team, edited}) =>
{
    const [id, setID] = useState('');
    const [color, setColor] = useState('#000000');
    const [name, setName] = useState('');
    const [priority, setPriority] = useState(0);
    const [capacity, setCapacity] = useState('2');
    const [picture, setPicture] = useState('');
    const [pictureUpload, setPictureUpload] = useState(null);
    const [members, setMembers] = useState([{name: '', id: ''}]);
    const [lead, setLead] = useState(null);

    const priority0Ref = useRef(null);
    const priority1Ref = useRef(null);
    const priority2Ref = useRef(null);

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

    const EditTeamSubmit = async () =>
    {
        if(pictureUpload !== null)
        {            
            const lastPeriod = pictureUpload.name.lastIndexOf(".");
            const newName = id + pictureUpload.name.substring(lastPeriod);
            const pictureRef = ref(storage, `teams/${newName}`);

            uploadBytes(pictureRef, pictureUpload).then(() =>
            {
                const pictureListRef = ref(storage, 'teams/')
                listAll(pictureListRef).then((response) =>
                {
                    response.items.forEach((picture) =>
                    {
                        getDownloadURL(picture).then((url) =>
                        {
                            if(url.includes(newName))
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
        if(team)
        {
            setID(team.id);
            setName(team.name);
            setColor(team.color);
            setPriority(team.priority);
            setPicture(team.picture);
            setMembers(team.users);
            setLead(team.lead);

            if(team.capacity)
            {
                setCapacity(team.capacity.toString());
            }
            else
            {
                setCapacity('2');
            }
        }
    }, [team]);

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
                    <img className={styles.picture} src={picture} alt='Team'></img>
                </div>

                <Form.Control ref={pictureInputRef} style={{display: 'none'}} type='file' accept='image/png, image/jpeg, image/jpg' onChange={ChangePicture.bind(this)} />
                <div className={styles.fileUpload} onClick={ClickPictureInput}>Change picture</div>
            </div>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Name</div>
                <input className={styles.formInput} type='text' placeholder="Team name" value={name} onChange={(e) => setName(e.target.value)}></input>
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Color</div>
                <Form.Control className={styles.colorPicker} type='color' value={color} onChange={(e) => setColor(e.target.value)} />
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

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Lead</div>
                <select className={styles.lead} name='lead' onChange={(e) => setLead(e.target.value)}>
                    {members && (
                        members.map(member => (
                            <option key={member.id} value={member.id}>{member.name}</option>
                        ))
                    )}
                </select>
            </Form.Group>


            <Button className={styles.submit} onClick={EditTeamSubmit}>Save</Button>
        </div>
    );

}

export {EditTeam as EditTeamForm}