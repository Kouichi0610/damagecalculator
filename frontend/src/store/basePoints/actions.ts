import { ActionTree } from 'vuex';
import { RootState } from '@/store/types';
import { BasePointsState } from './types'
import { state } from './index'

export const actions: ActionTree<BasePointsState, RootState> = {
  initialize: ({ commit }, target: string) => {
    if (state.target == target) return;
    commit('reset', target);
  }
}
