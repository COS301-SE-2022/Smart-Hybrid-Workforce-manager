import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom'

const Resources = () =>
{
  const navigate = useNavigate();
    useEffect(() =>
    {
        navigate("/layout")
    }, [])

  return (
    <h1>Redirecting...</h1>
  )
}

export default Resources