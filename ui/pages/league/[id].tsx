import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { League, Entry, Player } from '@/types';
import { Typography, Box, Stack } from '@mui/material';
import EntryDetails from '@/components/league/entry';
import Card from '@mui/material/Card';

interface LeaguePageProps {
  leagueId: string;
}

const LeaguePage: React.FC<LeaguePageProps> = () => {
  const router = useRouter();
  const leagueId = router.query.id as string;

  const [league, setLeague] = useState<League | null>(null);

  useEffect(() => {
    const fetchLeague = async () => {
      try {
        const response = await fetch(
          `http://localhost:8080/leagues/${leagueId}`,
        );
        if (response.ok) {
          const leagueData = await response.json();
          setLeague(leagueData);
        } else {
          console.error(
            'Failed to fetch league:',
            response.status,
            response.statusText,
          );
        }
      } catch (error) {
        console.error('Error fetching league:', error);
      }
    };

    fetchLeague();
  }, [leagueId]);

  if (!league) {
    return (
      <Stack spacing={2}>
        <Typography variant="h6">League not found.</Typography>
      </Stack>
    );
  }

  const { entries } = league;

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

  const getPlayerPoints = (player: Player) => {
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
    entry.players?.forEach((player) => {
      totalPoints += getPlayerPoints(player);
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
      <Typography variant="h4">{league.name}</Typography>
      <Typography variant="h6">Season: {league.season}</Typography>
      <Stack spacing={2}>
        <Typography variant="h6">Leaderboard</Typography>
        <Stack direction="row" spacing={2} sx={{ flexWrap: 'wrap' }}>
          {sortedEntries.map((entry, index) => (
            <Card
              key={entry.id}
              sx={{ p: 2, flexGrow: 1, minWidth: 250, position: 'relative' }}
            >
              <Stack spacing={2} padding={1}>
                <EntryDetails entry={entry} />
                <Typography variant="subtitle2" align="center">
                  {`${index + 1}`}/{entries.length}
                </Typography>
              </Stack>
            </Card>
          ))}
        </Stack>
      </Stack>
    </Stack>
  );
};

export default LeaguePage;
