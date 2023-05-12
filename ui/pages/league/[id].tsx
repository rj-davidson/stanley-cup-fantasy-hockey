import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { League, Entry, Player } from '@/types';
import { Typography, Box, Stack } from '@mui/material';
import EntryDetails from '@/components/league/entry';

interface LeaguePageProps {
  leagueId: string;
}

const LeaguePage: React.FC<LeaguePageProps> = () => {
  const router = useRouter();
  const leagueId = router.query.id as string; // Access the league ID from the query parameters

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
          // Handle error if the league cannot be fetched
          console.error(
            'Failed to fetch league:',
            response.status,
            response.statusText,
          );
        }
      } catch (error) {
        // Handle network or other errors
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

  const sortedEntries = entries.sort((a: Entry, b: Entry) => {
    const totalPointsA = calculateTotalPoints(a);
    const totalPointsB = calculateTotalPoints(b);
    return totalPointsB - totalPointsA;
  });

  const calculateTotalPoints = (entry: Entry) => {
    let totalPoints = 0;
    entry.forwards?.forEach((forward) => {
      totalPoints += getPlayerPoints(forward);
    });
    entry.defenders?.forEach((defender) => {
      totalPoints += getPlayerPoints(defender);
    });
    entry.goalies?.forEach((goalie) => {
      totalPoints += getPlayerPoints(goalie);
    });
    return totalPoints;
  };

  const getPlayerPoints = (player: Player) => {
    if (player.position === 'Goalie') {
      return (
        player.shutouts * 2 + player.wins * 2 + player.goals + player.assists
      );
    } else {
      return player.goals + player.assists;
    }
  };

  return (
    <Stack spacing={2}>
      <Typography variant="h4">{league.name}</Typography>
      <Typography variant="h6">Season: {league.season}</Typography>

      <Box>
        <Typography variant="h6">Entries</Typography>
        {sortedEntries.map((entry) => (
          <EntryDetails key={entry.id} entry={entry} />
        ))}
      </Box>
    </Stack>
  );
};

export default LeaguePage;
