import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { SpeciesState, defaultSpecies } from './types';
import { RootState } from '@/store/types';

export const state: SpeciesState = {
  species: defaultSpecies(),
}

const namespaced: boolean = true;

export const speciesState: Module<SpeciesState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
