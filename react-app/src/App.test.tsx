import React from 'react';
import { render } from '@testing-library/react';
import App from './App';

test('renders profile button', () => {
  const { getByText } = render(<App />);
  const linkElement = getByText(/get profile!/i);
  expect(linkElement).toBeInTheDocument();
});
