import { GetterTree } from 'vuex';
import { AttackerState } from './types';
import { RootState } from '@/store/types'
import { MoveInfo } from '../../moves/store/types'

export const getters: GetterTree<AttackerState, RootState> = {
  currentMove: (state: AttackerState): MoveInfo => {
    switch(state.current) {
      case 0:
        return state.moveA;
      case 1:
        return state.moveB;
      case 2:
        return state.moveC;
      case 3:
        return state.moveD;
      }
      return state.moveA;
    },
  moves: (state: AttackerState): MoveInfo[] => {
    return [state.moveA, state.moveB, state.moveC, state.moveD];
  }
}
