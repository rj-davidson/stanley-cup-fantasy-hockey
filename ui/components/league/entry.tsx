import React from 'react';
import { Entry, Player } from '@/types';
import { Typography, Box, Stack } from '@mui/material';
import PlayerDetails from '@/components/league/player';

interface EntryDetailsProps {
  entry: Entry;
}

const EntryDetails: React.FC<EntryDetailsProps> = ({ entry }) => {
  const { owner_name, forwards, defenders, goalies } = entry;

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

  // Calculate total points for each position
  const totalForwardPoints = forwards ? calculateTotalPoints(forwards) : 0;
  const totalDefenderPoints = defenders ? calculateTotalPoints(defenders) : 0;
  const totalGoaliePoints = goalies ? calculateTotalPoints(goalies) : 0;

  return (
    <Stack spacing={2}>
      <Typography variant="h6">{owner_name}</Typography>
      <Typography variant="body1">
        Total Points:{' '}
        {totalForwardPoints + totalDefenderPoints + totalGoaliePoints}
      </Typography>

      {forwards && (
        <Box>
          <Typography variant="subtitle1">
            Forwards (Total Points: {totalForwardPoints})
          </Typography>
          {forwards.map((forward) => (
            <PlayerDetails key={forward.id} player={forward} />
          ))}
        </Box>
      )}

      {defenders && (
        <Box>
          <Typography variant="subtitle1">
            Defenders (Total Points: {totalDefenderPoints})
          </Typography>
          {defenders.map((defender) => (
            <PlayerDetails key={defender.id} player={defender} />
          ))}
        </Box>
      )}

      {goalies && (
        <Box>
          <Typography variant="subtitle1">
            Goalies (Total Points: {totalGoaliePoints})
          </Typography>
          {goalies.map((goalie) => (
            <PlayerDetails key={goalie.id} player={goalie} />
          ))}
        </Box>
      )}
    </Stack>
  );
};

export default EntryDetails;
