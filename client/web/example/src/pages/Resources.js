import { useEffect } from 'react';

const Resources = () =>
{

    useEffect(() =>
    {
        window.location.replace('http://localhost:3000/layout');
    }, [])

  return (
    <h1>Redirecting...</h1>
  )
}

export default Resources