import { GetterTree } from 'vuex';
import { SpeciesState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<SpeciesState, RootState> = {
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
