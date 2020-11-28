import { MutationTree } from 'vuex';
import { Nature, NatureState } from './types';

export const mutations: MutationTree<NatureState> = {
  setTarget(state, payload: string) {
    state.target = payload;
  },
  setList(state, payload: Nature[]) {
    state.natures = payload;
    state.current = payload[0];
  },
  setCurrent(state, payload: Nature) {
    state.current = payload;
  },
  clearCurrent(state) {
    if (state.natures.length == 0) return;
    state.current = state.natures[0];
  }
}
