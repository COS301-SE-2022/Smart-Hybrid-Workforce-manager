import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import { useState, useEffect } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import UserTeamList from '../components/Team/UserTeamList'
import TeamLeadOption from '../components/Team/TeamLeadOption'
import { useNavigate } from 'react-router-dom'

const EditTeam = () =>
{
  const [teamName, setTeamName] = useState(window.sessionStorage.getItem("TeamName"));
  const [teamDescription, setTeamDescription] = useState(window.sessionStorage.getItem("TeamDescription"));
  const [teamCapacity, setTeamCapacity] = useState(window.sessionStorage.getItem("TeamCapacity"));
  const [teamLead, setTeamLead] = useState(window.sessionStorage.getItem("TeamLead"));
  const [teamPriority, setTeamPriority] = useState(window.sessionStorage.getItem("TeamPriority"));

  const [teamUsers, SetTeamUsers] = useState([]);
  //const [viewableUsers, SetViewableUsers] = useState([]);

  const navigate = useNavigate();

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8100/api/team/create", 
      {
        method: "POST",
        body: JSON.stringify({
          id: window.sessionStorage.getItem("TeamID"),
          name: teamName,
          description: teamDescription,
          capacity: parseInt(teamCapacity),
          priority: parseInt(teamPriority),
          team_lead_id: teamLead === "null" ? null : teamLead
        })
      });

      if(res.status === 200)
      {
        alert("Team Successfully Updated!");
        navigate("/team");
        // window.location.assign("./team");
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };

  //POST request
  const FetchTeamUsers = () =>
  {
    fetch("http://localhost:8100/api/team/user/information", 
        {
          method: "POST",
          body: JSON.stringify({
            team_id:window.sessionStorage.getItem("TeamID")
          })
        }).then((res) => res.json()).then(data => 
        {
          SetTeamUsers(data);
        });
  }

  //POST request
  const FetchViewableUsers = () =>
  {
    fetch("http://localhost:8100/api/user/information", 
        {
          method: "POST",
          body: JSON.stringify({})
        }).then((res) => res.json()).then(data => 
        {
          //SetViewableUsers(data);
        });
  }

  //Using useEffect hook. This will set the default values of the form once the components are mounted
  useEffect(() =>
  {
    setTeamName(window.sessionStorage.getItem("TeamName"));
    setTeamDescription(window.sessionStorage.getItem("TeamDescription"));
    setTeamCapacity(window.sessionStorage.getItem("TeamCapacity"));
    setTeamLead(window.sessionStorage.getItem("TeamLead"));
    setTeamPriority(window.sessionStorage.getItem("TeamPriority"));

    FetchTeamUsers();
    FetchViewableUsers();
  }, [])

  const PermissionConfiguration = () =>
  {
    navigate("/team-permissions");
    // window.location.assign("./team-permissions");
  }

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>EDIT TEAM</h1>Please update the team details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Team Name<br></br></Form.Label>
              <Form.Control className='form-input' type="text" placeholder="Enter your team name" value={teamName} onChange={(e) => setTeamName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Description<br></br></Form.Label>
              <Form.Control className='form-input-textarea' as="textarea" rows='5' placeholder="Enter your team description" value={teamDescription} onChange={(e) => setTeamDescription(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicEmail">
              <Form.Label className='form-label'>Capacity<br></br></Form.Label>
              <Form.Control className='form-input' type="text" placeholder="Enter your team capacity" value={teamCapacity} onChange={(e) => setTeamCapacity(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicTeamPriority">
              <Form.Label className='form-label'>Team Priority<br></br></Form.Label>
              <select className='combo-box' name='teampriority' value={teamPriority} onChange={(e) => setTeamPriority(e.target.value)}>
                <option value="0">low</option>
                <option value="1">medium</option>
                <option value="2">high</option>
              </select>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicTeamLead">
              <Form.Label className='form-label'>Team Lead<br></br></Form.Label>
              <select className='combo-box' name='teamlead' value={teamLead} onChange={(e) => setTeamLead(e.target.value)}>
                <option value="null">--none--</option>
                {teamUsers.length > 0 && (
                  teamUsers.map(teamUser => (
                    <TeamLeadOption id={teamUser.user_id} teamLeadId={teamLead} />
                  ))
                )}
              </select>
            </Form.Group>

            <Form.Group className='form-group' controlId="formTeamMembers">
              <Form.Label className='form-label'>Team Members<br></br></Form.Label>
              {teamUsers.length > 0 && (
                  teamUsers.map(teamUser => (
                    <UserTeamList id={teamUser.user_id} />
                  ))
                )}
            </Form.Group>
            
            <Button className='button-submit' variant='primary' type='submit'>Update Team</Button>
            <Button className='button-submit' variant='primary' onClick={PermissionConfiguration}>Configure Permissions</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default EditTeam