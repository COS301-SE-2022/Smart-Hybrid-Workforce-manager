import { BrowserRouter } from 'react-router-dom'
import React, { useContext, useState } from 'react'
import LogoutButton from '../LogoutButton.js'
import {act, render, fireEvent, cleanup, screen} from '@testing-library/react';
import { UserContext } from '../../../App';

import * as router from 'react-router'

const navigate = jest.fn();

const MockLogoutButton = () => {
  const [userData, setUserData] = useState(null);

  return(
    <BrowserRouter>
      <UserContext.Provider value={{userData, setUserData}}>
        <LogoutButton/>
      </UserContext.Provider>
    </BrowserRouter>
  )
};

beforeEach(() => {
  jest.spyOn(router, 'useNavigate').mockImplementation(() => navigate);
  jest.spyOn(window.localStorage.__proto__, 'removeItem');
});

afterEach(cleanup);

describe('When clicking on Logout button', () => {
  it('should navigate to /login', () => {
    render(<MockLogoutButton/> );

    expect(screen.getByTestId('button-user-profile')).toBeInTheDocument();
    fireEvent.click(screen.getByTestId('button-user-profile'));
    expect(navigate).toHaveBeenCalledWith('/login');
  });

  it('should clear local storage', () => {
    render(<MockLogoutButton/> );

    expect(screen.getByTestId('button-user-profile')).toBeInTheDocument();
    fireEvent.click(screen.getByTestId('button-user-profile'));
    expect(localStorage.removeItem).toHaveBeenCalledWith('auth_data');
  });
});
