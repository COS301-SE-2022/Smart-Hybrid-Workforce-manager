import { useState, useEffect, useContext, useRef, Fragment } from 'react';
import Button from 'react-bootstrap/Button';
import { UserContext } from '../../App';
import { useNavigate } from 'react-router-dom';
import { FaWheelchair, FaHouseUser, FaUserEdit } from 'react-icons/fa';
import styles from './profile.module.css';
import { MdClose } from 'react-icons/md';

const ProfileComponent = () =>
{
    const [user, setUser] = useState({});
    const [roles, setRoles] = useState([]);
    const [teams, SetTeams] = useState([]);
    const [edited, setEdited] = useState(true);

    const backgroundDimmerRef = useRef(null);
    const editFormRef = useRef(null);
  
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
                fetch("http://localhost:8080/api/user/information", 
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
                    sessionStorage.setItem("UserID", data[0].id);
                    setUser(data[0]);
                });
            };

            const FetchUserRoles = () =>
            {
                fetch("http://localhost:8080/api/role/user/information", 
                {
                    method: "POST",
                    mode: "cors",
                    body: JSON.stringify({
                        //user_id: userData.user_id
                    }),
                    headers:{
                        'Content-Type': 'application/json',
                        'Authorization': `bearer ${userData.token}`
                    }
                }).then((res) => res.json()).then(data => 
                {
                    var userRoles = data;
                    userRoles.forEach((association) =>
                    {
                        fetch("http://localhost:8080/api/role/information", 
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
                fetch("http://localhost:8080/api/team/user/information", 
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
                    SetTeams(data);
                    console.log(data);
                });
            }

            FetchUser();
            FetchUserRoles();
            FetchUserTeams();
        }
    }, [userData, edited]);

    const EditProfile = () =>
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

    const renderWheelchair = () =>
    {
        if(user.parking === 'DISABLED')
        {
            return <FaWheelchair />
        }
    }

    const renderHome = () =>
    {
        if(user.work_from_home)
        {
            return <FaHouseUser />
        }
    }

    return (
        <div className={styles.profileContainer}>

            <div ref={backgroundDimmerRef} className={styles.backgroundDimmer}></div>

            <div ref={editFormRef} className={styles.formContainer}>
                <div className={styles.formClose} onClick={CloseEditProfile}><MdClose /></div>
                <EditProfile user={user} edited={setEdited}/>
            </div>

            <div className={styles.personalInformationContainer}>
                <div className={styles.profileName}>{user.first_name + ' ' + user.last_name}</div>
                <div className={styles.profileEmail}>{user.email}</div>
                <div className={styles.profileIcons}>
                    {renderWheelchair()}
                    {renderHome()}
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
                    {user.teams && (
                        user.teams.map(team =>
                        (
                            <div className={styles.profileTeam}>{team.team_id}</div>
                        ))
                    )}
                </div>
            </div>

            <div className={styles.profileImageContainer}>
                <img className={styles.image} src={user.picture} alt='user'></img>
            </div>
            <div className={styles.profileEdit} onClick={EditProfile.bind(this)}>
                <FaUserEdit />
            </div>
        </div>
    )
}

export default ProfileComponent