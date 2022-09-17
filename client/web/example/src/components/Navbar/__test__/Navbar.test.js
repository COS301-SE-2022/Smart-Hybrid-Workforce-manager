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
