import React from 'react';
import { Player, Team } from '@/types';
import { Typography, Stack, Grid } from '@mui/material';

interface PlayerDetailsProps {
  player: Player;
  teams: Team[];
}

const PlayerDetails: React.FC<PlayerDetailsProps> = ({ player, teams }) => {
  // Calculate points based on position
  let points;
  if (player.position === 'Goalie') {
    points =
      player.shutouts * 2 + player.wins * 2 + player.goals + player.assists;
  } else {
    points = player.goals + player.assists;
  }

  const team = teams.find((team) => team.id === player.team_id);
  // Log in console if team is not found
  if (!team) {
    console.error('Team not found for player:', player);
  }

  const playerNameStyle = {
    color: team?.eliminated ? 'grey' : 'black',
    textDecoration: team?.eliminated ? 'line-through' : 'none',
  };

  return (
    <Stack>
      <Grid container alignItems="center" justifyContent="space-between">
        <Grid item>
          <Typography variant="body1" style={playerNameStyle}>
            {player.name}
          </Typography>
        </Grid>
        <Grid item>
          <Typography variant="body1" style={playerNameStyle}>
            {points}
          </Typography>
        </Grid>
      </Grid>
      <Typography variant="body2" color="textSecondary">
        {`${team?.name}, ${player.position}`}
      </Typography>
    </Stack>
  );
};

export default PlayerDetails;
