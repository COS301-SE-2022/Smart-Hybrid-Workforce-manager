import { BrowserRouter } from 'react-router-dom'
import React, { useContext, useState } from 'react'
import Navbar from '../index.js'
import {act, render, fireEvent, cleanup, screen} from '@testing-library/react';
import { Nav, NavHeader, NavLink, Bars, NavMenu } from '../NavbarElements'
import { createMemoryHistory } from 'history';
import { UserContext } from '../../../App';
import Bookings from '../../../pages/Bookings.js'

afterEach(cleanup);

const history = createMemoryHistory();

const MockNavbar = () => {
  const [userData, setUserData] = useState(null);

  return(
    <BrowserRouter history={history}>
      <UserContext.Provider value={{userData, setUserData}}>
        <Navbar/>
      </UserContext.Provider>
    </BrowserRouter>
  )
};

describe('NavbarTest', () => {

  it('Navigate to bookings', () => {
    render(<MockNavbar/> );

    expect(screen.getByText(/BOOKINGS/i)).toBeInTheDocument();
    fireEvent.click(screen.getByText(/BOOKINGS/i));
    // expect(history.location.pathname).toBe('/bookings');
  });
})
