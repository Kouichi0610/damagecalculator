import { MoveInfo } from '../../moves/store/types';

export const MoveCount = 4;

export interface AttackerState {
  current: number;
  moveA: MoveInfo;
  moveB: MoveInfo;
  moveC: MoveInfo;
  moveD: MoveInfo;
}

