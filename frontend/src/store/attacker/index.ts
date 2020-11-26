import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { AttackerState } from './types';
import { RootState } from '@/store/types';

export const state: AttackerState = {
  currentIndex: 0,
  moveA: '',
  moveB: '',
  moveC: '',
  moveD: '',
  attacker: '',
  physicals: [],
  specials: [],
  result: [],
}

const namespaced: boolean = true;

export const attackerState: Module<AttackerState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
