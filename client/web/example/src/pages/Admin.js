import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import AdminCard from "../components/AdminCard/AdminCard"

function Admin()
{
  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='card-container'>
          <AdminCard name='Teams' description='Create and manage teams.' 
          path='/teams' image='https://introducingsa.co.za/wp-content/uploads/sites/142/2022/03/Home-office.png'/>

          <AdminCard name='Resources' description='Create and manage resources.' 
          path='/resources' image='https://synivate.com/wp-content/uploads/conference-room-meetings-1-400x249.jpg'/>
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Admin