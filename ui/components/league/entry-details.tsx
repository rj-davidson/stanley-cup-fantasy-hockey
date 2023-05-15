import React from 'react';
import { Entry, Player, Team } from '@/types';
import { Box, Stack, Typography } from '@mui/material';
import PlayerDetails from '@/components/league/player-details';

interface EntryDetailsProps {
  entry: Entry;
  teams: Team[];
  players: Player[];
}

const EntryDetails: React.FC<EntryDetailsProps> = ({
  entry,
  teams,
  players,
}) => {
  const { owner_name, playerIDs } = entry;

  const calculateTotalPoints = (players: Player[]) => {
    return players.reduce((total, player) => {
      return total + getPlayerPoints(player);
    }, 0);
  };

  const getPlayerById = (id: number): Player | undefined => {
    return players.find((player) => player.id === id);
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
  const sortedPlayers = playerIDs
    ? (playerIDs
        .map((id) => getPlayerById(id))
        .filter((player) => player !== undefined) as Player[])
    : [];

  // Calculate total points for each position
  const points = calculateTotalPoints(sortedPlayers);

  return (
    <Stack spacing={1}>
      <Stack direction={'row'} justifyContent={'space-between'}>
        <Typography variant="h5">{owner_name}</Typography>
        <Typography variant="h6">{points}</Typography>
      </Stack>

      {sortedPlayers.length > 0 && (
        <Box>
          <Stack spacing={1}>
            {sortedPlayers.map((player) => (
              <PlayerDetails key={player.id} player={player} teams={teams} />
            ))}
          </Stack>
        </Box>
      )}
    </Stack>
  );
};

export default EntryDetails;
