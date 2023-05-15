import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { League, Entry, Player, LeagueBundle } from '@/types';
import CircularProgress from '@mui/material/CircularProgress';
import { Typography, Box, Stack, Grid, Container } from '@mui/material';
import EntryDetails from '@/components/league/entry-details';
import Card from '@mui/material/Card';

interface LeaguePageProps {
  leagueId: string;
}

const LeaguePage: React.FC<LeaguePageProps> = () => {
  const router = useRouter();
  const leagueId = router.query.id as string;

  const [leagueBundle, setLeagueBundle] = useState<LeagueBundle | null>(null);
  const [loading, setLoading] = useState(true); // Add a loading state

  useEffect(() => {
    const fetchLeague = async () => {
      try {
        setLoading(true);
        const response = await fetch(
          `${process.env.NEXT_PUBLIC_API_URL}/leagues/${leagueId}`,
        );
        if (response.ok) {
          const resp = await response.json();
          setLeagueBundle(resp);
        } else {
          console.error(
            'Failed to fetch league:',
            response.status,
            response.statusText,
          );
        }
      } catch (error) {
        console.error('Error fetching league:', error);
      } finally {
        setLoading(false); // Stop loading when fetch is complete
      }
    };

    fetchLeague().then((r) => r);
  }, [leagueId]);

  if (loading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        height="100vh"
      >
        <CircularProgress />
      </Box>
    );
  }

  if (!leagueBundle?.league && !loading) {
    return (
      <Stack spacing={2}>
        <Typography variant="h6">League not found.</Typography>
      </Stack>
    );
  }
  if (!leagueBundle?.league && !loading) {
    return (
      <Stack spacing={2}>
        <Typography variant="h6">League not found.</Typography>
      </Stack>
    );
  }

  if (!leagueBundle) {
    return null; // or replace with your preferred fallback UI
  }

  const { entries, league, teams, players } = leagueBundle;

  if (!entries || entries.length === 0) {
    return (
      <Stack spacing={2}>
        <Typography variant="h4">{league.name}</Typography>
        <Typography variant="h6">Season: {league.season}</Typography>
        <Box>
          <Typography variant="h6">No entries found.</Typography>
        </Box>
      </Stack>
    );
  }

  const getPlayerPoints = (playerID: number) => {
    const player = players.find((player) => player.id === playerID);

    if (!player) {
      return 0;
    }

    if (player.position === 'Goalie') {
      return (
        player.shutouts * 2 + player.wins * 2 + player.goals + player.assists
      );
    } else {
      return player.goals + player.assists;
    }
  };

  const calculateTotalPoints = (entry: Entry) => {
    let totalPoints = 0;
    entry.playerIDs?.forEach((playerID) => {
      totalPoints += getPlayerPoints(playerID);
    });
    return totalPoints;
  };

  const sortedEntries = entries.sort((a: Entry, b: Entry) => {
    const totalPointsA = calculateTotalPoints(a);
    const totalPointsB = calculateTotalPoints(b);
    return totalPointsB - totalPointsA;
  });

  return (
    <Stack spacing={4}>
      <Stack>
        <Typography variant="h4">{league.name}</Typography>
        <Stack direction="row" spacing={2}>
          <Typography variant="subtitle1">{`Season: ${league.season}`}</Typography>
          <Typography variant="subtitle1">{`Entries: ${entries.length}`}</Typography>
          <Typography variant="subtitle1">{`Teams: ${teams.length}`}</Typography>
          <Typography variant="subtitle1">{`Players: ${players.length}`}</Typography>
        </Stack>
        <Typography variant="h5">Standings</Typography>
      </Stack>
      <Container>
        <Grid container spacing={3} paddingLeft={-3}>
          {sortedEntries.map((entry, index) => (
            <Grid item xs={12} sm={6} md={4} lg={3} key={entry.id}>
              <Card sx={{ p: 2 }} variant="outlined">
                <Stack spacing={2} padding={1}>
                  <EntryDetails entry={entry} teams={teams} players={players} />
                  <Typography variant="subtitle2" align="center">
                    {`${index + 1}`}/{entries.length}
                  </Typography>
                </Stack>
              </Card>
            </Grid>
          ))}
        </Grid>
      </Container>
    </Stack>
  );
};

export default LeaguePage;
