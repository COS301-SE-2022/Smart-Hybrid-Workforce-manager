import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import { useState, useEffect } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'

function ProfileConfiguration()
{
  const [workFromHome, SetWorkFromHome] = useState("")
  const [parking, SetParking] = useState("")
  const [officeDays, SetOfficeDays] = useState("")
  const [startTime, SetStartTime] = useState("")
  const [endTime, SetEndTime] = useState("")

  let handleSubmit = async (e) =>
  {
    alert(workFromHome)
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8100/api/", 
      {
        method: "POST",
        body: JSON.stringify({
        })
      });

      if(res.status === 200)
      {
        alert("Profile Configuration Succesfully Updated!");
        window.location.assign("./");
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
    if (window.sessionStorage.getItem("Parking") === "STANDARD") {
      SetParking(1)
    }
    if (window.sessionStorage.getItem("Parking") === "DISABLED") {
      SetParking(2)
    }
    if (window.sessionStorage.getItem("Parking") === "NONE") {
      SetParking(3)
    }
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
                <option>-- Select Parking Type --</option>
                <option value="1">Standard</option>
                <option value="2">Disabled</option>
                <option value="3">None</option>
              </Form.Select>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Office Days<br></br></Form.Label>
              <Form.Control name="iOfficeDays" className='form-input' type="text" placeholder="3" value={officeDays} onChange={(e) => SetOfficeDays(e.target.value)} />
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