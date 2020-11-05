import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { SpeedOrderState } from './types';
import { RootState } from '@/store/types';
import { DefaultSpeedCorrector, SpeedCorrector } from './types'

export const state: SpeedOrderState = {
  list: [],
  abilityOwner: DefaultSpeedCorrector(),
  abilityOther: DefaultSpeedCorrector(),
  // TODO:サーバから取り直す
  item: [new SpeedCorrector('持ち物なし', 0, 1.0), new SpeedCorrector('くろいてっきゅう', 0, 0.5), new SpeedCorrector('こだわりスカーフ', 0, 1.5)],
}

const namespaced: boolean = true;

export const speedOrder: Module<SpeedOrderState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
