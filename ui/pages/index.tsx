import { Inter } from 'next/font/google';
import {
  Box,
  Button,
  Container,
  Grid,
  List,
  ListItem,
  ListItemText,
  Stack,
  Typography,
} from '@mui/material';
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
      <Stack spacing={5}>
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

        <Container maxWidth="md">
          <Typography variant="h6" sx={{ mt: 2 }}>
            Overview
          </Typography>
          <Typography variant="body1" align="left" sx={{ mt: 2 }}>
            This is a sample application built with Next.js, Go, and PostgreSQL.
            The source code is available on GitHub at{' '}
            <Link href="https://github.com/rj-davidson/stanley-cup-fantasy-hockey">
              github.com/rj-davidson/stanley-cup-fantasy-hockey
            </Link>
            . The application allows users to create leagues with entries for
            the Stanley Cup Playoffs. Entries are a list of players that the
            entry owner thinks will score the most points in the playoffs.
            Points in hockey are goals or assists. Goalie points are calculated
            differently with 2 points for a win and 2 bonus points for a
            shutout. The stats for this application are pulled from the NHL API.
            Stats are fetched every 20 minutes and appear to be updated by the
            NHL on game completion. The application is hosted on a small ubuntu
            server using Docker containers.
          </Typography>
          <Typography variant="h6" sx={{ mt: 2 }}>
            Tech Stack
          </Typography>
          <List>
            <ListItem>
              <ListItemText primary="Next.js" />
              <ListItemText secondary="React" />
              <ListItemText secondary="Material UI" />
              <ListItemText secondary="TypeScript" />
            </ListItem>
            <ListItem>
              <ListItemText primary="Go" />
              <ListItemText secondary="Fiber" />
              <ListItemText secondary="Ent" />
              <ListItemText secondary="Cron" />
            </ListItem>
            <ListItem>
              <ListItemText primary="PostgreSQL" />
            </ListItem>
          </List>
        </Container>
      </Stack>
    </Box>
  );
}

export default HomePage;
