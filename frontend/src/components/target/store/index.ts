import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { TargetState } from './types';
import { Species } from './species'
import { RootState } from '@/store/types';
import {StatePatterns} from './statePattern'

export const state: TargetState = {
  species: Species.default(),

  currentAbility: '',
  nature: '',
  individuals: {hp: 31, at: 31, df: 31, sa: 31, sd: 31, sp: 31},
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
