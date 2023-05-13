import React from 'react';
import { Entry, Player } from '@/types';
import { Typography, Box, Stack } from '@mui/material';
import PlayerDetails from '@/components/league/player';

interface EntryDetailsProps {
  entry: Entry;
}

const EntryDetails: React.FC<EntryDetailsProps> = ({ entry }) => {
  const { owner_name, players } = entry;

  // Calculate total points for each position
  const calculateTotalPoints = (players: Player[]) => {
    return players.reduce((total, player) => {
      return total + getPlayerPoints(player);
    }, 0);
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

  // Sort players by points in descending order
  const sortedPlayers = players
    ? [...players].sort((a, b) => {
        const pointsA = getPlayerPoints(a);
        const pointsB = getPlayerPoints(b);

        if (pointsA !== pointsB) {
          return pointsB - pointsA; // Sort by points in descending order
        }

        if (a.position !== b.position) {
          if (a.position === 'Forward') return -1;
          if (a.position === 'Defenseman')
            return a.position === b.position ? -1 : 1;
          if (a.position === 'Goalie') return 1;
        }

        return a.name.localeCompare(b.name);
      })
    : [];

  // Calculate total points for each position
  const points = calculateTotalPoints(sortedPlayers);

  return (
    <Stack spacing={2}>
      <Stack direction={'row'} justifyContent={'space-between'}>
        <Typography variant="h5">{owner_name}</Typography>
        <Typography variant="h6">{points}</Typography>
      </Stack>

      {sortedPlayers.length > 0 && (
        <Box>
          <Stack spacing={1}>
            {sortedPlayers.map((player) => (
              <PlayerDetails key={player.id} player={player} />
            ))}
          </Stack>
        </Box>
      )}
    </Stack>
  );
};

export default EntryDetails;
