import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import styles from './team.module.css';
import { useContext, useEffect, useRef, useState } from 'react';
import { UserContext } from '../../App';
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

    //Permission states
    const viewTeamRef = useRef(null);
    const [viewTeam, setViewTeam] = useState(false);
    const [viewTeamID, setViewTeamID] = useState('');

    const updateTeamRef = useRef(null);
    const [updateTeam, setUpdateTeam] = useState(false);
    const [updateTeamID, setUpdateTeamID] = useState('');

    const deleteTeamRef = useRef(null);
    const [deleteTeam, setDeleteTeam] = useState(false);
    const [deleteTeamID, setDeleteTeamID] = useState('');

    const viewTeamMemberRef = useRef(null);
    const [viewTeamMember, setViewTeamMember] = useState(false);
    const [viewTeamMemberID, setViewTeamMemberID] = useState('');

    const updateTeamMemberRef = useRef(null);
    const [updateTeamMember, setUpdateTeamMember] = useState(false);
    const [updateTeamMemberID, setUpdateTeamMemberID] = useState('');

    const deleteTeamMemberRef = useRef(null);
    const [deleteTeamMember, setDeleteTeamMember] = useState(false);
    const [deleteTeamMemberID, setDeleteTeamMemberID] = useState('');

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
                                fetch("http://deskflow.co.za:8080/api/team/create", 
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
                                        //Team
                                        if(viewTeam && viewTeamID === '')
                                        {
                                            AddPermission(id, 'VIEW', 'TEAM', 'IDENTIFIER');
                                        }
                                        else if(!viewTeam && viewTeamID !== '')
                                        {
                                            RemovePermission(viewTeamID);
                                        }

                                        if(updateTeam && updateTeamID === '')
                                        {
                                            AddPermission(id, 'CREATE', 'TEAM', 'IDENTIFIER');
                                        }
                                        else if(!updateTeam && updateTeamID !== '')
                                        {
                                            RemovePermission(updateTeamID);
                                        }

                                        if(deleteTeam && deleteTeamID === '')
                                        {
                                            AddPermission(id, 'DELETE', 'TEAM', 'IDENTIFIER');
                                        }
                                        else if(!deleteTeam && deleteTeamID !== '')
                                        {
                                            RemovePermission(deleteTeamID);
                                        }

                                        //Team member
                                        if(viewTeamMember && viewTeamMemberID === '')
                                        {
                                            AddPermission(id, 'VIEW', 'TEAM', 'USER');
                                        }
                                        else if(!viewTeamMember && viewTeamMemberID !== '')
                                        {
                                            RemovePermission(viewTeamMemberID);
                                        }

                                        if(updateTeamMember && updateTeamMemberID === '')
                                        {
                                            AddPermission(id, 'CREATE', 'TEAM', 'USER');
                                        }
                                        else if(!updateTeamMember && updateTeamMemberID !== '')
                                        {
                                            RemovePermission(updateTeamMemberID);
                                        }

                                        if(deleteTeamMember && deleteTeamMemberID === '')
                                        {
                                            AddPermission(id, 'DELETE', 'TEAM', 'USER');
                                        }
                                        else if(!deleteTeamMember && deleteTeamMemberID !== '')
                                        {
                                            RemovePermission(deleteTeamMemberID);
                                        }

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
            fetch("http://deskflow.co.za:8080/api/team/create", 
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
                    //Team
                    if(viewTeam && viewTeamID === '')
                    {
                        AddPermission(id, 'VIEW', 'TEAM', 'IDENTIFIER');
                    }
                    else if(!viewTeam && viewTeamID !== '')
                    {
                        RemovePermission(viewTeamID);
                    }

                    if(updateTeam && updateTeamID === '')
                    {
                        AddPermission(id, 'CREATE', 'TEAM', 'IDENTIFIER');
                    }
                    else if(!updateTeam && updateTeamID !== '')
                    {
                        RemovePermission(updateTeamID);
                    }

                    if(deleteTeam && deleteTeamID === '')
                    {
                        AddPermission(id, 'DELETE', 'TEAM', 'IDENTIFIER');
                    }
                    else if(!deleteTeam && deleteTeamID !== '')
                    {
                        RemovePermission(deleteTeamID);
                    }

                    //Team member
                    if(viewTeamMember && viewTeamMemberID === '')
                    {
                        AddPermission(id, 'VIEW', 'TEAM', 'USER');
                    }
                    else if(!viewTeamMember && viewTeamMemberID !== '')
                    {
                        RemovePermission(viewTeamMemberID);
                    }

                    if(updateTeamMember && updateTeamMemberID === '')
                    {
                        AddPermission(id, 'CREATE', 'TEAM', 'USER');
                    }
                    else if(!updateTeamMember && updateTeamMemberID !== '')
                    {
                        RemovePermission(updateTeamMemberID);
                    }

                    if(deleteTeamMember && deleteTeamMemberID === '')
                    {
                        AddPermission(id, 'DELETE', 'TEAM', 'USER');
                    }
                    else if(!deleteTeamMember && deleteTeamMemberID !== '')
                    {
                        RemovePermission(deleteTeamMemberID);
                    }

                    alert("Team Successfully Updated!");
                    edited(true);
                }
            });
        }
    }

    const AddPermission = (id, type, category, tenant) =>
    {
        console.log(`${id}\n${type}\n${category}\n${tenant}`)
        fetch("http://deskflow.co.za:8080/api/permission/create", 
        {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                permission_id: id,
                permission_id_type: "TEAM",
                permission_type: type,
                permission_category: category,
                permission_tenant: tenant
            }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}`
            }
        })
    }

    const RemovePermission = (id) =>
    {
        console.log(`${id}`)
        fetch("http://deskflow.co.za:8080/api/permission/remove", 
        {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                id: id,
            }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
            }
        });
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

            fetch("http://deskflow.co.za:8080/api/permission/information", 
            {
                method: "POST",
                mode: 'cors',
                body: JSON.stringify({
                    permission_id: team.id,
                    permission_id_type: 'TEAM'
                }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}`
            }
            }).then((res) => res.json()).then(data => 
            {
                setViewTeam(false);
                setUpdateTeam(false);
                setDeleteTeam(false);
                setViewTeamMember(false);
                setUpdateTeamMember(false);
                setDeleteTeamMember(false);

                data.forEach((permission) =>
                {
                    if(permission.permission_type === 'VIEW')
                    {
                        if(permission.permission_category === 'TEAM')
                        {
                            if(permission.permission_tenant === 'IDENTIFIER')
                            {
                                setViewTeam(true);
                                setViewTeamID(permission.id);
                            }
                            else if(permission.permission_tenant === 'USER')
                            {
                                setViewTeamMember(true);
                                setViewTeamMemberID(permission.id);
                            }
                        }
                    }
                    else if(permission.permission_type === 'CREATE')
                    {
                        if(permission.permission_category === 'TEAM')
                        {
                            if(permission.permission_tenant === 'IDENTIFIER')
                            {
                                setUpdateTeam(true);
                                setUpdateTeamID(permission.id);
                            }
                            else if(permission.permission_tenant === 'USER')
                            {
                                setUpdateTeamMember(true);
                                setUpdateTeamMemberID(permission.id);
                            }
                        }
                    }
                    else if(permission.permission_type === 'DELETE')
                    {
                        if(permission.permission_category === 'TEAM')
                        {
                            if(permission.permission_tenant === 'IDENTIFIER')
                            {
                                setDeleteTeam(true);
                                setDeleteTeamID(permission.id);
                            }
                            else if(permission.permission_tenant === 'USER')
                            {
                                setDeleteTeamMember(true);
                                setDeleteTeamMemberID(permission.id);
                            }
                        }
                    }
                });
            });
        }
    }, [team, userData.token]);

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

    useEffect(() =>
    {

    },[userData.token])

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

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <div className={styles.formLabel}>Team permissions</div>

                <div className={styles.checkboxContainer}>
                    <input className={styles.checkbox} ref={viewTeamRef} type='checkbox' checked={viewTeam} onChange={(e) => setViewTeam(e.target.checked)}></input>
                    <label>View team information</label><br></br>
                    
                    <input className={styles.checkbox} ref={updateTeamRef} type='checkbox' checked={updateTeam} onChange={(e) => setUpdateTeam(e.target.checked)}></input>
                    <label>Update team information</label><br></br>

                    <input className={styles.checkbox} ref={deleteTeamRef} type='checkbox' checked={deleteTeam} onChange={(e) => setDeleteTeam(e.target.checked)}></input>
                    <label>Delete team information</label><br></br>

                    <input className={styles.checkbox} ref={viewTeamMemberRef} type='checkbox' checked={viewTeamMember} onChange={(e) => setViewTeamMember(e.target.checked)}></input>
                    <label>View team members</label><br></br>
                    
                    <input className={styles.checkbox} ref={updateTeamMemberRef} type='checkbox' checked={updateTeamMember} onChange={(e) => setUpdateTeamMember(e.target.checked)}></input>
                    <label>Update team members</label><br></br>

                    <input className={styles.checkbox} ref={deleteTeamMemberRef} type='checkbox' checked={deleteTeamMember} onChange={(e) => setDeleteTeamMember(e.target.checked)}></input>
                    <label>Delete team members</label><br></br>
                </div>
            </Form.Group>

            <Form.Group className={styles.formGroup} controlId="formBasicName">
                <form>
                    <div className={styles.formLabel}>Team priority</div>

                    <input ref={priority0Ref} type='radio' name='priority' value={0} onChange={(e) => setPriority(e.target.value)}></input>
                    <label className={styles.radioLabel}>0</label>

                    <input ref={priority1Ref} type='radio' name='priority' value={1} onChange={(e) => setPriority(e.target.value)}></input>
                    <label className={styles.radioLabel}>1</label>

                    <input ref={priority2Ref} type='radio' name='priority' value={2} onChange={(e) => setPriority(e.target.value)}></input>
                    <label className={styles.radioLabel}>2</label>
                </form>
            </Form.Group>


            <Button className={styles.submit} onClick={() => EditTeamSubmit()}>Save</Button>
        </div>
    );

}

export {EditTeam as EditTeamForm}