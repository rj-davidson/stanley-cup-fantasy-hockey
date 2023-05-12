import React, { ReactNode } from 'react';
import Navigation from '@/components/layout/navigation';
import Header from '@/components/layout/header';
import Body from '@/components/layout/body';
import Footer from '@/components/layout/footer';
import { Grid } from '@mui/material';

interface LayoutProps {
  children: ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <Grid container direction="column" style={{ minHeight: '100vh' }}>
      <Grid item>
        <Navigation />
        <Header />
        <Body>{children}</Body>
      </Grid>
      <Grid item>
        <Footer />
      </Grid>
    </Grid>
  );
};

export default Layout;
