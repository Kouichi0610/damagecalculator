import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { AttackerState, MoveCount } from './types';
import { RootState } from '@/store/types';
import { MoveInfo } from '../../moves/store/types';

export const state: AttackerState = {
  current: 0,
  moveA: MoveInfo.empty(),
  moveB: MoveInfo.empty(),
  moveC: MoveInfo.empty(),
  moveD: MoveInfo.empty(),
}

const namespaced: boolean = true;

export const attacker: Module<AttackerState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
