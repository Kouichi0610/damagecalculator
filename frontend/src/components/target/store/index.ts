import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { TargetState } from './types';
import { Species } from './types'
import { Individuals } from './types'
import { BasePoints } from './types'
import { RootState } from '@/store/types';

export const state: TargetState = {
  name: '',
  types: [''],
  species: {hp: 0, at: 0, df: 0, sa: 0, sd: 0, sp: 0},
  weight: 0.0,
  abilities: [''],

  nature: '',
  individuals: {hp: 31, at: 31, df: 31, sa: 31, sd: 31, sp: 31},
  basePoints: {hp: 0, at: 0, df: 0, sa: 0, sd: 0, sp: 0},
  // TODO:species-individuals-basepoints-patternsセットでコンポーネントにする
  hppattern: [],
  atpattern: [],
  dfpattern: [],
  sapattern: [],
  sdpattern: [],
  sppattern: [],
}

const namespaced: boolean = true;

export const target: Module<TargetState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
