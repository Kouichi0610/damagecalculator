import { GetterTree } from 'vuex';
import { TargetSelectState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<TargetSelectState, RootState> = {
  initialTotal: (state: TargetSelectState) => {
    return state.initialTotal;
  },
  initialType: (state: TargetSelectState) => {
    return state.initialType;
  }
}

