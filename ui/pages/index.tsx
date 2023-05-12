import { Inter } from 'next/font/google';
import { Box, Button, Grid, Typography } from '@mui/material';
import Link from 'next/link';

const inter = Inter({ subsets: ['latin'] });

function HomePage() {
  return (
    <Box
      display="flex"
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
    >
      <Typography variant="h3" align="center" gutterBottom>
        Welcome to Fantasy Hockey
      </Typography>

      <Grid container spacing={2} justifyContent="center">
        <Grid item>
          <Link href="/league" passHref>
            <Button variant="contained" color="primary">
              View Leagues
            </Button>
          </Link>
        </Grid>
        <Grid item>
          <Link href="/league/create" passHref>
            <Button variant="contained" color="primary">
              Create a League
            </Button>
          </Link>
        </Grid>
      </Grid>
    </Box>
  );
}

export default HomePage;
