import { NextPage } from 'next';
import { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import {
  List,
  ListItem,
  ListItemText,
  Divider,
  Box,
  Typography,
} from '@mui/material';
import CreateLeagueForm from '@/components/league/create-league';
import { League } from '@/types';

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

  const handleCreateLeague = (league: League) => {
    // TODO: Send the create league request to the API
  };

  useEffect(() => {
    // TODO: Fetch the list of public leagues from the API
  }, []);

  return (
    <Box display="flex" alignItems="flex-start" justifyContent="space-between">
      <Box flexGrow={1}>
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
      <Box ml={4} minWidth="300px">
        <Typography variant="h5" gutterBottom>
          Create League
        </Typography>
        <CreateLeagueForm onCreateLeague={handleCreateLeague} />
      </Box>
    </Box>
  );
};

export default League;
