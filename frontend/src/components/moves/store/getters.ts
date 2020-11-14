import { GetterTree } from 'vuex';
import { MovesState, MoveInfo } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<MovesState, RootState> = {
  physicals: (state: MovesState): MoveInfo[] => {
    return state.info.physicals();
  },
  specials: (state: MovesState): MoveInfo[] => {
    return state.info.specials();
  },

}
