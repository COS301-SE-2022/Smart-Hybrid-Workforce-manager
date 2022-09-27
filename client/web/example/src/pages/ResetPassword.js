import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import { useRef, useState, useEffect, useContext } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'
import { useNavigate } from "react-router-dom"
import { UserContext } from '../App.js'

function ResetPassword()
{
    const [identifier, SetIdentifier] = useState("")
    const [password, SetPassword] = useState("")
    const [confirm, SetConfirmPassword] = useState("")
    
    const passwordRef = useRef();
    const confirmPasswordRef = useRef();
    const lengthRef = useRef();
    const charactersRef = useRef();
    const specialRef = useRef();
    const caseRef = useRef();

    const navigate=useNavigate();
    const {userData,setUserData}=useContext(UserContext)

  let handleSubmit = async (e) =>
  {
      e.preventDefault();
      if (password != confirm) {
          alert("Passwords do not match!")
          return
      }
    try
    {
      let res = await fetch("http://localhost:8080/api/user/resetpassword", 
      {
        method: "POST",
        mode: 'cors',
        body: JSON.stringify({
          user_id: window.sessionStorage.getItem("UserID"),
          password: password,
        }),
        headers:{
            'Content-Type': 'application/json',
            'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
        }
      });

      if(res.status === 200)
      {
        alert("Password Succesfully Updated!");
        navigate("/")
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };
    
    useEffect(() => 
    {
        const validatePassword = new RegExp('^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,}$');

        const CheckLowerCase = () =>
        {
            for(var i = 0; i < password.length; i++)
            {
                if(password.charAt(i).match(/[A-z]/i) && password.charAt(i) === password.charAt(i).toLowerCase())
                {
                    return true;
                }
            }

            return false;
        }

        const CheckUpperCase = () =>
        {
            for(var i = 0; i < password.length; i++)
            {
                if(password.charAt(i).match(/[A-z]/i) && password.charAt(i) === password.charAt(i).toUpperCase())
                {
                    return true;
                }
            }

            return false;
        }

        const CheckCharacters = () =>
        {
            for(var i = 0; i < password.length; i++)
            {
                if(password.charAt(i).match(/[A-z]/i))
                {
                    for(var j = 0; j < password.length; j++)
                    {
                        if(password.charAt(j).match(/[0-9]/i))
                        {
                            return true;
                        }
                    }
                }
            }

            return false;
        }

        const CheckSpecial = () =>
        {
            for(var i = 0; i < password.length; i++)
            {
                if(password.charAt(i).match(/[!@#$%^&*]/i))
                {
                    return true;
                }
            }

            return false;
        }

        if(!validatePassword.test(password))
        {
            passwordRef.current.style.borderColor = '#ff2e5b';
        }
        else
        {
            passwordRef.current.style.borderColor = '#2eff69';
        }

        if(password.length >= 8)
        {
            lengthRef.current.style.color = '#2eff69';
        }
        else
        {
            lengthRef.current.style.color = '#ff2e5b';
        }

        if(CheckCharacters())
        {
            charactersRef.current.style.color = '#2eff69';
        }
        else
        {
            charactersRef.current.style.color = '#ff2e5b';
        }

        if(CheckLowerCase() && CheckUpperCase())
        {
            caseRef.current.style.color = '#2eff69';
        }
        else
        {
            caseRef.current.style.color = '#ff2e5b';
        }

        if(CheckSpecial())
        {
            specialRef.current.style.color = '#2eff69';
        }
        else
        {
            specialRef.current.style.color = '#ff2e5b';
        }

    }, [password]);

  //Using useEffect hook. This will ste the default values of the form once the components are mounted
  useEffect(() =>
  {
    SetIdentifier(window.sessionStorage.getItem("Identifier"))
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
            <p className='form-header'><h1>Reset Password</h1>Please enter your new password and confirm.</p>
            
            <Form className='form' onSubmit={handleSubmit}>
                      
            <Form.Group className='form-group' controlId="formBasicPassword">
                <Form.Label className='form-label'>Password<br></br></Form.Label>
                <Form.Control ref={passwordRef} className='password-input' type="password" placeholder="Enter your password" value={password} onChange={(e) => SetPassword(e.target.value)} />
                <div className='password-requirements'>
                    <ul>
                        <li ref={lengthRef}>8 or more characters</li>
                        <li ref={charactersRef}>Letters and numbers</li>
                        <li ref={caseRef}>Uppercase and lowercase</li>
                        <li ref={specialRef}>Special characters</li>
                    </ul>
                </div>
            </Form.Group>
                  
            <Form.Group className='form-group' controlId="formBasicConfirmPassword">
                <Form.Label className='form-label'>Password<br></br></Form.Label>
                <Form.Control ref={confirmPasswordRef} className='password-input' type="password" placeholder="Confirm your password" value={confirm} onChange={(e) => SetConfirmPassword(e.target.value)}/>
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Update Password</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default ResetPassword