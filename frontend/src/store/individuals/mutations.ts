import { MutationTree } from 'vuex';
import { Individuals, IndividualsState } from './types';

export const mutations: MutationTree<IndividualsState> = {
  reset(state, target: string) {
    state.target = target;
    state.targetsIndividuals = Individuals.default();
  },
  changeSlowest(state) {
    state.targetsIndividuals = state.targetsIndividuals.changeSlowest();
  },
  changeWeakest(state) {
    state.targetsIndividuals = state.targetsIndividuals.changeWeakest();
  }
}
