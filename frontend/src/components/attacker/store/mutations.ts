import { MutationTree } from 'vuex';
import { AttackerState, MoveCount } from './types';

export const mutations: MutationTree<AttackerState> = {
  setCurrent(state, payload: number) {
    if (payload < 0 || payload >= MoveCount) {
      return;
    }
    state.current = payload;
  },

}
