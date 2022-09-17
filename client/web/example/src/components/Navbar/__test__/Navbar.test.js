import { BrowserRouter } from 'react-router-dom'
import React, { useContext, useState } from 'react'
import Navbar from '../Navbar.js'
import {act, render, fireEvent, cleanup, screen} from '@testing-library/react';
import { UserContext } from '../../../App';

import * as router from 'react-router'

const navigate = jest.fn();

const MockNavbar = () => {
  const [userData, setUserData] = useState(null);

  return(
    <BrowserRouter>
      <UserContext.Provider value={{userData, setUserData}}>
        <Navbar/>
      </UserContext.Provider>
    </BrowserRouter>
  )
};

beforeEach(() => {
  jest.spyOn(router, 'useNavigate').mockImplementation(() => navigate)
});

afterEach(cleanup);

describe('When clicking on Calendar', () => {
  it('should navigate to /', () => {
    render(<MockNavbar/> );

    expect(screen.getByText(/Calendar/i)).toBeInTheDocument();
    fireEvent.click(screen.getByText(/Calendar/i));
    expect(navigate).toHaveBeenCalledWith('/');
  });
});

describe('When clicking on Office Map', () => {
  it('should navigate to /map', () => {
    render(<MockNavbar/> );

    expect(screen.getByText(/Office Map/i)).toBeInTheDocument();
    fireEvent.click(screen.getByText(/Office Map/i));
    expect(navigate).toHaveBeenCalledWith('/map');
  });
});

describe('When clicking on Statistics', () => {
  it('should navigate to /statistics', () => {
    render(<MockNavbar/> );

    expect(screen.getByText(/Statistics/i)).toBeInTheDocument();
    fireEvent.click(screen.getByText(/Statistics/i));
    expect(navigate).toHaveBeenCalledWith('/statistics');
  });
});

describe('When clicking on Bookings', () => {
  it('should open drop down menu', () => {
    render(<MockNavbar/> );

    expect(screen.getByText(/Bookings/i)).toBeInTheDocument();
    expect(screen.getByText(/Desk/i)).not.toBeVisible();
    expect(screen.getByText(/Meeting Room/i)).not.toBeVisible();

    fireEvent.click(screen.getByText(/Bookings/i));
    expect(screen.getByText(/Desk/i)).toBeVisible();
    expect(screen.getByText(/Meeting Room/i)).toBeVisible();
  });

  it('should navigate to /bookings-desk', () => {
    render(<MockNavbar/> );

    fireEvent.click(screen.getByText(/Desk/i));
    expect(navigate).toHaveBeenCalledWith('/bookings-desk');
  });

  it('should navigate to /bookings-meetingroom', () => {
    render(<MockNavbar/> );

    fireEvent.click(screen.getByText(/Meeting Room/i));
    expect(navigate).toHaveBeenCalledWith('/bookings-meetingroom');
  });
});
