import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import Button from 'react-bootstrap/Button'
import { useState, useEffect } from 'react';
import TeamListItem from '../components/Team/TeamListItem';

function Teams()
{
  const [teams, SetTeams] = useState([])

  //POST request
  const FetchTeams = () =>
  {
    fetch("http://localhost:8100/api/team/information", 
        {
          method: "POST",
          body: JSON.stringify({
          })
        }).then((res) => res.json()).then(data => 
          {
            SetTeams(data);
          });
  }

  const AddTeam = () =>
  {
    window.location.assign("./team-create");
  }

  //Using useEffect hook. This will send the POST request once the component is mounted
  useEffect(() =>
  {
    FetchTeams()
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='resources-map'>
          {teams.length > 0 && (
            teams.map(team => 
            {
              return <TeamListItem id={team.id} name={team.name} description={team.description} capacity={team.capacity} picture={team.picture} lead = {team.team_lead_id} priority={team.priority} />
            }
          )
          )}
        </div>

        <div className='button-resource-container'>
          <Button className='button-resource' variant='primary' onClick={AddTeam}>Add Team</Button>
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Teams