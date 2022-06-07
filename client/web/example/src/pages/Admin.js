import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import AdminCard from "../components/AdminCard/AdminCard"

function Admin()
{
  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar/>
        <div className='admin-card-container'>
          <AdminCard name='Users' description='Create and manage users.' 
            path='/users' type='Users'/>
          
          <AdminCard name='Teams' description='Create and manage teams.' 
          path='/team' type='Teams'/>
        </div>
        <div className='admin-card-container'>
          <AdminCard name='Resources' description='Create and manage resources.' 
            path='/resources' type='Resources'/>
          
          <AdminCard name='Roles' description='Create and manage roles.' 
          path='/role' type='Roles'/>
        </div>
      </div>  
      <Footer/>
    </div>
  )
}

export default Admin