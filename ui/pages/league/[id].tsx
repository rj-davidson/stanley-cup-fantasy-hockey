import React from 'react';
import { League, Entry, Player } from '@/types';
import { Typography, Box, Stack } from '@mui/material';
import EntryDetails from '@/components/league/entry';

interface LeaguePageProps {
  league: League;
}

const LeaguePage: React.FC<LeaguePageProps> = ({ league }) => {
  const { entries } = league;

  // Sort entries by total points in descending order
  const sortedEntries = entries.sort((a: Entry, b: Entry) => {
    const totalPointsA = calculateTotalPoints(a);
    const totalPointsB = calculateTotalPoints(b);
    return totalPointsB - totalPointsA;
  });

  // Calculate total points for an entry
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

  // Calculate points based on position
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
