import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { TargetState, Species } from './types';
import Individuals from './types'
import BasePoints from './types'
import { RootState } from '@/store/types';
import * as nature from "../../../domain/nature"

export const state: TargetState = {
  name: '',
  types: [''],
  species: {hp: 0, at: 0, df: 0, sa: 0, sd: 0, sp: 0},
  weight: 0.0,
  abilities: [''],

  nature: nature.NatureFactory('てれや'),
  individuals: new Individuals(31,31,31,31,31,31),
  basePoints: {hp: 0, at: 0, df: 0, sa: 0, sd: 0, sp: 0},
}

const namespaced: boolean = true;

export const target: Module<TargetState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
