import { MutationTree } from 'vuex';
import { SpeciesState, PokeData } from './types';

export const mutations: MutationTree<SpeciesState> = {
  setData(state, payload: PokeData) {
    state.data = payload;
  }
}
