import { MutationTree } from 'vuex';
import { TargetState, Species } from './types';

export const mutations: MutationTree<TargetState> = {
  setName(state, payload: string) {
    state.name = payload;
  },
  setTypes(state, payload: string[]) {
    state.types = payload;
  },
  setWeight(state, payload: number) {
    state.weight = payload;
  },
  setSpecies(state, payload: Species) {
    state.species = payload;
  },
  setAbilities(state, payload: string[]) {
    state.abilities = payload;
    for (var i = 0; i < state.abilities.length; i++) {
      console.log(' ' + state.abilities[i]);
    }
  }
}

