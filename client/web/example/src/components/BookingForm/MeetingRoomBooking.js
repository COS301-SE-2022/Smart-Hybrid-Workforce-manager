import React, { useState, useEffect, useContext } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from "react-router-dom"
import { UserContext } from '../../App'

const MeetingRoomBooking = (props, ref) =>
{
    const [startDate, setStartDate] = useState("");
    const [startTime, setStartTime] = useState("");
    const [endTime, setEndTime] = useState("");

    const [teams, SetTeams] = useState([]);
    const [roles, SetRoles] = useState([]);

    const [teamSelectedId, SetTeamSelectedId] = useState(null) // explicit nulling
    const [roleSelectedId, SetRoleSelectedId] = useState(null) // explicit nulling

    const [aditionalAttendees, SetAditionalAttendees] = useState(0) // Use a number not string
    const [attendeesDesks, SetAttendeesDesks] = useState(false) // Use a bool
    const [aditionalAttendeesDesks, SetAditionalAttendeesDesks] = useState(false) // Use a bool
    
    const {userData} = useContext(UserContext);

    const navigate = useNavigate();

    let handleSubmit = async (e) =>
    {
        e.preventDefault();
        try
        {
            let res = await fetch("http://localhost:8080/api/booking/meetingroom/create", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify(
                {
                    "booking": {
                    id: null,
                    user_id: userData.user_id,
                    resource_type: "MEETINGROOM",
                    resource_preference_id: null,
                    resource_id: null,
                    start: startDate + "T" + startTime + ":43.511Z",
                    end: startDate + "T" + endTime + ":43.511Z",
                    booked: false
                    },
                    team_id: (teamSelectedId === "null") ? null : teamSelectedId,
                    role_id: (roleSelectedId === "null") ? null : roleSelectedId,
                    additional_attendees: Number(aditionalAttendees),
                    desks_attendees: attendeesDesks,
                    desks_aditional_attendees: aditionalAttendeesDesks,
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                }
            });

            if(res.status === 200)
            {
                alert("Booking Successfully Created!");
                navigate("/bookings-meeting");
            }
        }
        catch(err)
        {
            console.log(err);
        }
    }; 

    //POST request
    const FetchTeams = () =>
    {
        fetch("http://localhost:8080/api/team/information", 
        {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({}),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
            }
        }).then((res) => res.json()).then(data => 
        {
            SetTeams(data);
        });
    }

    //POST request
    const FetchRoles = () =>
    {
        fetch("http://localhost:8080/api/role/information", 
        {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({}),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
            }
        }).then((res) => res.json()).then(data => 
        {
            SetRoles(data);
        });
    }

    //Using useEffect hook. This will set the default values of the form once the components are mounted
    useEffect(() =>
    {
        FetchTeams();
        FetchRoles();
    }, [])

    return (
        <div ref={ref} className='form-container-meetingroom-booking'>
            <div className='form-header'><h1>MEETING ROOM BOOKING</h1>Please enter your booking details.</div>
            
            <Form className='form' onSubmit={handleSubmit}>

                <Form.Group className='form-group' controlId="formBasicTeam">
                    <Form.Label className='form-label'>Team<br></br></Form.Label>
                    <select className='combo-box' name='teamId' value={teamSelectedId} onChange={(e) => SetTeamSelectedId(e.target.value)}>
                        <option value="null">--none--</option>
                        {teams.length > 0 && (
                            teams.map(team => (
                                <option value={team.id}>{ team.name }</option>
                            ))
                        )}
                    </select>
                </Form.Group>

                <Form.Group className='form-group' controlId="formBasicRole">
                    <Form.Label className='form-label'>Role<br></br></Form.Label>
                    <select className='combo-box' name='roleId' value={roleSelectedId} onChange={(e) => SetRoleSelectedId(e.target.value)}>
                        <option value="null">--none--</option>
                        {roles.length > 0 && (
                            roles.map(role => (
                                <option value={role.id}>{ role.role_name }</option>
                            ))
                        )}
                    </select>
                </Form.Group>

                <Form.Group className='form-group' controlId="formBasicName">
                    <Form.Label className='form-label'>Aditional Attendees Count<br></br></Form.Label>
                    <Form.Control name="sAttendees" className='form-input' type="number" placeholder="0" min="0" value={aditionalAttendees} onChange={(e) => SetAditionalAttendees(e.target.value)} />
                </Form.Group>

                <Form.Group className='form-group' controlId="formBasicName">
                    <label className="container">
                        Book Desks for attendees  
                        <input className='checkbox' type="checkbox" checked={attendeesDesks} onChange={(e) => SetAttendeesDesks(e.target.checked)}/>
                        <span className="checkmark"></span>
                    </label>
                </Form.Group>
                
                <Form.Group className='form-group' controlId="formBasicName">
                    <label className="container">
                        Book Desks for additional attendees  
                        <input className='checkbox' type="checkbox" checked={aditionalAttendeesDesks} onChange={(e) => SetAditionalAttendeesDesks(e.target.checked)}/>
                        <span className="checkmark"></span>
                    </label>
                </Form.Group>
                
                <Form.Group className='form-group' controlId="formBasicName">
                    <Form.Label className='form-label'>Date<br></br></Form.Label>
                    <Form.Control className='form-input' type="date" value={startDate} onChange={(e) => setStartDate(e.target.value)} />
                </Form.Group>

                <Form.Group className='form-group' controlId="formBasicName">
                    <Form.Label className='form-label'>Start Time<br></br></Form.Label>
                    <Form.Control className='form-input' type="time" placeholder="hh:mm" value={startTime} onChange={(e) => setStartTime(e.target.value)} />
                </Form.Group>
                
                <Form.Group className='form-group' controlId="formBasicName">
                    <Form.Label className='form-label'>End Time<br></br></Form.Label>
                    <Form.Control className='form-input' type="time" placeholder="hh:mm" value={endTime} onChange={(e) => setEndTime(e.target.value)} />
                </Form.Group>

                <Button className='button-submit' variant='primary' type='submit'>Create Booking</Button>
            </Form>
        </div>
    )
}

export default React.forwardRef(MeetingRoomBooking)