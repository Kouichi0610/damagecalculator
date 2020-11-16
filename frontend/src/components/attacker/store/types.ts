import { DefendersResult } from '../../target/store/defenderDamages'

export const MoveCount = 4;

export interface AttackerState {
  target: string;
  current: number;
  results: DefendersResult[];
}

