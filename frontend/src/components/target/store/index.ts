import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { TargetState } from './types';
import { Species } from './species'
import { RootState } from '@/store/types';
import { StatePatterns } from './statePattern'
import { Individuals } from './individuals';
import { BasePoints } from './basePoints'

export const state: TargetState = {
  species: Species.default(),

  currentAbility: '',
  nature: '',
  individuals: new Individuals(),
  basePoints: BasePoints.default(),
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
