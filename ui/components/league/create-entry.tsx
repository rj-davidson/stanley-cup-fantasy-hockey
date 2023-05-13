import { useState, useEffect } from 'react';
import { Stack, TextField, Button, Autocomplete } from '@mui/material';
import { Entry, Player } from '@/types';

interface Props {
  onCreateEntry: (entry: Entry) => void;
  players: Player[];
  numForwards: number;
  numDefenders: number;
  numGoalies: number;
}

export default function CreateEntryForm(props: Props) {
  const [ownerName, setOwnerName] = useState('');
  const [selectedPlayers, setSelectedPlayers] = useState<Player[]>([]);

  useEffect(() => {
    setSelectedPlayers(
      new Array(props.numForwards + props.numDefenders + props.numGoalies).fill(
        null,
      ),
    );
  }, [props.numForwards, props.numDefenders, props.numGoalies]);

  const allPlayersSelected = () => {
    return selectedPlayers.every((player) => player !== null);
  };

  const handleFinalizeEntry = () => {
    if (!allPlayersSelected()) {
      alert('Please select all players before saving the entry.');
      return;
    }
    const entry: Entry = {
      owner_name: ownerName,
      players: selectedPlayers.map((p) => p),
    };
    props.onCreateEntry(entry);
  };

  const handlePlayerChange =
    (index: number) => (event: any, value: Player | null) => {
      setSelectedPlayers((oldArray: Player[]) => {
        const newArray = [...oldArray];
        newArray[index] = value || oldArray[index];
        return newArray;
      });
    };

  const playerSelect = (
    label: string,
    players: Player[],
    selectedIndex: number,
  ) => {
    let filteredPlayers: Player[] = [];

    switch (label) {
      case 'Forward':
        filteredPlayers = players.filter(
          (player) => player.position === 'Forward',
        );
        break;
      case 'Defenseman':
        filteredPlayers = players.filter(
          (player) => player.position === 'Defenseman',
        );
        break;
      case 'Goalie':
        filteredPlayers = players.filter(
          (player) => player.position === 'Goalie',
        );
        break;
      default:
        break;
    }

    return (
      <Autocomplete
        key={selectedIndex}
        id={`player-select-${selectedIndex}`}
        options={filteredPlayers.filter(
          (p: Player) => !selectedPlayers.includes(p),
        )}
        getOptionLabel={(option) => option.name}
        style={{ width: 300 }}
        renderInput={(params) => <TextField {...params} label={label} />}
        value={selectedPlayers[selectedIndex]}
        onChange={handlePlayerChange(selectedIndex)}
      />
    );
  };

  return (
    <Stack spacing={2}>
      <TextField
        required
        label="Owner Name"
        value={ownerName}
        onChange={(e) => setOwnerName(e.target.value)}
      />
      {Array.from({ length: props.numForwards }).map((_, index) =>
        playerSelect('Forward', props.players, index),
      )}
      {Array.from({ length: props.numDefenders }).map((_, index) =>
        playerSelect('Defenseman', props.players, index + props.numForwards),
      )}
      {Array.from({ length: props.numGoalies }).map((_, index) =>
        playerSelect(
          'Goalie',
          props.players,
          index + props.numForwards + props.numDefenders,
        ),
      )}
      <Button
        variant="contained"
        color="primary"
        onClick={handleFinalizeEntry}
        disabled={!allPlayersSelected()}
      >
        Finalize Entry
      </Button>
    </Stack>
  );
}
