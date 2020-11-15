import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { TargetState } from './types';
import { Species } from './species'
import { RootState } from '@/store/types';
import { StatePatterns } from './statePattern'
import { Individuals } from './individuals';

export const state: TargetState = {
  species: Species.default(),

  currentAbility: '',
  nature: '',
  individuals: new Individuals(),
  basePoints: {hp: 0, at: 0, df: 0, sa: 0, sd: 0, sp: 0},
  statePatterns: StatePatterns.default(),
}

const namespaced: boolean = true;

export const target: Module<TargetState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
