import { ActionTree } from 'vuex';
import { RootState } from '@/store/types';
import { SpeedState, SpeedList, SpeedInfo, GetSpeed } from './types'
import { state } from './index'
import { TargetCondition } from '../target/targetCondition'

export const actions: ActionTree<SpeedState, RootState> = {
  getSpeed: ({ commit }, target: TargetCondition) => {
    new GetSpeed().getSpeed(target)
    .then((speed) => {
      commit('setTargetSpeed', speed);
    });
  },
  speedOrders: ({commit}, trickRoom: boolean): SpeedInfo[] =>  {
    return [];
  },
  loadList: ({ commit }, level: number) => {
    if (level == 0) return;
    if (state.level == level) return;
    new SpeedList().speedList(level)
    .then((list) => {
      commit('setList', list);
      commit('setLevel', level);
    })
  }
}
