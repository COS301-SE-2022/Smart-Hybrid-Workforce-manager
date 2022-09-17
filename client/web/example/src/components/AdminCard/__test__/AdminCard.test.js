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
  jest.spyOn(window, 'open');
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

  it('should navigate to /users', () => {
    render(<MockAdminCard name='Users' description='Create and manage users.' path='/users' type='Users'/> );

    expect(screen.getByTestId('admin-card')).toBeInTheDocument();
    fireEvent.click(screen.getByTestId('admin-card'));
    expect(navigate).toHaveBeenCalledWith('/users');
    expect(window.open).not.toHaveBeenCalled();
  });
});

describe('When rendering as Teams type', () => {
  it('should return Users icon', () => {
    render(<MockAdminCard name='Teams' description='Create and manage teams.' path='/team' type='Teams'/> );
    expect(screen.getByTestId('admin-icon-teams')).toBeInTheDocument();
  });

  it('should display the correct header', () => {
    render(<MockAdminCard name='Teams' description='Create and manage teams.' path='/team' type='Teams'/> );

    expect(screen.getByTestId('admin-card-text-header')).toBeInTheDocument();
    expect(screen.getByTestId('admin-card-text-header').textContent).toBe('Teams');
  });

  it('should display the correct body', () => {
    render(<MockAdminCard name='Teams' description='Create and manage teams.' path='/team' type='Teams'/> );

    expect(screen.getByTestId('admin-card-text-body')).toBeInTheDocument();
    expect(screen.getByTestId('admin-card-text-body').textContent).toBe('Create and manage teams.');
  });

  it('should navigate to /team', () => {
    render(<MockAdminCard name='Teams' description='Create and manage teams.' path='/team' type='Teams'/> );

    expect(screen.getByTestId('admin-card')).toBeInTheDocument();
    fireEvent.click(screen.getByTestId('admin-card'));
    expect(navigate).toHaveBeenCalledWith('/team');
    expect(window.open).not.toHaveBeenCalled();
  });
});

describe('When rendering as Resources type', () => {
  it('should return Users icon', () => {
    render(<MockAdminCard name='Resources' description='Create and manage resources.' path='/resources' type='Resources'/> );
    expect(screen.getByTestId('admin-icon-resources')).toBeInTheDocument();
  });

  it('should display the correct header', () => {
    render(<MockAdminCard name='Resources' description='Create and manage resources.' path='/resources' type='Resources'/> );

    expect(screen.getByTestId('admin-card-text-header')).toBeInTheDocument();
    expect(screen.getByTestId('admin-card-text-header').textContent).toBe('Resources');
  });

  it('should display the correct body', () => {
    render(<MockAdminCard name='Resources' description='Create and manage resources.' path='/resources' type='Resources'/> );

    expect(screen.getByTestId('admin-card-text-body')).toBeInTheDocument();
    expect(screen.getByTestId('admin-card-text-body').textContent).toBe('Create and manage resources.');
  });

  it('should call window.open to /resources', () => {
    render(<MockAdminCard name='Resources' description='Create and manage resources.' path='/resources' type='Resources'/> );

    expect(screen.getByTestId('admin-card')).toBeInTheDocument();
    fireEvent.click(screen.getByTestId('admin-card'));
    expect(window.open).toHaveBeenCalledWith('/resources');
    expect(navigate).not.toHaveBeenCalled();
  });
});

describe('When rendering as Roles type', () => {
  it('should return Users icon', () => {
    render(<MockAdminCard name='Roles' description='Create and manage roles.' path='/role' type='Roles'/> );
    expect(screen.getByTestId('admin-icon-default')).toBeInTheDocument();
  });

  it('should display the correct header', () => {
    render(<MockAdminCard name='Roles' description='Create and manage roles.' path='/role' type='Roles'/> );

    expect(screen.getByTestId('admin-card-text-header')).toBeInTheDocument();
    expect(screen.getByTestId('admin-card-text-header').textContent).toBe('Roles');
  });

  it('should display the correct body', () => {
    render(<MockAdminCard name='Roles' description='Create and manage roles.' path='/role' type='Roles'/> );

    expect(screen.getByTestId('admin-card-text-body')).toBeInTheDocument();
    expect(screen.getByTestId('admin-card-text-body').textContent).toBe('Create and manage roles.');
  });

  it('should navigate to /role', () => {
    render(<MockAdminCard name='Roles' description='Create and manage roles.' path='/role' type='Roles'/> );

    expect(screen.getByTestId('admin-card')).toBeInTheDocument();
    fireEvent.click(screen.getByTestId('admin-card'));
    expect(navigate).toHaveBeenCalledWith('/role');
    expect(window.open).not.toHaveBeenCalled();
  });
});
