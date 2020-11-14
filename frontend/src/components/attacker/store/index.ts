import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { AttackerState, MoveCount } from './types';
import { RootState } from '@/store/types';
import { DefendersResult } from '../defenderDamages'

export const state: AttackerState = {
  // target 入れ替え時初期化
  target: '',
  current: 0,
  results: DefendersResult.Defaults(MoveCount),
}

const namespaced: boolean = true;

export const attacker: Module<AttackerState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
