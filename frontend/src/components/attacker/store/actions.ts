import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types';
import { AttackerState } from './types'
import { MoveInfo } from '@/components/moves/store/types';

export const actions: ActionTree<AttackerState, RootState> = {
  setMove: ({commit}, info: MoveInfo) => {
    commit('setMove', info);
  }
}

