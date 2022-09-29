import { useState, useEffect, useContext, useRef, Fragment } from 'react';
import Button from 'react-bootstrap/Button';
import { UserContext } from '../../App';
import { useNavigate } from 'react-router-dom';
import { FaWheelchair, FaHouseUser, FaUserEdit } from 'react-icons/fa';
import styles from './profile.module.css';
import { MdClose } from 'react-icons/md';
import EditProfile from './EditProfile';

const ProfileComponent = () =>
{
    const [user, setUser] = useState({});
    const [roles, setRoles] = useState([]);
    const [teams, setTeams] = useState([]);
    const [edited, setEdited] = useState(true);

    const backgroundDimmerRef = useRef(null);
    const editFormRef = useRef(null);
    const houseRef = useRef(null);
    const wheelchairRef = useRef(null);
  
  const {userData, setUserData}=useContext(UserContext);

  const navigate = useNavigate();
  

    //Using useEffect hook. This will ste the default values of the form once the components are mounted
    useEffect(() =>
    {
        if(edited)
        {
            //POST requests
            const FetchUser = () =>
            {
                fetch("http://deskflow.co.za:8080/api/user/information", 
                {
                    method: "POST",
                    mode: "cors",
                    body: JSON.stringify({
                        identifier : userData.user_identifier
                    }),
                    headers:{
                        'Content-Type': 'application/json',
                        'Authorization': `bearer ${userData.token}`
                    }
                }).then((res) => res.json()).then(data => 
                {
                    console.log(data[0]);
                    setUser(data[0]);
                });
            };

            const FetchUserRoles = () =>
            {
                fetch("http://deskflow.co.za:8080/api/role/user/information", 
                {
                    method: "POST",
                    mode: "cors",
                    body: JSON.stringify({
                        user_id: userData.user_id
                    }),
                    headers:{
                        'Content-Type': 'application/json',
                        'Authorization': `bearer ${userData.token}`
                    }
                }).then((res) => res.json()).then(data => 
                {
                    var userRoles = data;
                    setRoles([]);
                    userRoles.forEach((association) =>
                    {
                        fetch("http://deskflow.co.za:8080/api/role/information", 
                        {
                            method: "POST",
                            mode: "cors",
                            body: JSON.stringify({
                                id: association.role_id
                            }),
                            headers:{
                                'Content-Type': 'application/json',
                                'Authorization': `bearer ${userData.token}`
                            }
                        }).then((res) => res.json()).then(data => 
                        {
                            setRoles((prev) =>
                            (
                                [
                                    ...prev,
                                    data[0]
                                ]
                            ))
                        });
                    });
                });
            };

            const FetchUserTeams = () =>
            {
                fetch("http://deskflow.co.za:8080/api/team/user/information", 
                {
                    method: "POST",
                    mode: "cors",
                    body: JSON.stringify({
                        user_id: userData.user_id
                    }),
                    headers:{
                        'Content-Type': 'application/json',
                        'Authorization': `bearer ${userData.token}`
                    }
                }).then((res) => res.json()).then(data => 
                {
                    setTeams([]);
                    data.forEach((association) =>
                    {
                        fetch("http://deskflow.co.za:8080/api/team/information", 
                        {
                            method: "POST",
                            mode: "cors",
                            body: JSON.stringify({
                                id: association.team_id
                            }),
                            headers:{
                                'Content-Type': 'application/json',
                                'Authorization': `bearer ${userData.token}`
                            }
                        }).then((res) => res.json()).then(data => 
                        {
                            setTeams((prev) =>
                            (
                                [
                                    ...prev,
                                    data[0]
                                ]
                            ));
                        });
                    });
                });
            }

            FetchUser();
            FetchUserRoles();
            FetchUserTeams();
            setEdited(false);
        }
    }, [userData, edited]);

    const OpenEditProfile = () =>
    {
        if(backgroundDimmerRef.current)
        {
            backgroundDimmerRef.current.style.display = 'block';
        }

        if(editFormRef.current)
        {
            editFormRef.current.style.display = 'block';
        }
    }

    const CloseEditProfile = () =>
    {
        if(backgroundDimmerRef.current)
        {
            backgroundDimmerRef.current.style.display = 'none';
        }
        
        if(editFormRef.current)
        {
            editFormRef.current.style.display = 'none';
        }
    }

    useEffect(() =>
    {
        if(houseRef.current)
        {
            if(user.work_from_home)
            {
                houseRef.current.style.display = 'block';
            }
            else
            {
                houseRef.current.style.display = 'none';
            }
        }

        if(wheelchairRef.current)
        {
            if(user.parking === 'DISABLED')
            {
                wheelchairRef.current.style.display = 'block';
            }
            else
            {
                wheelchairRef.current.style.display = 'none';
            }
        }
    },[user, edited])

    return (
        <div className={styles.profileContainer}>

            <div ref={backgroundDimmerRef} className={styles.backgroundDimmer}></div>

            <div ref={editFormRef} className={styles.formContainer}>
                <div className={styles.formClose} onClick={() => CloseEditProfile()}><MdClose /></div>
                <EditProfile user={user} edited={setEdited}/>
            </div>

            <div className={styles.personalInformationContainer}>
                <div className={styles.profileName}>{user.first_name + ' ' + user.last_name}</div>
                <div className={styles.profileEmail}>{user.email}</div>
                <div className={styles.profileIcons}>
                    <div ref={wheelchairRef} className={styles.icon}>
                        <FaWheelchair />
                    </div>
                    
                    <div ref={houseRef} className={styles.icon}>
                        <FaHouseUser />
                    </div>
                </div>
            </div>

            <div className={styles.preferencesContainer}>
                <div className={styles.profileDays}>{user.office_days} office days per week</div>
                
                <div className={styles.profileTime}>Preferred times: {user.preferred_start_time ? user.preferred_start_time.substring(11,16) : '00:00'} - {user.preferred_end_time ? user.preferred_end_time.substring(11,16) : '00:00'}</div>
                
                <div className={styles.profileRoles}>
                    {roles.map((role) => (
                        <div key={role.id} className={styles.pill} style={{backgroundColor: role.color}}>{role.name}</div>
                    ))}
                </div>
            </div>

            <div className={styles.profileTeamsContainer}>
                <div className={styles.profileTeamsTitle}>Teams</div>

                <div className={styles.profileTeamsCarousel}>
                    {teams.map(team => (
                        <div className={styles.teamCard} style={{background: "linear-gradient(180deg, " + team.color + "66  0%, rgba(255,255,255,0.4) 50%)"}}>
                            <img className={styles.teamPicture} src={team.picture} alt='team'></img>
                            <div className={styles.teamName}>{team.name}</div>
                        </div>
                        
                    ))}
                </div>
            </div>

            <div className={styles.profileImageContainer}>
                <img className={styles.image} src={user.picture} alt='user'></img>
            </div>
            <div className={styles.profileEdit} onClick={() => OpenEditProfile()}>
                <FaUserEdit />
            </div>
        </div>
    )
}

export default ProfileComponent