import { useEffect, useState } from 'react';
import { League, Player } from '@/types';
import { useRouter } from 'next/router';
import {
  Typography,
  CircularProgress,
  Button,
  Stack,
  Grid,
} from '@mui/material';
import CreateLeagueForm from '@/components/league/create-league';
import CreateEntryForm from '@/components/league/create-entry';
import Card from '@mui/material/Card';

const CreateLeaguePage = () => {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [forwards, setForwards] = useState<Player[]>([]);
  const [defensemen, setDefensemen] = useState<Player[]>([]);
  const [goalies, setGoalies] = useState<Player[]>([]);
  const [league, setLeague] = useState<League | null>(null);
  const [entryForms, setEntryForms] = useState<number>(1);

  useEffect(() => {
    const fetchPlayers = async () => {
      try {
        const response = await fetch('http://localhost:8080/players'); // Replace with your API endpoint
        const data = await response.json();

        const forwardPlayers = data.filter(
          (player: Player) => player.position === 'F',
        );
        const defensemenPlayers = data.filter(
          (player: Player) => player.position === 'D',
        );
        const goaliePlayers = data.filter(
          (player: Player) => player.position === 'G',
        );

        setForwards(forwardPlayers);
        setDefensemen(defensemenPlayers);
        setGoalies(goaliePlayers);
      } catch (error) {
        console.error('Failed to fetch players:', error);
      }
    };

    fetchPlayers();
  }, []);

  const handleSaveLeague = async (newLeague: League) => {
    setIsLoading(true);
    try {
      // Here you would call your API to create the league
      // await yourApi.createLeague(newLeague);

      // Set the league variable with the created league
      setLeague(newLeague);
    } catch (error) {
      console.error('Failed to create league:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleAddEntry = () => {
    setEntryForms((prevCount) => prevCount + 1);
  };

  const handleSubmitLeague = async () => {
    if (!league) {
      return;
    }

    if (!league.entries || league.entries.length < 3) {
      alert('Please save at least 3 entries.');
      return;
    }

    try {
      setIsLoading(true);
      // Here you would call your API to submit the league
      // await yourApi.submitLeague(league);

      // Redirect to /league after successfully submitting the league
      router.push('/league');
    } catch (error) {
      console.error('Failed to submit league:', error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <>
      <Stack spacing={3}>
        <Grid container padding={2} sx={{ display: league ? 'none' : 'block' }}>
          <Typography variant="h4">Create a League</Typography>
          <CreateLeagueForm onCreateLeague={handleSaveLeague} />
          {isLoading && <CircularProgress />}
        </Grid>
        <Stack>
          <Typography variant="h4">Add Entries</Typography>
          <Grid
            container
            spacing={2}
            visibility={league ? 'visible' : 'hidden'}
          >
            {[...Array(entryForms)].map((_, index) => (
              <Grid item key={index} xs={12} sm={6} lg={4}>
                <Card sx={{ padding: 3 }}>
                  <CreateEntryForm
                    onCreateEntry={handleAddEntry}
                    forwards={forwards}
                    defenders={defensemen}
                    goalies={goalies}
                    numForwards={league?.num_forwards || 0}
                    numDefenders={league?.num_defenders || 0}
                    numGoalies={league?.num_goalies || 0}
                  />
                </Card>
              </Grid>
            ))}
            <Grid item xs={12}>
              <Button variant="outlined" onClick={handleAddEntry}>
                Add Entry
              </Button>
              <Button
                variant="contained"
                color="primary"
                onClick={handleSubmitLeague}
                disabled={
                  !league || !league.entries || league.entries.length < 3
                }
              >
                Submit League
              </Button>
            </Grid>
          </Grid>
        </Stack>
      </Stack>
    </>
  );
};

export default CreateLeaguePage;
