import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import React, { useContext, useState } from 'react'
import Login from '../pages/Login.js'
import {act, render, fireEvent, cleanup, screen} from '@testing-library/react';
import { UserContext } from '../App';
import { createMemoryHistory } from 'history';

// global.fetch = jest.fn(() => Promise.resolve({
//     json: () => Promise.resolve([]),
//   })
// )

// beforeAll(() => {
//   global.fetch = () => 
//     Promise.resolve({status: 200,
//       json: () => Promise.resolve()
//   })
// })

// beforeEach(() => {
//   // fetch.resetMocks();
//   fetch.mockClear();
// });

// const fetchMock = jest.spyOn(global, 'fetch').mockImplementation(() => Promise.resolve(200, { json: () => Promise.resolve({}) }))

// const mockFetch = jest.fn().mockImplementation(() => ({
//   json: jest.fn().mockResolvedValue({}),
// }));

// jest.mock('node-fetch', () => {
//   return jest.fn().mockImplementation((...p) => {
//     return mockFetch(...p);
//   });
// });

beforeEach(() => {
  // window.alert = jest.fn();
});

afterEach(cleanup);

const history = createMemoryHistory();

const MockLogin = () => {
  const [userData, setUserData] = useState(null);

  return(
    <Router history={history}>
      <UserContext.Provider value={{userData, setUserData}}>
        <Login/>
      </UserContext.Provider>
    </Router>
  )
};

describe('LoginTest', () => {
  
  jest.spyOn(Object.getPrototypeOf(window.sessionStorage), 'setItem')
  Object.setPrototypeOf(window.sessionStorage.setItem, jest.fn())

  it('Failed login', () => {
    global.fetch = () => 
      Promise.resolve({status: 401,
        json: () => Promise.resolve()
    })

    render(<MockLogin />);
    // const sesh = sessionStorage.getItem('auth_data')
    // jest.spyOn(global, 'alert').mockImplementation(() => {})
    // Object.defineProperty(sessionStorage, "setItem", { writable: true });
    // Object.defineProperty(window, 'sessionStorage', { value: mock,configurable:true,enumerable:true,writable:true });
    // jest.spyOn(sessionStorage, "setItem")

  //   jest.spyOn(Object.getPrototypeOf(window.localStorage), 'setItem')
  // Object.setPrototypeOf(window.localStorage.setItem, jest.fn())

        // const asdfauth = sessionStorage.getItem("auth_data");
    // history.push('/login')

    fireEvent.change(screen.getByPlaceholderText(/Enter your email/i), { target: { value: 'failEmail' } })
    fireEvent.change(screen.getByPlaceholderText(/Enter your password/i), { target: { value: 'failPass' } })

    fireEvent.click(screen.getByText(/Sign In/i));

    // expect(screen.getByText(/Fail/i)).toBeInTheDocument()
    // expect(window.alert).toBeCalled()
    // expect(history.location.pathname).toBe('/asdfasdfasdf');
    // expect(sessionStorage.getItem).toReturn(sesh)
    expect(window.sessionStorage.setItem).not.toHaveBeenCalled()

    // expect(window.alert).toHaveBeenCalledWith("Successfully Logged In!");
    // expect(global.alert).toHaveBeenCalled();
    // expect(navigate).toHvaeBeenCalled();
    // expect(jsdomAlert).toHaveBeenCalledWith("Successfully Logged In!");
  });

  
})