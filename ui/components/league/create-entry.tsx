import { useState, useEffect } from 'react';
import { Stack, TextField, Button, Autocomplete } from '@mui/material';
import { Entry, Player } from '@/types';

interface Props {
  onCreateEntry: (entry: Entry) => void;
  forwards: Player[];
  defenders: Player[];
  goalies: Player[];
  numForwards: number;
  numDefenders: number;
  numGoalies: number;
}

export default function CreateEntryForm(props: Props) {
  const [ownerName, setOwnerName] = useState('');
  const [selectedForwards, setSelectedForwards] = useState<Player[]>([]);
  const [selectedDefenders, setSelectedDefenders] = useState<Player[]>([]);
  const [selectedGoalies, setSelectedGoalies] = useState<Player[]>([]);

  useEffect(() => {
    setSelectedForwards(new Array(props.numForwards).fill(null));
    setSelectedDefenders(new Array(props.numDefenders).fill(null));
    setSelectedGoalies(new Array(props.numGoalies).fill(null));
  }, [props.numForwards, props.numDefenders, props.numGoalies]);

  const allPlayersSelected = () => {
    const allPlayers = [
      ...selectedForwards,
      ...selectedDefenders,
      ...selectedGoalies,
    ];
    return allPlayers.every((player) => player !== null);
  };

  const handleCreateEntry = () => {
    if (!allPlayersSelected()) {
      alert('Please select all players before saving the entry.');
      return;
    }

    // handle entry creation
  };

  const handlePlayerChange =
    (setSelectedPlayers: Function, index: number) =>
    (event: any, value: Player | null) => {
      setSelectedPlayers((oldArray: Player[]) => {
        const newArray = [...oldArray];
        newArray[index] = value || oldArray[index];
        return newArray;
      });
    };

  const playerSelect = (
    label: string,
    players: Player[],
    selectedPlayers: Player[],
    setSelectedPlayers: Function,
  ) =>
    selectedPlayers.map((player: Player, index: number) => (
      <Autocomplete
        key={index}
        id={`player-select-${index}`}
        options={players.filter((p: Player) => !selectedPlayers.includes(p))}
        getOptionLabel={(option) => option.name}
        style={{ width: 300 }}
        renderInput={(params) => <TextField {...params} label={label} />}
        value={player}
        onChange={handlePlayerChange(setSelectedPlayers, index)}
      />
    ));

  return (
    <Stack spacing={2}>
      <TextField
        required
        label="Owner Name"
        value={ownerName}
        onChange={(e) => setOwnerName(e.target.value)}
      />
      {playerSelect(
        'Forwards',
        props.forwards,
        selectedForwards,
        setSelectedForwards,
      )}
      {playerSelect(
        'Defenders',
        props.defenders,
        selectedDefenders,
        setSelectedDefenders,
      )}
      {playerSelect(
        'Goalies',
        props.goalies,
        selectedGoalies,
        setSelectedGoalies,
      )}
      <Button
        variant="contained"
        color="primary"
        onClick={handleCreateEntry}
        disabled={!allPlayersSelected()}
      >
        Finalize Entry
      </Button>
    </Stack>
  );
}
