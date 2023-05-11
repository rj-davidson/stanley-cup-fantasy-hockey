import React, { ReactNode } from 'react';
import { makeStyles } from '@mui/styles';
import { Container, Grid } from '@mui/material';

const useStyles = makeStyles(() => ({
  root: {
    flexGrow: 1,
    paddingTop: '20px',
    paddingBottom: '20px',
  },
}));

interface Props {
  children: ReactNode;
}

export default function Body({ children }: Props) {
  const classes = useStyles();

  return (
    <Container maxWidth="lg" className={classes.root}>
      <Grid marginTop={7}> </Grid>
      {children}
    </Container>
  );
}
