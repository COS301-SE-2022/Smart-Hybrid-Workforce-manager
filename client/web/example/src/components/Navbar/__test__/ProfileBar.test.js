import { BrowserRouter } from 'react-router-dom'
import React, { useContext, useState } from 'react'
import ProfileBar from '../ProfileBar.js'
import {act, render, fireEvent, cleanup, screen} from '@testing-library/react';
import { UserContext } from '../../../App';

import * as router from 'react-router'

const navigate = jest.fn();

const MockProfileBar = () => {
  const [userData, setUserData] = useState(null);

  return(
    <BrowserRouter>
      <UserContext.Provider value={{userData, setUserData}}>
        <ProfileBar/>
      </UserContext.Provider>
    </BrowserRouter>
  )
};

beforeEach(() => {
  jest.spyOn(router, 'useNavigate').mockImplementation(() => navigate);
  jest.spyOn(window.localStorage.__proto__, 'removeItem');
});

afterEach(cleanup);

describe('When clicking on Profile picture', () => {
  it('should navigate to /profile', () => {
    render(<MockProfileBar/> );

    expect(screen.getByTestId('profilepic-container')).toBeInTheDocument();
    fireEvent.click(screen.getByTestId('profilepic-container'));
    expect(navigate).toHaveBeenCalledWith('/profile');
  });
});

describe('When clicking on Logout', () => {
  it('should navigate to /login', () => {
    render(<MockProfileBar/> );

    expect(screen.getByText(/Logout/i)).toBeInTheDocument();
    fireEvent.click(screen.getByText(/Logout/i));
    expect(navigate).toHaveBeenCalledWith('/login');
  });

  it('should clear local storage', () => {
    render(<MockProfileBar/> );

    expect(screen.getByText(/Logout/i)).toBeInTheDocument();
    fireEvent.click(screen.getByText(/Logout/i));
    expect(localStorage.removeItem).toHaveBeenCalledWith('auth_data');
  });
});
