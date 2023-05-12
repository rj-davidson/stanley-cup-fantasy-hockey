import React from 'react';
import { Player } from '@/types';
import { Typography, Stack, Grid } from '@mui/material';

interface PlayerDetailsProps {
  player: Player;
}

const PlayerDetails: React.FC<PlayerDetailsProps> = ({ player }) => {
  // Calculate points based on position
  let points;
  if (player.position === 'Goalie') {
    points =
      player.shutouts * 2 + player.wins * 2 + player.goals + player.assists;
  } else {
    points = player.goals + player.assists;
  }

  return (
    <Stack spacing={2}>
      <Grid container alignItems="center" justifyContent="space-between">
        <Grid item>
          <Typography variant="body1">{player.name}</Typography>
        </Grid>
        <Grid item>
          <Typography variant="body1">Points: {points}</Typography>
        </Grid>
      </Grid>
      <Typography variant="body2" color="textSecondary">
        {player.team}
      </Typography>
    </Stack>
  );
};

export default PlayerDetails;
