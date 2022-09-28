import { useRef, useState, useEffect, useContext } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from 'react-router-dom'
import { UserContext } from '../App';

function Signup()
{
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const emailRef = useRef();
    const nameRef = useRef();
    const surnameRef = useRef();

    const passwordRef = useRef();
    const lengthRef = useRef();
    const charactersRef = useRef();
    const specialRef = useRef();
    const caseRef = useRef();

    const navigate = useNavigate();

    const {userData} = useContext(UserContext);

    useEffect(() => 
    {
        const validateEmail = new RegExp('^[a-zA-Z0-9.!#$%&\'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:[.][a-zA-Z0-9-]+)+$');

        if(!validateEmail.test(email))
        {
            emailRef.current.style.borderColor = '#ff2e5b';
        }
        else
        {
            emailRef.current.style.borderColor = '#2eff69';
        }
    },[email]);

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

    },[password]);

    useEffect(() => 
    {
        if(firstName.length <= 0)
        {
            nameRef.current.style.borderColor = '#ff2e5b';
        }
        else
        {
            nameRef.current.style.borderColor = '#2eff69';
        }
    },[firstName]);

    useEffect(() => 
    {
        if(lastName.length <= 0)
        {
            surnameRef.current.style.borderColor = '#ff2e5b';
        }
        else
        {
            surnameRef.current.style.borderColor = '#2eff69';
        }
    },[lastName]);


    let handleSubmit = async (e) =>
    {
    e.preventDefault();
    console.log(passwordRef.current.password)
    try
    {
        let res = await fetch("http://localhost:8080/api/user/register", 
        {
        method: "POST",
        mode: "cors",
        body: JSON.stringify({
            id: "",
            first_name: firstName,
            last_name: lastName,
            email: email,
            password: password
        }),
        headers:{
            'Content-Type': 'application/json'
        }
        });

        if(res.status === 200)
        {
            alert("Account Successfully Created!\nPlease verify your login details");
            navigate("/login");
        }
    }
    catch(err)
    {
        alert("Account Successfully Created!\nPlease verify your login details");
        navigate("/login");
        console.log(err);
    }
    };  

    return (
    <div className='page-container'>
        <div className='content-login'>
        <div className='login-grid'>
            <div className='form-container-signup'>
                <p className='form-header'><h1>CREATE AN ACCOUNT</h1>Please enter your details.</p>
                
                <Form className='form' onSubmit={handleSubmit}>
                    <Form.Group className='form-group' controlId="formBasicName">
                        <Form.Label className='form-label'>First Name<br></br></Form.Label>
                        <Form.Control ref={nameRef} className='form-input' type="text" placeholder="Enter your first name" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
                    </Form.Group>

                    <Form.Group className='form-group' controlId="formBasicName">
                        <Form.Label className='form-label'>Surname<br></br></Form.Label>
                        <Form.Control ref={surnameRef} className='form-input' type="text" placeholder="Enter your surname" value={lastName} onChange={(e) => setLastName(e.target.value)} />
                    </Form.Group>

                    <Form.Group className='form-group' controlId="formBasicEmail">
                        <Form.Label className='form-label'>Email<br></br></Form.Label>
                        <Form.Control ref={emailRef} className='form-input' type="text" placeholder="Enter your email" value={email} onChange={(e) => setEmail(e.target.value)} />
                    </Form.Group>

                    <Form.Group className='form-group' controlId="formBasicPassword">
                        <Form.Label className='form-label'>Password<br></br></Form.Label>
                        <Form.Control ref={passwordRef} className='password-input' type="password" placeholder="Enter your password" value={password} onChange={(e) => setPassword(e.target.value)} />
                        <div className='password-requirements'>
                            <ul>
                                <li ref={lengthRef}>8 or more characters</li>
                                <li ref={charactersRef}>Letters and numbers</li>
                                <li ref={caseRef}>Uppercase and lowercase</li>
                                <li ref={specialRef}>Special characters</li>
                            </ul>
                        </div>
                    </Form.Group>

                    <Button className='button-submit' variant='primary' type='submit'>Create Account</Button>
                </Form>
            </div>

            

            <div className='image-container'>
                <img className='login-image' src='https://i.pinimg.com/originals/3b/79/c7/3b79c7a4a275b5ee1dbb76731f9736b8.png' alt='office'></img>
            </div>
        </div>

        </div>
    </div>
    )
}

export default Signup