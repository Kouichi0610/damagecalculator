import { GetterTree } from 'vuex';
import { DefenderState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<DefenderState, RootState> = {
  /*
  isInitialized: (state: NatureState): boolean => {
    return state.natures.length > 0;
  },
  natures: (state: NatureState): Nature[] => {
    return state.natures;
  },
  current: (state: NatureState): Nature => {
    return state.current;
  }
  */
}
