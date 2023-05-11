import { NextPage } from 'next';
import { useState, useEffect } from 'react';
import Link from 'next/link';
import {
  Button,
  List,
  ListItem,
  ListItemText,
  Divider,
  Box,
  Typography,
  Grid,
} from '@mui/material';
import { League } from '@/types';
import { useRouter } from 'next/router';

interface LeagueListItemProps {
  league: League;
  onClick: () => void;
}

const LeagueListItem = (props: LeagueListItemProps) => {
  const { league, onClick } = props;

  return (
    <>
      <ListItem button onClick={onClick}>
        <ListItemText primary={league.name} />
      </ListItem>
      <Divider />
    </>
  );
};

const League: NextPage = () => {
  const [leagues, setLeagues] = useState<League[]>([]);
  const router = useRouter();

  useEffect(() => {
    const fetchLeagues = async () => {
      try {
        const response = await fetch('http://localhost:8080/leagues');
        if (response.ok) {
          const leagues = await response.json();
          const publicLeagues = leagues.filter(
            (league: League) => league.public,
          );
          setLeagues(publicLeagues);
        } else {
          console.error('Failed to fetch leagues:', response.statusText);
        }
      } catch (error) {
        console.error('Failed to fetch leagues:', error);
      }
    };
    fetchLeagues();
  }, []);

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
          <Button
            variant="contained"
            color="primary"
            onClick={() =>
              window.scrollTo({
                top: document.body.scrollHeight,
                behavior: 'smooth',
              })
            }
          >
            Join a League
          </Button>
        </Grid>
        <Grid item>
          <Link href="/league/create" passHref>
            <Button variant="contained" color="primary">
              Create a League
            </Button>
          </Link>
        </Grid>
      </Grid>

      <Box mt={4} width="100%">
        <Typography variant="h5" gutterBottom>
          Public Leagues
        </Typography>
        <List>
          {leagues.map((league) => (
            <LeagueListItem
              key={league.id}
              league={league}
              onClick={() => router.push(`/league/${league.id}`)}
            />
          ))}
        </List>
      </Box>
    </Box>
  );
};

export default League;
