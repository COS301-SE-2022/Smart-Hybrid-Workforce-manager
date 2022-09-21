import { BrowserRouter } from 'react-router-dom'
import React, { useContext, useState } from 'react'
import UserListItem from '../UserListItem.js'
import {act, render, fireEvent, cleanup, screen} from '@testing-library/react';

import * as router from 'react-router'

const navigate = jest.fn();

const MockUserListItem = ({id, name, email}) => {
  return(
    <BrowserRouter>
        <UserListItem id={id} name = {name} email = {email}/>
    </BrowserRouter>
  )
};

beforeEach(() => {
  jest.spyOn(router, 'useNavigate').mockImplementation(() => navigate);
  jest.spyOn(window.localStorage.__proto__, 'setItem');
});

afterEach(cleanup);

it('should display the name', () => {
  render(<MockUserListItem id='test_id' name='test_name' email='test_email'/> );
  
  expect(screen.getByText('test_name')).toBeInTheDocument();
});

describe('When clicking on popup', () => {
  it('should navigate to /user-edit', () => {
    render(<MockUserListItem id='test_id' name='test_name' email='test_email'/> );

    expect(screen.getByTestId('resource-edit-icon')).toBeInTheDocument();
    fireEvent.click(screen.getByTestId('resource-edit-icon'));
    expect(navigate).toHaveBeenCalledWith('/user-edit');
  });

  it('should set UserID in local storage', () => {
    render(<MockUserListItem id='test_id' name='test_name' email='test_email'/> );

    expect(screen.getByTestId('resource-edit-icon')).toBeInTheDocument();
    fireEvent.click(screen.getByTestId('resource-edit-icon'));
    expect(localStorage.setItem).toHaveBeenCalledWith('UserID', 'test_id');
  });

  it('should set UserName in local storage', () => {
    render(<MockUserListItem id='test_id' name='test_name' email='test_email'/> );

    expect(screen.getByTestId('resource-edit-icon')).toBeInTheDocument();
    fireEvent.click(screen.getByTestId('resource-edit-icon'));
    expect(localStorage.setItem).toHaveBeenCalledWith('UserName', 'test_name');
  });

  it('should set UserEmail in local storage', () => {
    render(<MockUserListItem id='test_id' name='test_name' email='test_email'/> );

    expect(screen.getByTestId('resource-edit-icon')).toBeInTheDocument();
    fireEvent.click(screen.getByTestId('resource-edit-icon'));
    expect(localStorage.setItem).toHaveBeenCalledWith('UserEmail', 'test_email');
  });
});
