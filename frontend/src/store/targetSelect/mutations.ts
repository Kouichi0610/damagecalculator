import { MutationTree } from 'vuex';
import { RootState } from '@/store/types'
import { TargetSelectState, Species } from './types';

export const mutations: MutationTree<TargetSelectState> = {
  setInitialTotal(state, payload: number) {
    state.initialTotal = payload;
  },
  setInitialType(state, payload: string) {
    state.initialType = payload;
  },
  setCandidates(state, payload: Species[]) {
    state.candidates = payload;
  }
}
