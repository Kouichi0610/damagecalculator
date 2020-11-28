import { ActionTree } from 'vuex';
import { RootState } from '@/store/types';
import { IndividualsState } from './types'
import { state } from './index'

export const actions: ActionTree<IndividualsState, RootState> = {
  initialize: ({ commit }, target: string) => {
    if (state.target == target) return;
    commit('reset', target);
  }
}
