import { StatsCalculator } from '@/domain/stats';
import { MutationTree } from 'vuex';
import { StatePatterns } from './statePattern';
import Individuals, { TargetState } from './types';
import Species from './types';
import { BasePoints } from './types';

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
    state.currentAbility = payload[0];
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

