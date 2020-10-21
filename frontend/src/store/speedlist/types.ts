
export interface SpeedInfo {
  info: string;     // 同速のポケモン
  species: number;  // 素早さ種族値
}

export interface SpeedListState {
  list: SpeedInfo[];
}