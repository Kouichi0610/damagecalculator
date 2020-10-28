import { MutationTree } from 'vuex';
import { NatureState } from './types';
import { Nature } from './types';

export const mutations: MutationTree<NatureState> = {
  setlist(state, payload: Nature[]) {
    state.list = payload;
    state.current = state.list[0].name;
  },
  change(state, payload: string) {
    state.current = payload;
  }
}
