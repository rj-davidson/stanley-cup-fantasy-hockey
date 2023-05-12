import { useEffect, useState } from 'react';
import { Entry, League, Player } from '@/types';
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
          (player: Player) => player.position === 'Forward',
        );
        const defensemenPlayers = data.filter(
          (player: Player) => player.position === 'Defenseman',
        );
        const goaliePlayers = data.filter(
          (player: Player) => player.position === 'Goalie',
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
      // ensure entries is initialized as an empty array
      setLeague({ ...newLeague, entries: [] });
    } catch (error) {
      console.error('Failed to create league:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleFinalizeEntry = (newEntry: Entry) => {
    if (league) {
      // @ts-ignore
      setLeague((currentLeague) => {
        if (currentLeague) {
          return {
            ...currentLeague,
            entries: [...currentLeague.entries, newEntry],
          };
        }
      });
      setEntryForms((prevCount) => prevCount - 1);
    }
  };

  const handleAddEntryForm = () => {
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

      const response = await fetch('http://localhost:8080/leagues', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(league),
      });

      if (response.ok) {
        // Redirect to /league after successfully submitting the league
        router.push('/league');
      } else {
        console.error('Failed to submit league:', response.status);
      }
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
          <Grid
            container
            spacing={2}
            visibility={league ? 'visible' : 'hidden'}
          >
            <Typography variant="h4">Add Entries</Typography>
            {[...Array(entryForms)].map((_, index) => (
              <Grid item key={index} xs={12} sm={6} lg={4}>
                <Card sx={{ padding: 3 }}>
                  <CreateEntryForm
                    onCreateEntry={handleFinalizeEntry}
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

            {league?.entries.map((entry, index) => (
              <Grid item key={index} xs={12} sm={6} lg={4}>
                <Card sx={{ padding: 3 }}>
                  <Typography variant="h6">{entry.owner_name}</Typography>
                </Card>
              </Grid>
            ))}

            <Grid item xs={12}>
              <Button variant="outlined" onClick={handleAddEntryForm}>
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
