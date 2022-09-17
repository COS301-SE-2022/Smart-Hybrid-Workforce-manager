import { BrowserRouter } from 'react-router-dom'
import React, { useContext, useState } from 'react'
import AdminCard from '../AdminCard.js'
import {act, render, fireEvent, cleanup, screen} from '@testing-library/react';
import { UserContext } from '../../../App';

import * as router from 'react-router'

const navigate = jest.fn();

const MockAdminCard = ({name, description, path, type}) => {
  const [userData, setUserData] = useState(null);

  return(
    <BrowserRouter>
      <UserContext.Provider value={{userData, setUserData}}>
        <AdminCard name={name} description={description} path={path} type={type} />
      </UserContext.Provider>
    </BrowserRouter>
  )
};

beforeEach(() => {
  jest.spyOn(router, 'useNavigate').mockImplementation(() => navigate);
  jest.spyOn(window.localStorage.__proto__, 'removeItem');
});

afterEach(cleanup);

describe('When rendering as Users type', () => {
  it('should return Users icon', () => {
    render(<MockAdminCard name='Users' description='Create and manage users.' path='/users' type='Users'/> );
    expect(screen.getByTestId('admin-icon-users')).toBeInTheDocument();
  });

  it('should display the correct header', () => {
    render(<MockAdminCard name='Users' description='Create and manage users.' path='/users' type='Users'/> );

    expect(screen.getByTestId('admin-card-text-header')).toBeInTheDocument();
    expect(screen.getByTestId('admin-card-text-header').textContent).toBe('Users');
  });

  it('should display the correct body', () => {
    render(<MockAdminCard name='Users' description='Create and manage users.' path='/users' type='Users'/> );

    expect(screen.getByTestId('admin-card-text-body')).toBeInTheDocument();
    expect(screen.getByTestId('admin-card-text-body').textContent).toBe('Create and manage users.');
  });
});