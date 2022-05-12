import { render, screen } from '@testing-library/react';
import App from './App';

test('renders learn react link', () => {
  render(<App />);
  const linkElement = screen.getByText(/SMART-HYBRID WORKFORCE MANAGER/i);
  expect(linkElement).toBeInTheDocument();
});
