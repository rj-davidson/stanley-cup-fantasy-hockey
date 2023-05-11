export interface Entry {
  id?: number;
  owner_name: string;
  forwards?: number[];
  defenders?: number[];
  goalies?: number[];
}

export interface League {
  id?: number;
  name: string;
  season: number;
  public: boolean;
  num_forwards: number;
  num_defenders: number;
  num_goalies: number;
  entries?: Entry[];
  edit_key: string;
  code: string;
}

export interface Player {
  id: number;
  name: string;
  position: string;
  goals: number;
  assists: number;
  shutouts: number;
  wins: number;
  team: string;
}
