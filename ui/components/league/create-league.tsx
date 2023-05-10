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
      is_public: isPublic,
      num_forwards: parseInt(numForwards),
      num_defenders: parseInt(numDefenders),
      num_goalies: parseInt(numGoalies),
      edit_key: editKey,
      code,
    };
    props.onCreateLeague(league);
  };

  return (
    <Stack spacing={2}>
      <TextField
        required
        label="Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <FormControl>
        <InputLabel>Season</InputLabel>
        <Select
          value={season}
          onChange={(e) => setSeason(e.target.value as number)}
        >
          <MenuItem value={new Date().getFullYear()}>
            {new Date().getFullYear()}
          </MenuItem>
          <MenuItem value={new Date().getFullYear() + 1}>
            {new Date().getFullYear() + 1}
          </MenuItem>
          <MenuItem value={new Date().getFullYear() + 2}>
            {new Date().getFullYear() + 2}
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
      <TextField
        required
        label="League Code (For Joining League)"
        value={code}
        onChange={(e) => setCode(e.target.value)}
      />
      <Button variant="contained" color="primary" onClick={handleCreateLeague}>
        Create League
      </Button>
    </Stack>
  );
}
