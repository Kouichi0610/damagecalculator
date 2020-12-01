import { GetterTree } from 'vuex';
import { AttackerState, Move } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<AttackerState, RootState> = {
  selectedMoves: (state: AttackerState): string[] => {
    let res: string[] = [];
    res.push(state.moveA);
    res.push(state.moveB);
    res.push(state.moveC);
    res.push(state.moveD);
    return res;
  },
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
