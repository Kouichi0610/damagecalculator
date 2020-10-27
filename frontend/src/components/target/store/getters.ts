import { GetterTree } from 'vuex';
import { TargetState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<TargetState, RootState> = {
  name: (state: TargetState) => {
    return state.name;
  },
  types: (state: TargetState) => {
    return state.types;
  },
  species: (state: TargetState) => {
    return state.species;
  },
  weight: (state: TargetState) => {
    return state.weight;
  },
  nature: (state: TargetState) => {
    return state.nature;
  },
  individuals: (state: TargetState) => {
    return state.individuals;
  },
  basePoints: (state: TargetState) => {
    return state.basePoints;
  }
}
