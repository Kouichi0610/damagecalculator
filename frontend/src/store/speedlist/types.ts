
export interface SpeedInfo {
  info: string;     // 同速のポケモン一覧
  speed: number;  // 素早さ(実数)
}

export interface SpeedListState {
  list: SpeedInfo[];
}