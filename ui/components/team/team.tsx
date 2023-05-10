import React from 'react';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';

interface RosterMember {
  NHLTeam: string;
  Position: string;
  Name: string;
  Points: number;
  Eliminated: boolean;
}

export interface Props {
  owner: string;
  roster: RosterMember[];
}

const TeamRoster: React.FC<Props> = ({ owner, roster }) => {
  const totalPoints = roster.reduce((sum, player) => sum + player.Points, 0);

  return (
    <Card>
      <CardHeader title={owner} />
      <Stack spacing={2} sx={{ p: 2 }}>
        {roster.map((player) => (
          <Typography
            key={player.Name}
            variant="body1"
            sx={{ color: player.Eliminated ? 'red' : 'black' }}
          >
            {player.Name} - NHL Team: {player.NHLTeam}, Position:{' '}
            {player.Position}, Points: {player.Points}
          </Typography>
        ))}
        <Typography variant="h6">Total Points: {totalPoints}</Typography>
      </Stack>
    </Card>
  );
};

export default TeamRoster;
