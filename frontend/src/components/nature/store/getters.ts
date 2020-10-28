import { GetterTree } from 'vuex';
import { NatureState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<NatureState, RootState> = {
  initialized: (state: NatureState) => {
    return state.list.length > 0;
  },
  natures: (state: NatureState) => {
    return state.list;
  },
  current: (state: NatureState) => {
    return state.current;
  }
}
