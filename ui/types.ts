export interface League {
  id?: number;
  name: string;
  season: number;
  is_public: boolean;
  num_forwards: number;
  num_defenders: number;
  num_goalies: number;
  edit_key: string;
  code: string;
}
