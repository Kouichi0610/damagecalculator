import { GetterTree } from 'vuex';
import { Nature, NatureState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<NatureState, RootState> = {
  isInitialized: (state: NatureState): boolean => {
    return state.natures.length > 0;
  },
  natures: (state: NatureState): Nature[] => {
    return state.natures;
  },
  current: (state: NatureState): Nature => {
    return state.current;
  }
}
