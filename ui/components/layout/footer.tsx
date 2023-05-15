import React from 'react';
import { Container, Typography } from '@mui/material';

const Footer: React.FC = () => {
  const currentYear = new Date().getFullYear();
  return (
    <footer>
      <Container maxWidth="md">
        <Typography variant="body1" align="center">
          &copy; {currentYear} hockey.rjd.app - All rights reserved.
        </Typography>
      </Container>
    </footer>
  );
};

export default Footer;
