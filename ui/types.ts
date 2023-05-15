export interface Entry {
  id?: number;
  owner_name: string;
  playerIDs?: number[];
}

export interface League {
  id?: number;
  name: string;
  season: number;
  public: boolean;
  num_forwards: number;
  num_defenders: number;
  num_goalies: number;
}

export interface Player {
  id: number;
  name: string;
  position: string;
  goals: number;
  assists: number;
  shutouts: number;
  wins: number;
  team_id: number;
}

export interface LeagueBundle {
  league: League;
  players: Player[];
  entries: Entry[];
  teams: Team[];
}

export interface CreateLeagueBundle {
  league: League;
  entries: Entry[];
}

export interface Team {
  id: number;
  name: string;
  eliminated: boolean;
}
