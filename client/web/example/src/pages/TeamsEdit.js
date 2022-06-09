import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import { useState, useEffect } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'

const EditTeam = () =>
{
  const [teamName, setTeamName] = useState("");
  const [teamDescription, setTeamDescription] = useState("");
  const [teamCapacity, setTeamCapacity] = useState("");

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
          teamCapacity: teamCapacity
        })
      });

      if(res.status === 200)
      {
        alert("Team Successfully Updated!");
        window.location.assign("./team");
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };

  //Using useEffect hook. This will ste the default values of the form once the components are mounted
  useEffect(() =>
  {
    setTeamName(window.sessionStorage.getItem("TeamName"));
    setTeamDescription(window.sessionStorage.getItem("TeamDescription"));
    setTeamCapacity(window.sessionStorage.getItem("TeamCapacity"));
  }, [])

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

            <Button className='button-submit' variant='primary' type='submit'>Update Team</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default EditTeam