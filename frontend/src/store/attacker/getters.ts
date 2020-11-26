import { GetterTree } from 'vuex';
import { AttackerState, Move } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<AttackerState, RootState> = {
  currentMove: (state: AttackerState): string => {
    switch(state.currentIndex) {
      case 0: return state.moveA;
      case 1: return state.moveB;
      case 2: return state.moveC;
      case 3: return state.moveD;
    }
    return '';
  },
  physicals: (state: AttackerState): Move[] => {
    return state.physicals;
  },
  specials: (state: AttackerState): Move[] => {
    return state.specials;
  },
}
