import { MutationTree } from 'vuex';
import { StatePatterns } from './statePattern';
import { TargetState } from './types';
import { BasePoints } from './types';
import { Species } from './species'

export const mutations: MutationTree<TargetState> = {
  setSpecies(state, payload: Species) {
    state.species = payload;
  },
  setCurrentAbility(state, payload: string) {
    state.currentAbility = payload;
  },
  changeSlowest(state, isSlowest: boolean) {
    state.individuals.slowest(isSlowest);
  },
  changeWeakest(state, isWeakest: boolean) {
    state.individuals.weakest(isWeakest);
  },
  changeBasePoints(state, payload: BasePoints) {
    state.basePoints = payload;
  },
  changeNature(state, payload: string) {
    state.nature = payload;
  },
  setStatsPattern(state, payload: StatePatterns) {
    state.statePatterns = payload;
  },

}

