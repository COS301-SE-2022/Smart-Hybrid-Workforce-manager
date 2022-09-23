import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import { useState, useEffect, useContext } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'
import { useNavigate } from "react-router-dom"
import { UserContext } from '../App.js'

function ProfileConfiguration()
{
  const [identifier, SetIdentifier] = useState("")
  const [firstName, SetFirstName] = useState("")
  const [lastName, SetLastName] = useState("")
  const [email, SetEmail] = useState("")
  const [picture, SetPicture] = useState("")
  const [dateCreated, SetDateCreated] = useState("")
  const [workFromHome, SetWorkFromHome] = useState("")
  const [parking, SetParking] = useState("")
  const [officeDays, SetOfficeDays] = useState("")
  const [startTime, SetStartTime] = useState("")
  const [endTime, SetEndTime] = useState("")

  const navigate=useNavigate();
  const {userData,setUserData}=useContext(UserContext)

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8080/api/user/update", 
      {
        method: "POST",
        mode: 'cors',
        body: JSON.stringify({
          id: window.sessionStorage.getItem("UserID"),
          identifier: identifier,
          first_name: firstName,
          last_name: lastName,
          email: email,
          picture: picture,
          date_created: dateCreated,
          work_from_home: workFromHome,
          parking: parking,
          office_days: parseInt(officeDays),
          preferred_start_time: "0000-01-01T" + startTime + ":00Z",
          preferred_end_time: "0000-01-01T" + endTime + ":00Z"
        }),
        headers:{
            'Content-Type': 'application/json',
            'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
        }
      });

      if(res.status === 200)
      {
        alert("Profile Configuration Succesfully Updated!");
        navigate("/")
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
    SetParking(window.sessionStorage.getItem("Parking"))
    SetIdentifier(window.sessionStorage.getItem("Identifier"))
    SetFirstName(window.sessionStorage.getItem("FirstName"))
    SetLastName(window.sessionStorage.getItem("LastName"))
    SetEmail(window.sessionStorage.getItem("Email"))
    SetPicture(window.sessionStorage.getItem("Picture"))
    SetDateCreated(window.sessionStorage.getItem("DateCreated"))
    SetWorkFromHome(window.sessionStorage.getItem("WorkFromHome"))
    SetOfficeDays(window.sessionStorage.getItem("OfficeDays"))
    SetStartTime(window.sessionStorage.getItem("StartTime").substring(11,16))
    SetEndTime(window.sessionStorage.getItem("EndTime").substring(11,16))
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>Configure Profile</h1>Please enter your profile configuration settings.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Work From Home&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;</Form.Label>
              <input type="checkbox" defaultChecked={workFromHome} onChange={(e) => SetWorkFromHome(e.target.checked)}/>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Parking Type<br></br></Form.Label>
              <Form.Select aria-label="Default select example" className='form-input' onChange={(e) => SetParking(e.target.value)} value={parking}>
                <option value="STANDARD">Standard</option>
                <option value="DISABLED">Disabled</option>
                <option value="NONE">None</option>
              </Form.Select>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Office Days<br></br></Form.Label>
              <Form.Control name="iOfficeDays" className='form-input' type="number" placeholder="0" min="0" value={officeDays} onChange={(e) => SetOfficeDays(e.target.value)} />
            </Form.Group>
            
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Preffered Start Time<br></br></Form.Label>
              <Form.Control name="eStartTime" className='form-input' type="text" placeholder="hh:mm" value={startTime} onChange={(e) => SetStartTime(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Preffered End Time<br></br></Form.Label>
              <Form.Control name="eEndTime" className='form-input' type="text" placeholder="hh:mm" value={endTime} onChange={(e) => SetEndTime(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Update Profile</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default ProfileConfiguration