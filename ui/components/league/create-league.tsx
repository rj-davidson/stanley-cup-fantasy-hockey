import { useState } from 'react';
import {
  Stack,
  TextField,
  Checkbox,
  FormControlLabel,
  Select,
  MenuItem,
  InputLabel,
  FormControl,
  Button,
} from '@mui/material';
import { League } from '@/types';

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
  const [editKey, setEditKey] = useState('');
  const [code, setCode] = useState('');

  const handleCreateLeague = () => {
    const league: League = {
      name,
      season,
      public: isPublic,
      num_forwards: parseInt(numForwards),
      num_defenders: parseInt(numDefenders),
      num_goalies: parseInt(numGoalies),
      entries: [],
      edit_key: editKey,
    };
    props.onCreateLeague(league);
  };

  const isFormComplete =
    name.trim() !== '' &&
    numForwards.trim() !== '' &&
    numDefenders.trim() !== '' &&
    numGoalies.trim() !== '' &&
    editKey.trim() !== '';

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
      <TextField
        required
        label="Admin Password (For Editing League)"
        value={editKey}
        onChange={(e) => setEditKey(e.target.value)}
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
