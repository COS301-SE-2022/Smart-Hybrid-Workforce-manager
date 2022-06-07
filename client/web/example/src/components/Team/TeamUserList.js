import { useState, useEffect } from 'react';

const TeamUserList = ({id}) =>
{
  const [teamName, SetTeamName] = useState("")

//POST request
  const FetchTeamInformation = () =>
  {
    fetch("http://localhost:8100/api/team/information", 
        {
          method: "POST",
          body: JSON.stringify({
            id:id
          })
        }).then((res) => res.json()).then(data => 
          {
            SetTeamName(data[0].name);
          });
    }
    
    useEffect(() =>
    {
        FetchTeamInformation();
    }, [])

    return (
        <div>
            {teamName}
        </div>
    )
}

export default TeamUserList