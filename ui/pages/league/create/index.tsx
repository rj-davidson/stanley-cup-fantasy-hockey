import { useEffect, useState } from 'react';
import { Entry, League, Player, CreateLeagueBundle } from '@/types';
import { useRouter } from 'next/router';
import {
  Typography,
  CircularProgress,
  Button,
  Stack,
  Grid,
  Card,
} from '@mui/material';
import CreateLeagueForm from '@/components/league/league-create';
import CreateEntryForm from '@/components/league/entry-create';

const CreateLeaguePage = () => {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [players, setPlayers] = useState<Player[]>([]);
  const [entries, setEntries] = useState<Entry[]>([]);
  const [league, setLeague] = useState<League>();
  const [leagueBundle, setLeagueBundle] = useState<CreateLeagueBundle>();
  const [entryForms, setEntryForms] = useState<number>(1);
  const [submittedIndices, setSubmittedIndices] = useState<number[]>([]);

  useEffect(() => {
    const fetchPlayers = async () => {
      try {
        const response = await fetch(
          `${process.env.NEXT_PUBLIC_API_URL}/players`,
        );
        const data = await response.json();
        setPlayers(data);
      } catch (error) {
        console.error('Failed to fetch players:', error);
      }
    };

    fetchPlayers();
  }, []);

  const handleSaveLeague = async (newLeague: League) => {
    setIsLoading(true);
    try {
      setLeague(newLeague);
    } catch (error) {
      console.error('Failed to save league:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleFinalizeEntry = (newEntry: Entry, index: number) => {
    setEntries((prevEntries) => [...prevEntries, newEntry]);
    setSubmittedIndices((prevIndices) => [...prevIndices, index]);
  };

  const handleAddEntryForm = () => {
    setEntryForms((prevCount) => prevCount + 1);
  };

  const handleSubmitLeague = async () => {
    if (!league || !entries) {
      return;
    }

    if (entries.length < 3) {
      alert('Please save at least 3 entries.');
      return;
    }

    setIsLoading(true);

    // Create the leagueBundle immediately before it's needed
    const newLeagueBundle = {
      league,
      entries,
    };

    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/leagues`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(newLeagueBundle), // Use the new leagueBundle
        },
      );
      setLeagueBundle(await response.json());

      if (response.ok) {
        await router.push(
          leagueBundle ? '/league/' + leagueBundle.league.id : '/league',
        );
        setIsLoading(false);
      } else {
        console.error('Failed to submit league:', response.status);
      }
    } catch (error) {
      console.error('Failed to submit league:', error);
    }
  };

  return (
    <Stack spacing={3}>
      <Grid container padding={2} sx={{ display: league ? 'none' : 'block' }}>
        <Typography variant="h4">Create a League</Typography>
        <CreateLeagueForm onCreateLeague={handleSaveLeague} />
        {isLoading && <CircularProgress />}
      </Grid>

      <Typography variant="h4">Add Entries</Typography>
      <Grid container spacing={2} visibility={league ? 'visible' : 'hidden'}>
        {entries.map((entry, index) => (
          <Grid item key={index} xs={12} sm={6} lg={4}>
            <Card sx={{ padding: 3 }}>
              <Typography variant="h6">{entry.owner_name}</Typography>
            </Card>
          </Grid>
        ))}

        {[...Array(entryForms)].reverse().map(
          (_, index) =>
            !submittedIndices.includes(index) && (
              <Grid item key={index} xs={12} sm={6} lg={4}>
                <Card sx={{ padding: 3 }}>
                  <CreateEntryForm
                    onCreateEntry={(entry) => handleFinalizeEntry(entry, index)}
                    players={players}
                    numForwards={league?.num_forwards || 0}
                    numDefenders={league?.num_defenders || 0}
                    numGoalies={league?.num_goalies || 0}
                  />
                </Card>
              </Grid>
            ),
        )}

        <Grid item xs={12}>
          <Button variant="outlined" onClick={handleAddEntryForm}>
            Add Entry
          </Button>

          <Button
            variant="contained"
            color="primary"
            onClick={handleSubmitLeague}
            disabled={!league || !entries || entries.length < 3 || isLoading}
          >
            Submit League
          </Button>
        </Grid>
      </Grid>
    </Stack>
  );
};

export default CreateLeaguePage;
