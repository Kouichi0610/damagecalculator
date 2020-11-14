import { DefenderDamages, DefendersResult } from '../defenderDamages'

export const MoveCount = 4;

export interface AttackerState {
  target: string;
  current: number;
  results: DefendersResult[];
}

