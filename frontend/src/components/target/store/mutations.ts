import { StatsCalculator } from '@/domain/stats';
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
  },
  changeSlowest(state, payload: boolean) {
    state.individuals.changeSlowest(payload);
  },
  changeWeakest(state, payload: boolean) {
    state.individuals.changeWeakest(payload);
  },
}

