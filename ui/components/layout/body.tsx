import React, { ReactNode } from 'react';
import { makeStyles } from '@mui/styles';
import { Container, Grid } from '@mui/material';

const useStyles = makeStyles(() => ({
  root: {
    flexGrow: 1,
    display: 'flex',
    flexDirection: 'column',
    minHeight: '100vh',
    paddingTop: '20px',
    paddingBottom: '20px',
  },
  content: {
    flex: 1,
  },
}));

interface Props {
  children: ReactNode;
}

export default function Body({ children }: Props) {
  const classes = useStyles();

  return (
    <Container maxWidth="lg" className={classes.root}>
      <Grid item className={classes.content} paddingTop={10}>
        {children}
      </Grid>
    </Container>
  );
}
