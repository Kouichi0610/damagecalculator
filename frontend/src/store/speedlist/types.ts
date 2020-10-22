
export interface SpeedInfo {
  info: string;     // 同速のポケモン
  speed: number;  // 素早さ(実数)
}

export interface SpeedListState {
  list: SpeedInfo[];
}