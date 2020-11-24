import { Module } from 'vuex'
import { getters } from './getters';
import { mutations } from './mutations';
import { Individuals, IndividualsState } from './types';
import { RootState } from '@/store/types';

export const state: IndividualsState = {
  targetsIndividuals: Individuals.default(),
}

const namespaced: boolean = true;

export const individualsState: Module<IndividualsState, RootState> = {
  namespaced,
  state,
  getters,
  mutations,
}
