import { MutationTree } from 'vuex';
import { RootState } from '@/store/types'
import { TargetSelectState } from './types';

export const mutations: MutationTree<TargetSelectState> = {
  setInitialTotal(state, payload: number) {
    console.log('setInitialTotal ' + payload);
    state.initialTotal = payload;
  },
  setInitialType(state, payload: string) {
    console.log('setInitialType ' + payload);
    state.initialType = payload;
  },
}
