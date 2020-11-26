import { GetterTree } from 'vuex';
import { AttackerState, Move } from './types';
import { RootState } from '@/store/types'
import { Result } from './sendDamage'

export const getters: GetterTree<AttackerState, RootState> = {
  move: (state: AttackerState): string => {
    return state.selectedMove[state.currentIndex];
  },
  physicals: (state: AttackerState): Move[] => {
    return state.physicals;
  },
  specials: (state: AttackerState): Move[] => {
    return state.specials;
  },
  results: (state: AttackerState): Result[] => {
    return state.result;
  }
}
