import { GetterTree } from 'vuex';
import { NatureState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<NatureState, RootState> = {
  natures: (state: NatureState) => {
    return state.list;
  },
  current: (state: NatureState) => {
    return state.current;
  }
}
