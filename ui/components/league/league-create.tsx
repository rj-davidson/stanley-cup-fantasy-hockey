import { useState } from 'react';
import {
  Stack,
  TextField,
  Checkbox,
  FormControlLabel,
  Select,
  MenuItem,
  FormControl,
  Button,
} from '@mui/material';
import { Entry, League } from '@/types';

interface Props {
  onCreateLeague: (league: League) => void;
}

export default function CreateLeagueForm(props: Props) {
  const [name, setName] = useState('');
  const [season, setSeason] = useState(new Date().getFullYear());
  const [isPublic, setIsPublic] = useState(false);
  const [numForwards, setNumForwards] = useState('');
  const [numDefenders, setNumDefenders] = useState('');
  const [numGoalies, setNumGoalies] = useState('');

  const handleCreateLeague = () => {
    const league: League = {
      name,
      season,
      public: isPublic,
      num_forwards: parseInt(numForwards),
      num_defenders: parseInt(numDefenders),
      num_goalies: parseInt(numGoalies),
    };
    props.onCreateLeague(league);
  };

  const isFormComplete =
    name.trim() !== '' &&
    numForwards.trim() !== '' &&
    numDefenders.trim() !== '' &&
    numGoalies.trim() !== '';

  return (
    <Stack spacing={2}>
      <TextField
        required
        label="Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <FormControl>
        <Select
          value={season}
          onChange={(e) => setSeason(e.target.value as number)}
        >
          <MenuItem value={new Date().getFullYear()}>
            {new Date().getFullYear()}
          </MenuItem>
        </Select>
      </FormControl>
      <FormControlLabel
        control={
          <Checkbox
            checked={isPublic}
            onChange={(e) => setIsPublic(e.target.checked)}
          />
        }
        label="Public"
      />
      <TextField
        required
        label="Number of Forwards"
        value={numForwards}
        onChange={(e) => setNumForwards(e.target.value)}
        type="number"
        inputProps={{ min: 0 }}
      />
      <TextField
        required
        label="Number of Defenders"
        value={numDefenders}
        onChange={(e) => setNumDefenders(e.target.value)}
        type="number"
        inputProps={{ min: 0 }}
      />
      <TextField
        required
        label="Number of Goalies"
        value={numGoalies}
        onChange={(e) => setNumGoalies(e.target.value)}
        type="number"
        inputProps={{ min: 0 }}
      />
      <Button
        variant="contained"
        color="primary"
        onClick={handleCreateLeague}
        disabled={!isFormComplete}
      >
        Confirm League Settings
      </Button>
    </Stack>
  );
}
