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
    state.individuals = {
      hp: 31,
      at: 31,
      df: 31,
      sa: 31,
      sd: 31,
      sp: isSlowest ? 0 : 31,
    };
  },
  changeWeakest(state, isWeakest: boolean) {
    state.individuals = {
      hp: 31,
      at: isWeakest ? 0 : 31,
      df: 31,
      sa: 31,
      sd: 31,
      sp: 31,
    };
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

